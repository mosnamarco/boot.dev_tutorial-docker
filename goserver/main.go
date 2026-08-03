package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	port, exists := os.LookupEnv("PORT")
	if exists {
		fmt.Println("port env detected. setting port to: ", port)
	} else {
		port = "8010"
		fmt.Println("no env set port detected. defaulting to: ", port)
	}

	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `
	<html>
		<head></head>
		<body>
			<p>Hello from Docker! I'm a Go server. </p>
		</body>
	</html>`

	w.Write([]byte(page))
}
