package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var port = "0.0.0.0:" + os.Getenv("PORT")
var numBaseStr = os.Getenv("PEOPLEBASE")
var numIntervalStr = os.Getenv("PEOPLEINTERVAL")
var powerBaseStr = os.Getenv("POWERBASE")
var powerIntervalStr = os.Getenv("POWERINTERVAL")
var buildingLongitude = os.Getenv("LONGITUDE")
var buildingLatitude = os.Getenv("LATITUDE")
var buildingHeight = os.Getenv("HEIGHT")
var numBase, powerBase, numInterval, powerInterval int
var err error

func init() {
	numBase, err = strconv.Atoi(numBaseStr)
	if err != nil {
		log.Println("Unsupply num Base")
		os.Exit(-1)
	}

	powerBase, err = strconv.Atoi(powerBaseStr)
	if err != nil {
		log.Println("Unsupply power Base")
		os.Exit(-1)
	}

	numInterval, err = strconv.Atoi(numIntervalStr)
	if err != nil {
		log.Println("Unsupply num Base")
		os.Exit(-1)
	}

	powerInterval, err = strconv.Atoi(powerIntervalStr)
	if err != nil {
		log.Println("Unsupply power Base")
		os.Exit(-1)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/num_people", GetNumberofPeople)
	mux.HandleFunc("/cur_power", GetCurrentPower)
	mux.HandleFunc("/alive", Empty)
	mux.HandleFunc("/coordinate", GetCoordinate)
	mux.HandleFunc("/height", GetHeight)
	log.Println("Listening at " + port)
	http.ListenAndServe(port, mux)
}

func Empty(w http.ResponseWriter, r *http.Request) {
	log.Println("handler empty")
	fmt.Fprintf(w, "alive")
	w.WriteHeader(http.StatusOK)
}

func GetNumberofPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("handler num_people")
	fmt.Fprintf(w, "%d人", rand.Intn(numInterval)+numBase)
	w.WriteHeader(http.StatusOK)
}

func GetCurrentPower(w http.ResponseWriter, r *http.Request) {
	log.Println("handler cur_power")
	fmt.Fprintf(w, "%dkW", rand.Intn(powerInterval)+powerBase)
	w.WriteHeader(http.StatusOK)
}

func GetCoordinate(w http.ResponseWriter, r *http.Request) {
	log.Println("handler coordinate")
	fmt.Fprintf(w, "Longitude: %v°, Latitude: %v°", buildingLongitude, buildingLatitude)
	w.WriteHeader(http.StatusOK)
}

func GetHeight(w http.ResponseWriter, r *http.Request) {
	log.Println("handler height")
	fmt.Fprintf(w, "%vm", buildingHeight)
	w.WriteHeader(http.StatusOK)
}
