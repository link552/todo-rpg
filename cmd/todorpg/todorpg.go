package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"todorpg/internal/web"
	"todorpg/internal/storage"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	storage.Init();
	defer storage.Deinit();

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	http.HandleFunc("GET /", web.GetIndex)
	http.HandleFunc("POST /task", web.PostTask)
	http.HandleFunc("PUT /task", web.PutTask)
	http.HandleFunc("DELETE /task", web.DeleteTask)
	http.HandleFunc("POST /complete-task", web.PostCompleteTask)

	fmt.Println("Server starting on port", httpPort)

	log.Fatal(http.ListenAndServe("localhost:" + httpPort, nil))
}
