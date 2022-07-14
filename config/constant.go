package config

// noinspection ALL
const (
	ENV             = "ENV"
	ENV_DEVELOPMENT = "development"
	ENV_QA          = "qa"
	ENV_MOCK        = "mock"
	ENV_PRODUCTION  = "production"
)

// noinspection ALL
const (
	HTTP_PORT          = "HTTP_PORT"
	HTTP_PROFILER_PORT = "HTTP_PROFILER_PORT"

	DATABASE_CONNECTION_STRING = "DATABASE_CONNECTION_STRING"

	DATABASE_DRIVER = "DATABASE_DRIVER"
	DATABASE_HOST   = "DATABASE_HOST"
	DATABASE_PORT   = "DATABASE_PORT"
	DATABASE_USER   = "DATABASE_USER"
	DATABASE_PASS   = "DATABASE_PASS"
	DATABASE_NAME   = "DATABASE_NAME"
	DATABASE_SSL    = "DATABASE_SSL"

	MIGRATION_PATH = "MIGRATION_PATH"

	USER_DEFAULT_ROLE = "USER_DEFAULT_ROLE"

	VAULT_TOKEN   = "VAULT_TOKEN"
	VAULT_ADDRESS = "VAULT_ADDRESS"
	SERVICE_NAME  = "SERVICE_NAME"

	EMAIL_ADDRESS  = "EMAIL_ADDRESS"
	EMAIL_PASSWORD = "EMAIL_PASSWORD"
	EMAIL_HOST     = "EMAIL_HOST"
	EMAIL_PORT     = "EMAIL_PORT"
)
