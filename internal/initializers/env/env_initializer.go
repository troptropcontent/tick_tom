package env_initializer

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/troptropcontent/tick_tom/internal/env"
)

func Init() {
	fmt.Println("Loading .env file from: " + env.EnvFilePath())
	err := godotenv.Load(env.EnvFilePath())
	if err != nil {
		panic("Error loading .env file")
	}
}
