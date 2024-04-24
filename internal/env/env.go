package env

import "os"

func Require(env_var string) string {
	value := os.Getenv(env_var)
	if value == "" {
		message := "The " + env_var + " seems to not be set"
		panic(message)
	}
	return value
}
