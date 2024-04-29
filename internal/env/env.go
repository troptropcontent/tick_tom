package env

import "os"

// This function return the value of the environment variable passed as parameter. It panics if the environment variable is not set.
func Require(env_var string) string {
	value := os.Getenv(env_var)
	if value == "" {
		message := "The " + env_var + " seems to not be set"
		panic(message)
	}
	return value
}

// This function return the current environment it can be either production, development or test. It panics if the environment is not one of the three.
func Current() string {
	env := os.Getenv("FOO_ENV")
	if env == "" {
		env = "development"
	}

	available_env := map[string]string{
		"production":  "",
		"development": "",
		"test":        "",
	}

	if _, ok := available_env[env]; !ok {
		panic("GO_ENV " + env + " is not a valid environment")
	}

	return env
}
