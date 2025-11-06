package envhelper

import "os"

// get an env variable or fallback to a default value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvAsInt(key, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if val, err := strconv.ParseInt(value); err == nil {
			return val
		} 
	}
	return fallback
}