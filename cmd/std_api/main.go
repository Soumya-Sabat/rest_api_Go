package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ghost/restAPI/internal/config"
	"github.com/ghost/restAPI/internal/http/handlers/student"
)

func main() {
	//load config 
	cfg := config.MustLoad()
	//set up the logger through the package
	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students",student.New())

	//setup server
	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	// fmt.Printf("server started at port : %s",cfg.HTTPServer.Addr)
	slog.Info("server started :",slog.String("address",cfg.Addr))

	//while working in the production server then we just cannot accept the close_interrutpt signal from the terminal 
	//the servers must gracefully shutdown - we will be using the go routines and channels for syncronisation
	done:=make(chan os.Signal,1)
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM) //we are providing the list of signals that are eexpected to come , which will be passed to the done channel

	go func(){
		err:=server.ListenAndServe()
		if err!=nil{
		log.Fatal("Failed to start the server")
	}
	}()

	<-done

	//server stop logic 
	slog.Info("shutting down the server") //slog - structured log
	//we can do server.Shutdown(), 'but IRL sometimes it does not cause the real shutdown of the server annd the port remains blocked, so we use a timer such that if the server does not shutdown within that time report it 
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err!=nil{
		slog.Error("failed to shutdoen the server",slog.String("error",err.Error()))
	}
	//the other way is 
	//if err:=server.Shutdown(ctx);  err!=nil{ 
	// slog.Error("failed to shutdoen the server",slog.String("error",err.Error()))
	// }
	slog.Info("server shutdown successfully")
}	
