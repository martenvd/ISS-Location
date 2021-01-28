package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template
var currentPosition ISSData

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	currentPosition = getCurrentISSLocation()
	tpl.ExecuteTemplate(w, "index.html", currentPosition)
}

func main() {
	log.Println("Listening on :80...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

// ISSData ...
type ISSData struct {
	Message     string `json:"message"`
	Timestamp   int64  `json:"timestamp"`
	IssPosition struct {
		Latitude  float64 `json:"latitude,string"`
		Longitude float64 `json:"longitude,string"`
	} `json:"iss_position"`
}

func getAddress(lat float64, lon float64) string {
	convLat := fmt.Sprintf("%f", lat)
	convLon := fmt.Sprintf("%f", lon)
	key := "*************" // Fill in your own opencagedata API key
	url := "https://api.opencagedata.com/geocode/v1/json?q=" + convLat + "+" + convLon + "&key=" + key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	return string(byteValue)
}

func getCurrentISSLocation() ISSData {
	resp, err := http.Get("http://api.open-notify.org/iss-now.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var currentPosition ISSData
	byteValue, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(byteValue, &currentPosition)

	return currentPosition
}
