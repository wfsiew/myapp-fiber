package config

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func Config(key string) string {
    // load .env file
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("Error loading .env file")
    }
        // Return the value of the variable
    return os.Getenv(key)
}