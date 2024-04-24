package env_initializer

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() {
	fmt.Println("Initializing env")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
