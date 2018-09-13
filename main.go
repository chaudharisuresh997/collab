package main

import (
	c "control"
	domain "domain"
	"fmt"
	"log"
	"net/http"
//"github.com/gorilla/handlers"
cors "github.com/rs/cors"
)

func main() {
	
    corsObj := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4200"},
    
   })
	router := c.NewRouter()
    //handler := corsObj.Handler(router)
/*	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	})*/
	domain.CustomFun()
	
	handler := corsObj.Handler(router)

	srv := &http.Server{
	  Handler: handler,
	  Addr:    ":" + "8081",
	}
  
	log.Fatal(srv.ListenAndServe())
	//log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:4200"}))(router)))
	fmt.Println("Shiv")
}
