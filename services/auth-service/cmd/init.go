package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Panic("❌ No .env file found, loading from system env")
	}

	log.Println("✅ loaded .env ...")

}
