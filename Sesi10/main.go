package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"time"
)

/* make json format file
{
	"status":{
		"water" : int,
		"wind" : int
	}
}
*/

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Status struct {
	Status      Data   `json:"status"`
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}

var status Status

func main() {
	go startService()
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
	err := exec.Command("rundle32", "url.dll,FileProtocolHandler", "http://localhost:8080").Start()
	if err != nil {
		log.Fatal(err)
	}
}

func startService() {
	ticker := time.NewTicker(4 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				doService()
			}
		}
	}()

	time.Sleep(60 * time.Second)
	ticker.Stop()
	done <- true
}

func doService() {
	file, err := os.ReadFile("file.json")
	if err != nil {
		fmt.Println("Error reading file.json: ", err)
		return
	}

	err = json.Unmarshal(file, &status)
	if err != nil {
		fmt.Println("Error unmarshaling file.json: ", err)
		return
	}

	water := GenerateRandomNumber()
	wind := GenerateRandomNumber()

	status.Status.Water = water
	status.Status.Wind = wind

	waterStatus, windStatus := SetStatus(status.Status)

	status.WaterStatus = waterStatus
	status.WindStatus = windStatus

	newFile, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		fmt.Println("Error marshal data to json: ", err)
		return
	}

	err = os.WriteFile("file.json", newFile, 0644)
	if err != nil {
		fmt.Println("Error writing data to files file.json: ", err)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error parsing index.html: ", err)
		return
	}

	err = index.Execute(w, status)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return
	}

}

func SetStatus(d Data) (string, string) {
	var waterStatus string
	var windStatus string

	//waterStatus
	if d.Water >= 1 && d.Water <= 5 {
		waterStatus = "aman"
	} else if d.Water >= 6 && d.Water <= 8 {
		waterStatus = "siaga"
	} else if d.Water >= 9 && d.Water <= 100 {
		waterStatus = "bahaya"
	}

	//windStatus
	if d.Wind >= 1 && d.Wind <= 6 {
		windStatus = "aman"
	} else if d.Wind >= 7 && d.Wind <= 15 {
		windStatus = "siaga"
	} else if d.Wind >= 16 && d.Wind <= 100 {
		windStatus = "bahaya"
	}

	return waterStatus, windStatus
}

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100-1) + 1
}
