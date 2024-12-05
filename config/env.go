package config

type Config struct {
	PublicHost string
	Port	   string
	DBUser     string
	DBPassword   string
	DBAddress     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv*("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost:3306"), 
				getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecom"),
		 
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
