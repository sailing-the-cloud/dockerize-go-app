package main

import (
	"fmt"
	configEnv "github.com/joho/godotenv"
	"os"
)

func main() {

	// Untuk dipakai dilokal
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	service := MakeHandler()
	service.HTTPServerMain()
}
