package server

import (
	"erpaa-backend/internal/database"
	"erpaa-backend/internal/handler"
	middleware "erpaa-backend/internal/middlewares"
	"erpaa-backend/internal/repository"
	"log"
	"net/http"
	"time"
)

const (
	FrontendPages = "/var/www/erpaa/app/frontend/pages"
	FrontendAsset = "/var/www/erpaa/app/frontend/assets"
)

func Server() error {

	// ===== DB & DEPENDENCY =====
	db := database.DataCon()
	userRepo := repository.NewUserImplemen(db)
	userHandler := handler.NewHandlerUser(userRepo)

	mux := http.NewServeMux()

	// ===== STATIC FILE =====
	mux.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir(FrontendAsset)),
		),
	)

	// ===== PUBLIC PAGE =====
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, FrontendPages+"/login.html")
	})

	// ===== AUTH =====
	mux.Handle("/login/auth", middleware.Chain(
		http.HandlerFunc(userHandler.FindUser),
		middleware.LoggingMiddleware,
		middleware.RecoverMiddleware,
	))

	mux.Handle("/login/register", middleware.Chain(
		http.HandlerFunc(userHandler.Registrasi),
		middleware.LoggingMiddleware,
		middleware.RecoverMiddleware,
	))

	mux.HandleFunc("/logout/", userHandler.Logout)

	// ===== PROTECTED PAGE =====
	mux.Handle("/dashboard/", middleware.AuthMiddleware(
		fileHandler("dashboard.html"),
	))

	mux.Handle("/inventory/", middleware.AuthMiddleware(
		fileHandler("inventory.html"),
	))
		mux.Handle("/inventory/stok/", middleware.AuthMiddleware(
		fileHandler("stokmanagement.html"),
	))

	mux.Handle("/product/", middleware.AuthMiddleware(
		fileHandler("product.html"),
	))

	mux.Handle("/shipping/", middleware.AuthMiddleware(
		fileHandler("shipping.html"),
	))

	mux.Handle("/salesforecasting/", middleware.AuthMiddleware(
		fileHandler("salesforecasting.html"),
	))

	mux.Handle("/integrasi/", middleware.AuthMiddleware(
		fileHandler("Integrasi.html"),
	))

	mux.Handle("/dashboard/daily", middleware.AuthMiddleware(
		fileHandler("DailyBrief.html"),
	))

	// ===== SERVER =====
	server := &http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("✅ ERPAA server running at :8090")
	return server.ListenAndServe()
}

// ===== HELPER =====
func fileHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, FrontendPages+"/"+filename)
	}
}
