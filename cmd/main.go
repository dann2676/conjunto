package main

import (
	asm "asamblea/internal/domain/assembly"
	ows "asamblea/internal/domain/owner"
	as "asamblea/internal/domain/unit"
	asmr "asamblea/internal/platform/assembly"
	or "asamblea/internal/platform/owner"
	ar "asamblea/internal/platform/unit"
	"asamblea/internal/providers/storage"
	"asamblea/internal/web"
	"asamblea/internal/web/assembly"
	"asamblea/internal/web/owner"
	"asamblea/internal/web/unit"
	"embed"
	"html/template"
	"log/slog"
	"os"
	"time"

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

	templ := template.Must(template.New("").Funcs(template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"mul": func(a, b float32) float32 {
			return a * b
		},
		"sub": func(a, b float32) float32 {
			return a - b
		},
		"minFloat": func(a, b float32) float32 {
			if a < b {
				return a
			}
			return b
		},
	}).ParseFS(fS, "templates/**/*.html"))

	// 2. Pasarle esas plantillas a Gin
	r.SetHTMLTemplate(templ)

	h := web.New()
	db, err := storage.Init()
	if err != nil {
		slog.Error("cound not init db", "error", err.Error())
		return
	}
	unitRepo := ar.New(db)
	ownerRepo := or.New(db)

	unitService := as.New(unitRepo)
	ownerService := ows.New(ownerRepo)

	owner := owner.New(ownerService, unitService)
	unit := unit.New(unitService, ownerService)

	// Define a simple GET endpoint
	r.GET("/ping", h.Ping)

	r.GET("/units", unit.GetAll)
	r.GET("/units/list", unit.GetAllList)
	r.POST("/units", unit.Create)
	r.DELETE("/units/:id", unit.Delete)
	r.DELETE("/units/:id/purge", unit.Purge)
	r.PUT("/units/:id", unit.Update)
	r.GET("/units/form/:id", unit.EditForm)
	r.GET("/units/:id/owners", unit.GetOwner)

	r.GET("/owners", owner.GetAll)
	r.GET("/owners/list", owner.GetAllList)
	r.POST("/owners", owner.Create)
	r.DELETE("/owners/:id", owner.Delete)
	r.DELETE("/owners/:id/purge", owner.Purge)
	r.PUT("/owners/:id", owner.Update)
	r.GET("/owners/form/:id", owner.EditForm)

	assemblyRepo := asmr.New(db)
	assemblyService := asm.New(assemblyRepo)
	assemblyHandler := assembly.New(assemblyService, unitService)

	r.GET("/assemblies", assemblyHandler.GetAll)
	r.GET("/assemblies/form/:id", assemblyHandler.EditForm)
	r.POST("/assemblies", assemblyHandler.Create)
	r.PUT("/assemblies/:id", assemblyHandler.Update)
	r.DELETE("/assemblies/:id", assemblyHandler.Delete)

	r.GET("/assemblies/:id/admin", assemblyHandler.Admin)
	r.GET("/assemblies/:id/quorum", assemblyHandler.GetQuorum)
	r.POST("/assemblies/:id/agenda", assemblyHandler.CreateAgendaItem)
	r.DELETE("/assemblies/:id/agenda/:item_id", assemblyHandler.DeleteAgendaItem)
	r.PUT("/assemblies/:id/agenda/:item_id/status/:status", assemblyHandler.UpdateAgendaItemStatus)
	r.PUT("/assemblies/:id/status/:status", assemblyHandler.UpdateStatus)

	r.GET("/assemblies/:id/attendance", assemblyHandler.AttendancePage)
	r.POST("/assemblies/:id/attendance", assemblyHandler.RegisterAttendance)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
