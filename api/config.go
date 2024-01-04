package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("error loading env file ", err)
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("error loading env file ", err)
	}

	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envSERVICESID() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("error loading env file ", err)
	}

	return os.Getenv("TWILIO_SERVICE_SID")
}
