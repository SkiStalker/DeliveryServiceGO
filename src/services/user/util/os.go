package util

import "os"

func GetOSEnvWithDefault(key string, default_value string) string {
	res := os.Getenv(key)
	if res == "" {
		res = default_value
	}
	return res
}
