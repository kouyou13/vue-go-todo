package config

import "github.com/gin-contrib/cors"

func CORSConfig() cors.Config {
	return cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	}
}
