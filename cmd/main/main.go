package main

import (
	"log"
	"net/http"
	"time"

	"github.com/QuanNguyenDong/bookstore/pkg/middleware"
	"github.com/QuanNguyenDong/bookstore/pkg/ratelimiter"
	"github.com/QuanNguyenDong/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	router := mux.NewRouter()
	limiter := ratelimiter.NewFixedWindowLimiter(1, time.Second)
	routes.RegisterBookStoreRoutes(router.PathPrefix("/v1").Subrouter())
	
	router.Use(limiter.Middleware)
	router.Use(middleware.LoggingMiddleware)
	
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}