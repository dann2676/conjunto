package main

import (
	ows "asamblea/internal/domain/owner"
	as "asamblea/internal/domain/unit"
	or "asamblea/internal/platform/owner"
	ar "asamblea/internal/platform/unit"
	"asamblea/internal/providers/storage"
	"asamblea/internal/web"
	"asamblea/internal/web/owner"
	"asamblea/internal/web/unit"
	"embed"
	"html/template"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var fS embed.FS

func main() {

	// Create a Gin router with default middleware (logger and recovery)
	var handler slog.Handler

	if os.Getenv("ENV") == "production" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		// Para tu proyecto personal en tu PC, este es el que quieres
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(handler))

	r := gin.Default()

	templ := template.Must(template.New("").ParseFS(fS, "templates/**/*.html"))

	// 2. Pasarle esas plantillas a Gin
	r.SetHTMLTemplate(templ)

	h := web.New()
	db, err := storage.Init()
	if err != nil {
		slog.Error("cound not init db", "error", err.Error())
		return
	}
	unitRepo := ar.New(db)
	unitService := as.New(unitRepo)
	unit := unit.New(unitService)

	ownerRepo := or.New(db)
	ownerService := ows.New(ownerRepo)
	owner := owner.New(ownerService, unitService)

	// Define a simple GET endpoint
	r.GET("/ping", h.Ping)

	r.GET("/units", unit.GetAll)
	r.GET("/units/list", unit.GetAllList)
	r.POST("/units", unit.Create)
	r.DELETE("/units/:id", unit.Delete)
	r.DELETE("/units/:id/purge", unit.Delete)
	r.PUT("/units/:id", unit.Update)
	r.GET("/units/form/:id", unit.EditForm)
	r.GET("/units/:id/owners", unit.EditForm)

	r.GET("/owners", owner.GetAll)
	r.GET("/owners/list", owner.GetAllList)
	r.POST("/owners", owner.Create)
	r.DELETE("/owners/:id", owner.Delete)
	r.DELETE("/owners/:id/purge", owner.Purge)
	r.PUT("/owners/:id", owner.Update)
	r.GET("/owners/form/:id", owner.EditForm)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
