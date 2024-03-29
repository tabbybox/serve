package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
	// fiberlog "github.com/gofiber/fiber/v2/log"
)

var Version = "0.2.1"

func main() {
	showVersion := false
	var port uint16 = 3000
	// var listenUri []string
	// noPortSwitching := false
	showHelp := false
	// beQuiet := false
	allowBrowse := false
	enableCors := false
	var sslCert string
	var sslKey string

	rootCmd := &cobra.Command{
		Use:   "serve [folder_name]",
		Short: "Serve files from a directory as HTTP server",
		Long:  "A simple HTTP server that serves files from a directory, fast, simple, without the weird npx/js quirks.",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				fmt.Println(Version)
				return
			}
			if len(args) == 0 || showHelp {
				fmt.Println("Please specify a directory to serve")
				cmd.Help()
				return
			}

			// the serving stuff
			app := fiber.New(fiber.Config{
				AppName:               "Serve",
				DisableStartupMessage: true,
			})

			if enableCors {
				app.Use(cors.New(cors.Config{
					AllowOrigins: "*",
				}))
			}

			app.Static("/", args[0], fiber.Static{
				Browse: allowBrowse,
			})

			finalUri := fmt.Sprintf(":%d", port)
			log.Println(fmt.Sprintf("Listening on https://127.0.0.1%s", finalUri))

			if (sslCert != "") && (sslKey != "") {
				log.Fatal(app.ListenTLS(finalUri, sslCert, sslKey))
			} else {
				log.Fatal(app.Listen(finalUri))
			}
		},
	}
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version number")
	rootCmd.Flags().Uint16VarP(&port, "port", "p", 3000, "Port to listen on")
	// rootCmd.Flags().StringArrayVarP(&listenUri, "url", "u", []string{}, "URI endpoint on which to listen, can input moar than one")
	// rootCmd.Flags().BoolVar(&noPortSwitching, "no-port-switching", false, "[INOP] Don't switch port when selected port is already in use")
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Print the help menu")
	// rootCmd.Flags().BoolVarP(&beQuiet, "quiet", "q", false, "Supress log messages")
	rootCmd.Flags().BoolVarP(&allowBrowse, "allow-browse", "b", false, "Allow browsing of the directory")
	rootCmd.Flags().BoolVarP(&enableCors, "enable-cors", "c", false, "Enable CORS, sets Access-Control-Allow-Origin to *")
	rootCmd.Flags().StringVar(&sslCert, "ssl-cert", "", "Path to SSL certificate")
	rootCmd.Flags().StringVar(&sslKey, "ssl-key", "", "Path to SSL key")

	rootCmd.Execute()
}
