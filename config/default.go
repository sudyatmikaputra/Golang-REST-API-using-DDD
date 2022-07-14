package config

func loadDefaultConfig() map[string]string {
	return map[string]string{
		"DATABASE_CONNECTION_STRING": "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"DATABASE_DRIVER":            "postgres",
		"DATABASE_HOST":              "localhost",
		"DATABASE_NAME":              "postgres",
		"DATABASE_PASS":              "password",
		"DATABASE_PORT":              "5432",
		"DATABASE_SSL":               "disable",
		"DATABASE_USER":              "postgres",
		"HTTP_PORT":                  ":8001",
		"HTTP_PROFILER_PORT":         ":9100",
		"MIGRATION_PATH":             "database/migrations",
		"USER_DEFAULT_ROLE":          "becdd3c3-6e9d-4fb9-9f05-6d183c87de16",
	}
}
