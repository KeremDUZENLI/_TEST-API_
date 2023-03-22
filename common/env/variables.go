package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	URL    string
	URL_MY string
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
	URL_MY = os.Getenv("URL_MY")
}
