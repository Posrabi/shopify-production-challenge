package cmd

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Posrabi/shopify-backend-project/src/common/middleware"
	"github.com/Posrabi/shopify-backend-project/src/common/utils"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	corsWrapper = cors.New(cors.Options{
		MaxAge:         int(10 * time.Minute / time.Second),
		AllowedHeaders: []string{"*"},
	})
	port       = "8081"
	addr       = "localhost:" + port
	invService *api.Service
)

func newServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Inventory tracking",
		Run:   runServerCmd,
	}
}

func runServerCmd(cmd *cobra.Command, args []string) {
	dir, err := os.Getwd()
	if err != nil {
		logrus.Panicf("unable to get current directory: %w", err)
	}
	if err := godotenv.Load(filepath.Join(dir, "build", ".env")); err != nil && utils.FileExists(".env") {
		logrus.Panicf("unable to load .env: %w", err)
	}
	sigs := make(chan os.Signal)

	go func() {
		sig := <-sigs
		time.Sleep(5 * time.Second)
		log.Panic("Exiting server on", sig) // panic allowed deferred calls to run unlike cmd os.Exit -> make sense in this situation
	}()

	db, err := api.NewSQLDB()
	if err != nil {
		logrus.Panicf("unable to connect to postgres database: %w", err)
	}

	invService = api.NewService(api.NewMasterRepository(db))

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.Path("/create").Methods("POST").HandlerFunc(middleware.ErrorWrapper("/create", createHandler))
	r.Path("/edit").Methods("POST").HandlerFunc(middleware.ErrorWrapper("/edit", editHandler))
	r.Path("/delete").Methods("POST").HandlerFunc(middleware.ErrorWrapper("/delete", deleteHandler))
	r.Path("/list").Methods("GET").HandlerFunc(middleware.ErrorWrapper("/list", listHandler))
	r.Path("/ship").Methods("POST").HandlerFunc(middleware.ErrorWrapper("/ship", shipHandler))

	logrus.Printf("Listening on %s", addr)
	logrus.Panic(http.ListenAndServe(addr, corsWrapper.Handler(r))) // the syntax is a bit confusing here but this will run forever until it fails. Nothing to worry about.
}
