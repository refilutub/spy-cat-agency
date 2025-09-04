package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"spy-cat-agency/internal/cats/dtos"
	"spy-cat-agency/internal/cats/mappers"
	"sync"
	"time"
)

// thread-safe in-memory cache for breed names
var (
	breedCacheTTL = 6 * time.Hour
	breedCache    = struct {
		mu        sync.RWMutex
		names     []dtos.BreedName
		version   string
		expiresAt time.Time
	}{}
	sharedHTTPClient = &http.Client{Timeout: 5 * time.Second}
)

func getBreedNames() ([]dtos.BreedName, error) {
	// Fast path: return cached if still valid
	now := time.Now()
	breedCache.mu.RLock()
	if breedCache.names != nil && now.Before(breedCache.expiresAt) {
		n := make([]dtos.BreedName, len(breedCache.names))
		copy(n, breedCache.names)
		breedCache.mu.RUnlock()
		return n, nil
	}
	breedCache.mu.RUnlock()

	// Slow path: attempt to refresh under write lock
	breedCache.mu.Lock()
	defer breedCache.mu.Unlock()
	// Double-check after acquiring lock
	if breedCache.names != nil && time.Now().Before(breedCache.expiresAt) {
		n := make([]dtos.BreedName, len(breedCache.names))
		copy(n, breedCache.names)
		return n, nil
	}

	// Try to get API version first; use it to avoid refetching breeds if unchanged
	apiVersion := ""
	if respAPIv, err := sharedHTTPClient.Get("https://api.thecatapi.com/v1"); err == nil {
		defer func() {
			if errClose := respAPIv.Body.Close(); errClose != nil {
				log.Printf("Error closing response body: %v", errClose)
			}
		}()
		if respAPIv.StatusCode == http.StatusOK {
			var v struct {
				Version string `json:"version"`
			}
			if err := json.NewDecoder(respAPIv.Body).Decode(&v); err == nil {
				apiVersion = v.Version
			}
		}
	}

	// If version unchanged and we have cached data, extend TTL and return it
	if apiVersion != "" && apiVersion == breedCache.version && breedCache.names != nil {
		breedCache.expiresAt = time.Now().Add(breedCacheTTL)
		n := make([]dtos.BreedName, len(breedCache.names))
		copy(n, breedCache.names)
		return n, nil
	}

	// Otherwise, fetch breeds
	res, err := sharedHTTPClient.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		// Fallback to stale cache if available
		if breedCache.names != nil {
			log.Printf("Error fetching breeds, serving stale cache: %v", err)
			breedCache.expiresAt = time.Now().Add(30 * time.Minute)
			n := make([]dtos.BreedName, len(breedCache.names))
			copy(n, breedCache.names)
			return n, nil
		}
		return nil, err
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()
	if res.StatusCode != http.StatusOK {
		// Fallback to stale cache on non-200
		if breedCache.names != nil {
			log.Printf("Unexpected status %d fetching breeds, serving stale cache", res.StatusCode)
			breedCache.expiresAt = time.Now().Add(30 * time.Minute)
			n := make([]dtos.BreedName, len(breedCache.names))
			copy(n, breedCache.names)
			return n, nil
		}
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var names []dtos.BreedName
	if err := json.NewDecoder(res.Body).Decode(&names); err != nil {
		// Fallback to stale cache on decode error
		if breedCache.names != nil {
			log.Printf("Error decoding breeds, serving stale cache: %v", err)
			breedCache.expiresAt = time.Now().Add(30 * time.Minute)
			n := make([]dtos.BreedName, len(breedCache.names))
			copy(n, breedCache.names)
			return n, nil
		}
		return nil, err
	}

	// Update cache
	breedCache.names = names
	if apiVersion != "" {
		breedCache.version = apiVersion
	}
	breedCache.expiresAt = time.Now().Add(breedCacheTTL)

	n := make([]dtos.BreedName, len(names))
	copy(n, names)
	return n, nil
}

func isValidBreed(breed string, breedNames []dtos.BreedName) bool {
	for _, b := range breedNames {
		if b.Name == breed {
			return true
		}
	}
	return false
}

func (s *spyCatService) CreateSpyCat(spyCatReq dtos.SpyCatRequest) (dtos.SpyCatSingleResponseDTO, error) {
	breedNames, err := getBreedNames()
	if err != nil {
		return dtos.SpyCatSingleResponseDTO{}, err
	}
	if !isValidBreed(spyCatReq.Breed, breedNames) {
		return dtos.SpyCatSingleResponseDTO{}, fmt.Errorf("invalid breed")
	}

	newSpyCatModel := mappers.CreateDTOToSpyCat(spyCatReq)
	persisted, err := s.repo.CreateSpyCat(newSpyCatModel)
	if err != nil {
		return dtos.SpyCatSingleResponseDTO{}, err
	}

	return mappers.SpyCatSingleToDTO(persisted), nil
}
