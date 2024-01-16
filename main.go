package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Please specify a directory to serve")
	}
	portPtr := flag.String("l", ":3000", "listen port")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir(path.Join(wd, args[0])))
	http.Handle("/", fs)
	flag.Parse()
	log.Print("Listening on " + *portPtr)

	err = http.ListenAndServe(*portPtr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
