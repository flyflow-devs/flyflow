package main

import (
	"github.com/flyflow-devs/flyflow/internal/config"
	"github.com/flyflow-devs/flyflow/internal/logger"
	"github.com/flyflow-devs/flyflow/internal/server"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)



var rootCmd = &cobra.Command{
	Use:   "flyflow",
	Short: "FlyFlow is LLM middleware",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		s := server.NewServer(cfg)
		logger.S.Fatal(http.ListenAndServe(":"+cfg.Port, s.Router))
	},
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database operations",
}

var autoMigrateCmd = &cobra.Command{
	Use:   "automigrate",
	Short: "Automatically migrate the database schema",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		server.InitDB(cfg, true)
		logger.S.Info("Database migration completed successfully.")
	},
}

func init() {
	cfg := config.NewConfig()
	logger.InitLogger(cfg.Env)

	dbCmd.AddCommand(autoMigrateCmd)
	rootCmd.AddCommand(dbCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
