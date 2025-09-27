package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configurations Config

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println()
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	portStr := os.Getenv("HTTP_PORT")

	if version == "" || serviceName == "" || portStr == "" {
		fmt.Println("Missing required environment variables: VERSION, SERVICE_NAME, or HTTP_PORT")
		os.Exit(1)
	}

	// port, err := strconv.Atoi(portStr)
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid HTTP_PORT value: %s (must be a number)\n", portStr)
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
