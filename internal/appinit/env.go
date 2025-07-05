package appinit

import (
	"log"

	"ExBot/internal/texts"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(texts.LogEnvNotFound)
	} else {
		log.Println(texts.LogEnvLoaded)
	}
}
