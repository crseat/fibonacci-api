//Package app provides logic for starting the server and wiring together the user, domain, and server side
package app

import (
	"context"
	"fibonacci-api/domain"
	"fibonacci-api/logger"
	"fibonacci-api/service"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

//Start the server on the given port.
func Start(port string) {
	logger.InfoLogger.Println(fmt.Sprintf("Starting server on localhost:%s ...", port))

	//wiring
	fibRepository := domain.NewFibRepository()
	Handler := fibHandler{
		fibService: service.NewFibonacciService(fibRepository),
	}

	//Create channel to monitor whether a shutdown has been initiated
	quit := make(chan bool)

	//Create Router object
	wg := &sync.WaitGroup{}
	router := &Router{
		Handler:   &Handler,
		WaitGroup: wg,
		quitChan:  &quit,
	}

	//Construct listen address and then server
	address := "localhost" + ":" + port
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ErrorLog:     logger.ErrorLogger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	//source: https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7
	go func() {
		<-quit
		logger.InfoLogger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		log.Println("waiting for running jobs to finish")
		wg.Wait()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.ErrorLogger.Fatal("Could not gracefully shutdown the server: ", err)
		}
	}()

	//Start http server
	//address := "localhost"
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.ErrorLogger.Fatal("Error starting server")
	}

	//<-done
	logger.InfoLogger.Println("Server stopped")
}
