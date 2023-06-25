package main

import (
	"io/ioutil"
	"log"

	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
	Port     string `yaml:"port"`
}

type RoutesConfig struct {
	Routes map[string]string `yaml:"routes"`
}

func main() {
	// Read YAML config file
	configFile := "/etc/detour/config.yaml" // Replace with your YAML config file path
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Parse YAML config
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}

	// Read YAML routes file
	routesFile := "/etc/detour/routes.yaml" // Replace with your YAML routes file path
	routesData, err := ioutil.ReadFile(routesFile)
	if err != nil {
		log.Fatalf("Error reading routes file: %s", err)
	}

	// Parse YAML routes
	var routesConfig RoutesConfig
	err = yaml.Unmarshal(routesData, &routesConfig)
	if err != nil {
		log.Fatalf("Error parsing routes file: %s", err)
	}

	redirectHandler := func(ctx *fasthttp.RequestCtx) {
		requestURI := string(ctx.Request.URI().Path())
		redirectURL, ok := routesConfig.Routes[requestURI]
		if !ok {
			ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
			return
		}

		// Set the appropriate redirect status code
		ctx.Response.SetStatusCode(fasthttp.StatusMovedPermanently)

		// Set the a server header
		ctx.Response.Header.Set("Server", "deTour proxy")

		// Set the Location header to the redirect URL
		ctx.Response.Header.Set("Location", redirectURL)
	}

	// Create the FastHTTP server
	server := &fasthttp.Server{
		Handler: redirectHandler,
	}

	// Start the server with TLS support
	err = server.ListenAndServeTLS(config.Port, config.CertFile, config.KeyFile)
	if err != nil {
		log.Fatalf("Error in ListenAndServeTLS: %s", err)
	}
}
