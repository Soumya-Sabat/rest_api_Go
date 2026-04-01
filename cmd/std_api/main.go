package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghost/restAPI/internal/config"
)

func main() {
	//load config 
	cfg := config.MustLoad()
	//set up the logger through the package
	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Welcome to student API"))
	})

	//setup server
	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	fmt.Printf("server started at port : %s",cfg.HTTPServer.Addr)
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal("Failed to start the server")
	}
}
