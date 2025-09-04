# Spy Cat Agency

A CRUD RESTful API for managing spy cats, missions, and targets.

## Overview

This application is built for the Spy Cat Agency assessment. It demonstrates:
- RESTful API design using Go and a modern framework
- SQL database integration
- Third-party API validation ([TheCatAPI](https://api.thecatapi.com/v1/breeds))
- Logger middleware for HTTP requests/responses
- Request validation and error handling
- Interactive API documentation (Swagger)

## Features

- Manage spy cats (create, update salary, delete, list, get)
- Manage missions and targets (create, update, delete, assign cats, mark complete)
- Validate cat breeds via TheCatAPI
- Swagger UI for API exploration

## Setup & Run

Clone and start the application (default config uses local Postgres):

```sh
git clone https://github.com/refilutub/spy-cat-agency.git
cd spy-cat-agency
docker-compose up --build
```
Now you can test the app! 
- The Swagger API specification is available at `http://localhost:8080/swagger/index.html`.
- The API base URL is `http://localhost:8080/api/`.