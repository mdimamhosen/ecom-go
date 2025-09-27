package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux,
		middleware.Logger,
		middleware.TestMid,
		middleware.CorsWithPreflight)

	iniRoutes(mux, manager)

	port := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, wrappedMux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// 01731659763
