package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"github.com/QuanNguyenDong/bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router.PathPrefix("/v1").Subrouter())
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}