package main

import (
	as "asamblea/internal/domain/apartment"
	ows "asamblea/internal/domain/owner"
	ar "asamblea/internal/platform/apartment"
	or "asamblea/internal/platform/owner"
	"asamblea/internal/providers/storage"
	"asamblea/internal/web"
	"asamblea/internal/web/apartment"
	"asamblea/internal/web/owner"
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
	apartmentRepo := ar.New(db)
	apartmentService := as.New(apartmentRepo)
	apartment := apartment.New(apartmentService)

	ownerRepo := or.New(db)
	ownerService := ows.New(ownerRepo)
	owner := owner.New(ownerService, apartmentService)

	// Define a simple GET endpoint
	r.GET("/ping", h.Ping)

	r.GET("/apartments", apartment.GetAll)
	r.GET("/apartments/list", apartment.GetAllList)
	r.POST("/apartments", apartment.Create)
	r.DELETE("/apartments/:id", apartment.Delete)
	r.PUT("/apartments/:id", apartment.Update)
	r.GET("/apartments/form/:id", apartment.EditForm)

	r.GET("/owners", owner.GetAll)
	r.GET("/owners/list", owner.GetAllList)
	r.POST("/owners", owner.Create)
	r.DELETE("/owners/:id", owner.Delete)
	r.PUT("/owners/:id", owner.Update)
	r.GET("/owners/form/:id", owner.EditForm)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
