package main

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type appHandler struct {
	t *Tree
	h func(*Tree, http.ResponseWriter, *http.Request)
}

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.h(a.t, w, r)
}

func indexHandler(treeData *Tree, w http.ResponseWriter, r *http.Request) {
	log.Info("indexHandler")

	fmt.Fprintf(w, "indexHandler \n")
}

func searchHandler(treeData *Tree, w http.ResponseWriter, r *http.Request) {
	log.Info("searchHandler")

	log.Info("Value from GET request")
	val, err := strconv.Atoi(r.URL.Query().Get("val"))
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Infof("GET value: <%v>", val)

	log.Infof("Search in Tree value: <%v>", val)
	res, err := treeData.search(val)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v\n", res)
}

func insertHandler(treeData *Tree, w http.ResponseWriter, r *http.Request) {
	log.Info("insertHandler")

	log.Info("Parse POST request")
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Info("Value from POST request")
	val, err := strconv.Atoi(r.Form.Get("val"))
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Infof("POST value: <%v>", val)

	log.Infof("Insert in Tree value: <%v>", val)
	err = treeData.insert(val)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "OK \n")
}

func deleteHandler(treeData *Tree, w http.ResponseWriter, r *http.Request) {
	log.Info("deleteHandler")

	log.Info("Value from DELETE request")
	val, err := strconv.Atoi(r.URL.Query().Get("val"))
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Infof("DELETE value: <%v>", val)

	log.Infof("Delete in Tree value: <%v>", val)
	res, err := treeData.delete(val, nil)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v\n", res)
}
