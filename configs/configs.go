package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Define the structure of your JSON object here
	Key1 string `json:"key1"`
	Key2 int    `json:"key2"`
	// Add more fields as needed
}

func ConfigsLoad() {
	godotenv.Load()
	serviceName := os.Getenv("SERICE_Name")
	fmt.Println("service name from config", serviceName)
	// Read JSON string from environment variable
	jsonObjString := os.Getenv("JSON_OBJ")
	if jsonObjString == "" {
		log.Fatal("JSON_OBJ environment variable is not set")
	}

	// Parse JSON object
	var config Config
	if err := json.Unmarshal([]byte(jsonObjString), &config); err != nil {
		log.Fatalf("Error decoding JSON object: %v", err)
	}

	fmt.Println(config)

}
