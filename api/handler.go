// Package api .
package api

import (
	"net/http"

	"github.com/crispgm/foosbot/internal/app"
	"github.com/gin-gonic/gin"
)

// Handler Vercel faas function
func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	app.LoadRoutes(router)
	router.ServeHTTP(w, r)
}
