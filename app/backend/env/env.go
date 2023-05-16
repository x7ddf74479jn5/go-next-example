package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const targetEnvName = "local"

func Load() {
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv(targetEnvName)))
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}
}
