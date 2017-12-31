package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type jsonObj struct {
	Name string
	Age  int
	Type string
	Text string
}

func Users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	var err error
	var data []byte
	var user jsonObj

	// Change the
	keys, ok := r.URL.Query()["diff"]
	var ageDiff int
	var nameDiff string

	if !ok || len(keys) < 1 {
	} else {
		ageDiff = 10
		nameDiff = "man"
	}

	// Todo: use slices instead?
	switch ID {

	case "0":
		user = jsonObj{"Bjorn" + nameDiff, (33 + ageDiff), "Elf", "Locked out"}
	case "1":
		user = jsonObj{"No one" + nameDiff, (109 + ageDiff), "God", "Not feeling well"}
	case "2":
		user = jsonObj{"Someone" + nameDiff, (10 + ageDiff), "Orc", "Food poisoning"}
	default:
		user = jsonObj{"Magnus" + nameDiff, (30 + ageDiff), "Human", "Static as it gets"}
	}

	data, err = json.MarshalIndent(user, "", "  ")
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func StartHTTPServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/json", JSONPage)
	router.HandleFunc("/user/{ID}", Users)
	router.HandleFunc("/todos/{todoID}", TodoShow)

	srv := &http.Server{Handler: router, Addr: ":8080"}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("HTTPServer: ListenAndServe() error: %s", err)
		}
	}()

	return srv
}

func RandomText(texts []string) string {
	fmt.Printf("%+v\n", time.Now().Unix())
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(texts)
	return texts[n]
}

func JSONPage(w http.ResponseWriter, r *http.Request) {
	var err error
	var data []byte

	var jsonBlob jsonObj

	// TODO: nicer way to write?
	if _, exists := r.URL.Query()["random"]; exists {
		jsonBlob = jsonObj{
			RandomText([]string{"Magnus", "Bjorn", "Someone", "No one"}),
			30,
			"Human",
			RandomText([]string{"Locked out", "Pipes broke", "Food poisoning", "Not feeling well"})}
	} else {
		jsonBlob = jsonObj{"Magnus", 30, "Human", "Static as it gets"}
	}

	// Check if prettyprint parameter is set
	// TODO: nicer way to write this?
	if _, exists := r.URL.Query()["pretty"]; exists {
		data, err = json.MarshalIndent(jsonBlob, "", "  ")
	} else if _, exists := r.URL.Query()["prettyprint"]; exists {
		data, err = json.MarshalIndent(jsonBlob, "", "  ")
	} else {
		data, err = json.Marshal(jsonBlob)
	}

	if err != nil {
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! JSON objects can be found here: /json")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
