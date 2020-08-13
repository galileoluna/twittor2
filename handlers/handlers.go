package handlers

import(
	"log"
	"net/http"
	"os"

	"github.com/galileoluna/twittor2/middlew"
	"github.com/galileoluna/twittor2/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

)

func Manejadores () {
	
	
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")



	
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+ PORT, handlers))
	
}