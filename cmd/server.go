package cmd

import (
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/api"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// serverCmd represents the api command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		router := mux.NewRouter()

		router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			// an example API handler
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		})

		apiHandler := api.NewApiHandler()
		defer apiHandler.Close()

		apiRouter := router.PathPrefix("/api").Subrouter()

		apiRouter.Methods("GET").Path("/leds").HandlerFunc(apiHandler.GetLeds)
		apiRouter.Methods("GET").Path("/animations").HandlerFunc(apiHandler.GetAnimations)
		apiRouter.Methods("POST").Path("/animations/play").HandlerFunc(apiHandler.SetAnimation)
		apiRouter.Methods("GET").Path("/statics").HandlerFunc(apiHandler.GetStatics)
		apiRouter.Methods("POST").Path("/statics/set").HandlerFunc(apiHandler.SetStatic)

		// Handle FE Assets
		spa := api.NewSpaHandler()
		router.PathPrefix("/").Handler(spa)

		srv := &http.Server{
			Handler:      router,
			Addr:         ":8080",
			WriteTimeout: 2 * time.Second,
			ReadTimeout:  2 * time.Second,
		}

		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
