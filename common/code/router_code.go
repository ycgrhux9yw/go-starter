package code

import (
	"bytes"
	"text/template"
)

// RouterCodeGenerate generates the main router setup file for the project.
// It creates a basic HTTP router using the Echo framework with health check
// and versioned API group endpoints.
func RouterCodeGenerate(projectName string) ([]byte, error) {
	const routerTemplate = `package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter initializes and returns a configured Echo router instance.
// It registers global middleware and all application routes.
func NewRouter() *echo.Echo {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	// Health check endpoint
	e.GET("/health", healthCheck)

	// API v1 group
	v1 := e.Group("/api/v1")
	registerV1Routes(v1)

	return e
}

// healthCheck handles the health check endpoint.
// Returns HTTP 200 with a simple status message.
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "{{.ProjectName}}",
	})
}

// registerV1Routes registers all v1 API routes to the provided group.
func registerV1Routes(g *echo.Group) {
	// Register your routes here
	// Example:
	// g.GET("/users", userHandler.GetAll)
	// g.POST("/users", userHandler.Create)
	_ = g
}
`

	tmpl, err := template.New("router").Parse(routerTemplate)
	if err != nil {
		return nil, err
	}

	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
