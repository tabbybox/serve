package main

import (
	// "flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	showVersion := false
	var port uint16 = 3000
	// var listenUri []string
	noPortSwitching := false
	showHelp := false

	rootCmd := &cobra.Command{
		Use:   "serve [folder_name]",
		Short: "Serve files from a directory as HTTP server",
		Long:  "A simple HTTP server that serves files from a directory, fast, simple, without the weird npx/js quirks.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				fmt.Println("0.2.0")
				return
			}
			if len(args) == 0 || showHelp {
				fmt.Println("Please specify a directory to serve")
				cmd.Help()
				return
			}

			// the serving stuff
			inPath := args[0]
			abspath := ""
			if path.IsAbs(inPath) {
				fmt.Print("absolute path it is")
				abspath = inPath
			} else {
				wd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				abspath, err = filepath.Abs(path.Join(wd, args[0]))
			}

			fmt.Println(abspath)
			server := http.FileServer(http.Dir(abspath))
			http.Handle("/", server)
			log.Println(fmt.Sprintf("Listening on :%d", port))
			finalUri := fmt.Sprintf(":%d", port)
			// fmt.Println(finalUri)
			err := http.ListenAndServe(finalUri, nil)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version number")
	rootCmd.Flags().Uint16VarP(&port, "port", "p", 3000, "Port to listen on")
	// rootCmd.Flags().StringArrayVarP(&listenUri, "url", "u", []string{}, "URI endpoint on which to listen, can input moar than one")
	rootCmd.Flags().BoolVar(&noPortSwitching, "no-port-switching", false, "Print the version number")
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Print the help menu")

	rootCmd.Execute()

	// args := os.Args[1:]

	// if len(args) == 0 {
	// 	log.Fatal("Please specify a directory to serve")
	// }
	// portPtr := flag.String("l", ":3000", "listen port")
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(wd)
	// fs := http.FileServer(http.Dir(path.Join(wd, args[0])))
	// http.Handle("/", fs)
	// flag.Parse()
	// log.Print("Listening on " + *portPtr)

	// err = http.ListenAndServe(*portPtr, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
