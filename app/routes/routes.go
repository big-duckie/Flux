package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var Context *cli.Context

func Init(c *cli.Context, router *mux.Router) error {
	Context = c

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/home", Home).Methods("GET")

	return nil
}
