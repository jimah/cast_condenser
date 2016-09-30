package main

import (
	"fmt"

	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hey!\n")
}

func main() {
	// TODO environment setup
	logrus.SetLevel(logrus.DebugLevel)

	port := "1337"
	host := "0.0.0.0"
	addr := fmt.Sprintf("%s:%s", host, port)

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/search", searchHandler)

	logrus.WithFields(logrus.Fields{
		"port": port,
		"host": host,
	}).Info("httprouter listening")

	logrus.Fatal(http.ListenAndServe(
		addr,
		router,
	))
}