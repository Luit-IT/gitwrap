package main

import (
	"flag"
	"http"
	"http/cgi"
	"log"
	"os"
	"strconv"
)

var (
	cgitPath = flag.String("cgit", "/var/www/htdocs/cgit/cgit",
		"Location of cgit executable.")
	cgitRes = flag.String("cgitres", "/var/www/htdocs/cgit",
		"Location of cgit resources.")
	config = flag.String("config", "",
		"Location of the cgit configuration file.")

	port = flag.Int("port", 5000, 
		"TCP port to listen on.")
	addr = flag.String("addr", "127.0.0.1",
		"Address to listen on.")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	cgiHandler := &cgi.Handler{
		Path: *cgitPath,
		Env:  []string{},
		InheritEnv: []string{
			"CGIT_CONFIG",
		},
	}
	if *config != "" {
		cgiHandler.Env = append(cgiHandler.Env,
			"CGIT_CONFIG="+*config)
	}

	fs := http.FileServer(http.Dir(*cgitRes))
	http.Handle("/cgit.css", fs)
	http.Handle("/cgit.png", fs)
	http.Handle("/", cgiHandler)
 
	err := http.ListenAndServe(*addr + ":" + strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Everything seems to work: daemonize (close file handles)
	os.Stdin.Close()
	os.Stdout.Close()
	os.Stderr.Close()
}
