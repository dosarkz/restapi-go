# Rest api


# Modules
- Auth JWT
- RBAC casbin
- Playground validation
- Mux routing
- Localization
- Session
- ORM gorm
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

# Casbin
Role Based Access Control (RBAC) implemented with [casbin](https://casbin.org/docs/en/syntax-for-models)
- ./auth_model.conf - file with ACL configs
- ./policy.csv - file with policy rules

All routes locked by default, so we should open root path prefix and close concrete endpoints. Deny rules override allow rules

# Locales and translations
Available Russian and English locales and validations
- locale destination: ```routes/lang/locales```
- yml file supports

Example:
```responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errors)```
