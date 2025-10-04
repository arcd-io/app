package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/arcd-io/arcd/internal/database"
	"github.com/arcd-io/arcd/internal/database/migration"
	"github.com/arcd-io/arcd/server/grpc"
	"github.com/arcd-io/arcd/server/http"
	"golang.org/x/net/http2"

	"golang.org/x/net/http2/h2c"

	"github.com/joho/godotenv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nethttp "net/http"
)

var rootCmd = &cobra.Command{
	Use:   "arcd",
	Short: "arcd starts the Arcd server",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := viper.GetString("dsn")
		database := database.New(dsn)

		migration.InitMigration(cmd.Context(), database)
		migration.AutoMigrate(cmd.Context(), database)

		// githubService := github.NewService(database)

		e := http.NewServer(database)
		grpc.NewServer(e, database)

		server := &nethttp.Server{
			Addr:    "localhost:8080",
			Handler: h2c.NewHandler(e, &http2.Server{}),
		}

		slog.Info("Server starting on localhost:8080")
		log.Fatal(server.ListenAndServe())
	},
}

func init() {
	slog.Info("initializing Fivemanage...")

	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file. Probably becasue we're in production", slog.Any("error", err))
	}

	rootCmd.Flags().String("dsn", "", "Database DSN")
	if err := viper.BindEnv("dsn", "DSN"); err != nil {
		bindError(err)
	}

	rootCmd.AddCommand(migration.RootCmd)
	migration.RootCmd.AddCommand(
		migration.InitCmd,
		migration.MigrateCmd,
		migration.CreateMigrationCmd,
		migration.UnlockCmd,
		migration.LockCmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func bindError(err error) {
	if err != nil {
		slog.Error("failed to bind env", slog.Any("error", err))
		os.Exit(1)
	}
}
