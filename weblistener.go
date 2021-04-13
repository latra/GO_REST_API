package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

type Message struct {
	Direction string `json:"IP"`
	Port      string `json:"Port"`
	Password  string `json:"Password"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pero tu que tramas, morena")
	fmt.Println("Endpoint Hit: homePage")
}
func sendToBots(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message sended via bots")
	discord, err := discordgo.New("Bot " + "")
	fmt.Println(err)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message Message
	json.Unmarshal(reqBody, &message)
	discord.ChannelMessageSend("404371919618703371", fmt.Sprintf("Servidord de terraria abierto en la IP %s\nPuerto: %s.\nContraseña: %s", message.Direction, message.Port, message.Password))
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/send", sendToBots).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
