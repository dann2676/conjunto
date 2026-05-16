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
	middleware "asamblea/internal/web/middelware"
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

	assemblyRepo := asmr.New(db)
	assemblyService := asm.New(assemblyRepo, unitRepo)
	assemblyHandler := assembly.New(assemblyService, unitService, ownerService)

	r.GET("/ping", h.Ping)
	// Define a simple GET endpoint

	r.Use(middleware.SecurityHeaders())

	// rutas admin — con auth
	admin := r.Group("/")
	admin.Use(middleware.BasicAuth())
	{
		admin.GET("/units", unit.GetAll)
		admin.GET("/units/list", unit.GetAllList)
		admin.POST("/units", unit.Create)
		admin.DELETE("/units/:id", unit.Delete)
		admin.DELETE("/units/:id/purge", unit.Purge)
		admin.PUT("/units/:id", unit.Update)
		admin.GET("/units/form/:id", unit.EditForm)
		admin.GET("/units/:id/owners", unit.GetOwner)

		admin.GET("/owners", owner.GetAll)
		admin.GET("/owners/list", owner.GetAllList)
		admin.POST("/owners", owner.Create)
		admin.DELETE("/owners/:id", owner.Delete)
		admin.DELETE("/owners/:id/purge", owner.Purge)
		admin.PUT("/owners/:id", owner.Update)
		admin.GET("/owners/form/:id", owner.EditForm)

		admin.GET("/assemblies", assemblyHandler.GetAll)
		admin.GET("/assemblies/form/:id", assemblyHandler.EditForm)
		admin.POST("/assemblies", assemblyHandler.Create)
		admin.PUT("/assemblies/:id", assemblyHandler.Update)
		admin.DELETE("/assemblies/:id", assemblyHandler.Delete)
		admin.GET("/assemblies/:id/admin", assemblyHandler.Admin)
		admin.GET("/assemblies/:id/quorum", assemblyHandler.GetQuorum)
		admin.POST("/assemblies/:id/agenda", assemblyHandler.CreateAgendaItem)
		admin.DELETE("/assemblies/:id/agenda/:item_id", assemblyHandler.DeleteAgendaItem)
		admin.PUT("/assemblies/:id/agenda/:item_id/status/:status", assemblyHandler.UpdateAgendaItemStatus)
		admin.PUT("/assemblies/:id/status/:status", assemblyHandler.UpdateStatus)
		admin.GET("/assemblies/:id/agenda/:item_id/results", assemblyHandler.GetVoteResults)
		admin.GET("/assemblies/:id/report", assemblyHandler.DownloadReport)
		admin.GET("/assemblies/:id/codes", assemblyHandler.GetCodes)
		admin.GET("/assemblies/:id/codes/:code/qr", assemblyHandler.DownloadQR)
		admin.GET("/assemblies/:id/codes/:code/qr-image", assemblyHandler.QRImage)

	}

	public := r.Group("/")
	public.Use(middleware.RateLimitAttendance())
	{
		public.POST("/assemblies/:id/attendance", assemblyHandler.RegisterAttendance)
		public.POST("/assemblies/:id/attendance/lookup", assemblyHandler.LookupOwner)
		public.POST("/assemblies/:id/vote", assemblyHandler.RegisterVote)
	}
	polling := r.Group("/")
	polling.Use(middleware.RateLimitPolling())
	{
		polling.GET("/assemblies/:id/attendance", assemblyHandler.AttendancePage)
		polling.GET("/assemblies/:id/vote", assemblyHandler.VotePage)
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
