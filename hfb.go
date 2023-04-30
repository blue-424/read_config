package main

import (
	Cnf "awesomeProject/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var configuration Cnf.Configuration

func init() {
	println()
	file, err := os.Open("cfg/hfb.cfg")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		log.Print(err.Error())
	}
	decoder := json.NewDecoder(file)
	errd := decoder.Decode(&configuration)
	if errd != nil {
		fmt.Println("error:", errd)
	}
	fmt.Println("config loaded")
}

//ToDo iwant to add something
func main() {

	//_, err := zap.NewProduction()
	//if err != nil {
	//	log.Print(err.Error())
	//}
	println(configuration.Username)
	println(configuration.Password)
	println(configuration.Flag)
	for i := range configuration.Products {
		println(configuration.Products[i])

	}
	println(configuration.Adr.ToString())
	println(configuration.Birth)
	///
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello!")
		if err != nil {
			return
		}
	})

	fmt.Printf("Hello Hossein !!! \n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
