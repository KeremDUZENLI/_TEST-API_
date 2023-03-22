package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	URL string
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
}
