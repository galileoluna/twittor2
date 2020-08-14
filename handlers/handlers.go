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
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")




	
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+ PORT, handlers))
	
}