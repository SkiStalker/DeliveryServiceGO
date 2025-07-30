package util

func GetDBDSN() string {
	pg_user := GetOSEnvWithDefault("POSTGRES_USER", "postgres")

	pg_pwd := GetOSEnvWithDefault("POSTGRES_PASSWORD", "postgres")

	pg_host := GetOSEnvWithDefault("POSTGRES_HOST", "localhost")
	pg_port := GetOSEnvWithDefault("POSTGRES_PORT", "5432")
	pg_db := GetOSEnvWithDefault("POSTGRES_DB", "company")

	return "postgresql://" + pg_user + ":" + pg_pwd + "@" + pg_host + ":" + pg_port + "/" + pg_db
}
