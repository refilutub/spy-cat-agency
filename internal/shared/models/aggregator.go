package models

import (
	cats "spy-cat-agency/internal/cats/domain/models"
	missions "spy-cat-agency/internal/missions/domain/models"
)

var AllModels = []interface{}{
	&cats.Cat{},
	&missions.Mission{},
	&missions.Target{},
}
