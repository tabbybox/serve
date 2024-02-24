package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/gofiber/fiber/v2"
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
				AppName: "Serve",
				DisableStartupMessage: true,
			});


			app.Static("/", args[0], fiber.Static{
				Browse: allowBrowse,
			})

			finalUri := fmt.Sprintf(":%d", port)
			log.Println(fmt.Sprintf("Listening on https://127.0.0.1%s", finalUri))

			err := app.Listen(finalUri)
			if err != nil {
				log.Fatal(err)
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

	rootCmd.Execute()
}
