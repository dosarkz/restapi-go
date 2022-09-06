# Simple rest api (Gorilla/mux,ddd,gorm,jwt) 
This project a basic example for personal use, I hope it will be useful =)

# Features
- JWT authentication
- [Go-playground validator](https://github.com/go-playground/validator)
- [Gorilla/mux](https://github.com/gorilla/mux)
- Localization and response translation
- Session
- Database migrations
- [GORM ORM](https://gorm.io/index.html)
- Logging and monitoring
- Rate limits

# Stack
- Golang 1.18
- Postgresql
- Redis
- Docker
- Prometheus
- Grafana

## Goland IDE 
- [Link to install](https://www.jetbrains.com/go)

The extensions which must have to install:
- Go linter
- Git commit template

## Docker 
local env
- APP_URL=0.0.0.0

Command to run containers
`DOCKER_BUILDKIT=1 docker-compose -f docker-compose.main.yml -f docker-compose.local.yml up -d --build`

## Windows 
Make sure that files of docker has unix format

# Deploy locally
- DATABASE_HOST=localhost
- DATABASE_PORT=5411 // your DOCKER_POSTGRES_PORT
- APP_URL=localhost

# Locales and translations
Available Russian and English locales and validations
- locale destination: ```routes/lang/locales```
- yml file supports

Example:
```responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errors)```
