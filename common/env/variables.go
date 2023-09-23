package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB  string
	URL string
)

func Load() {
	godotenv.Load(".env")

	DB = os.Getenv("DB")
	URL = os.Getenv("URL")
}
