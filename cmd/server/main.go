package main

import (
	"log"
	"net/http"

	strategy "github.com/czas/gosip/auth/saml"
	"github.com/czas/gosip/cmd/server/handlers"
)

func main() {

	auth := &strategy.AuthCnfg{}
	err := auth.ReadConfig("./config/private.spo-user.json")
	if err != nil {
		log.Fatalf("unable to get config: %v", err)
	}

	http.HandleFunc("/digest", handlers.GetDigest(auth))
	http.HandleFunc("/web", handlers.GetWeb(auth))
	http.HandleFunc("/file", handlers.GetFile(auth))
	http.HandleFunc("/", handlers.Proxy(auth))

	log.Fatal(http.ListenAndServe(":8081", nil))

}
