package main

import (
	"flag"
	"http"
	"http/cgi"
)

var (
	cgitPath = flag.String("cgit", "/var/www/htdocs/cgit/cgit",
		"Location of cgit executable.")
	cgitRes = flag.String("cgitres", "/var/www/htdocs/cgit",
		"Location of cgit resources.")
	config = flag.String("config", "",
		"Location of the cgit configuration file.")
)

func main() {
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

	http.ListenAndServe("127.0.0.1:5000", nil)
}
