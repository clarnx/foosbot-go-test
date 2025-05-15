// Package main CLI entry for foosbot
package main

import (
	"github.com/crispgm/foosbot/internal/app"
	"github.com/crispgm/foosbot/internal/def"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := def.LoadVariables()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	app.LoadRoutes(router)
	router.Run(def.Port)
}
