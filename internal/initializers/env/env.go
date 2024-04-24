package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	fmt.Println("Initializing env")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func Require(env_var string) string {
	value := os.Getenv(env_var)
	if value == "" {
		message := "The " + env_var + " seems to not be set"
		panic(message)
	}
	return value
}
