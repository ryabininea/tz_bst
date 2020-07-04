package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)
}

func main() {

	log.Info("Tree initialization")
	treeData, err := initTree("data.json")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.Handle("/", appHandler{treeData, indexHandler})

	r.Handle("/search", appHandler{treeData, searchHandler}).Methods(http.MethodGet)
	r.Handle("/insert", appHandler{treeData, insertHandler}).Methods(http.MethodPost)
	r.Handle("/delete", appHandler{treeData, deleteHandler}).Methods(http.MethodDelete)

	log.Info("Server is listening...")
	server := http.Server{Addr: "localhost:8181", Handler: r}
	log.Fatal(server.ListenAndServe())
}
