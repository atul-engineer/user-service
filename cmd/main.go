package main

import (
	"database/sql"
	"log"
	"net/http"
	"user-service/internal/db"
	"user-service/internal/handlers"
)

func main() {
	sqlDb, err := db.InitDb()
	if err != nil {
		log.Fatalln("error in database connection ", err.Error())
	}
	defer sqlDb.Close()

	srv := initServer(sqlDb)
	log.Println("running server at port ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("error listening: ", srv.Addr)
	}
}

func initServer(db *sql.DB) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handlers.CreateUser(db))

	server := http.Server{
		Addr: ":8000",
		Handler: mux,
	}
	return &server
}