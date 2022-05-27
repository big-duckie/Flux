package api

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var Context *cli.Context

func Init(c *cli.Context, router *mux.Router) error {
	Context = c

	subRouter := router.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("", Base).Methods("GET")
	subRouter.HandleFunc("/", Base).Methods("GET")

	// Authentication
	subRouter.HandleFunc("/auth", Auth).Methods("POST")

	// Passing an empty JSON to any update endpoint deletes that resource.

	// Modpack
	subRouter.HandleFunc("/modpack", ModpackList).Methods("GET")    // Lists modpacks
	subRouter.HandleFunc("/modpack", ModpackCreate).Methods("POST") // Creates modpack (Requires Auth)

	subRouter.HandleFunc("/modpack/{modpack}", Modpack).Methods("GET")        // Gets info about specific modpack
	subRouter.HandleFunc("/modpack/{modpack}", ModpackUpdate).Methods("POST") // Updates info about specific modpack (Requires Auth)

	subRouter.HandleFunc("/modpack/{modpack}/", BuildList).Methods("GET")    // Lists builds
	subRouter.HandleFunc("/modpack/{modpack}/", BuildCreate).Methods("POST") // Creates build (Requires Auth)

	subRouter.HandleFunc("/modpack/{modpack}/{build}", Build).Methods("GET")       // Gets build info
	subRouter.HandleFunc("/modpack/{modpack}/{build}", BuildUpdate).Methods("GET") // Updates build info (Requires Auth)

	// Mod
	subRouter.HandleFunc("/mod", ModList).Methods("GET")    // Lists mods
	subRouter.HandleFunc("/mod", ModCreate).Methods("POST") // Creates mod (Requires Auth)

	subRouter.HandleFunc("/mod/{mod}", Mod).Methods("GET")        // Gets mod info
	subRouter.HandleFunc("/mod/{mod}", ModUpdate).Methods("POST") // Updates mod info (Requires Auth)

	subRouter.HandleFunc("/mod/{mod}/", VersionList).Methods("GET")    // Lists versions
	subRouter.HandleFunc("/mod/{mod}/", VersionCreate).Methods("POST") // Creates version (Requires Auth)

	subRouter.HandleFunc("/mod/{mod}/{version}", Version).Methods("GET")        // Gets version info
	subRouter.HandleFunc("/mod/{mod}/{version}", VersionUpdate).Methods("POST") // Updates version info (Requires Auth)

	// User (Requires Auth)
	subRouter.HandleFunc("/user", UserList).Methods("GET")    // Lists users
	subRouter.HandleFunc("/user", UserCreate).Methods("POST") // Creates user

	subRouter.HandleFunc("/user/{username}", User).Methods("GET")        // Gets user info
	subRouter.HandleFunc("/user/{username}", UserUpdate).Methods("POST") // Updates user info

	// Client (Requires Auth)
	subRouter.HandleFunc("/client", ClientList).Methods("GET")    // Lists clients
	subRouter.HandleFunc("/client", ClientCreate).Methods("POST") // Creates client

	subRouter.HandleFunc("/client/{name}", Client).Methods("GET")        // Gets client info
	subRouter.HandleFunc("/client/{name}", ClientUpdate).Methods("POST") // Updates client info

	// API Key (Requires Auth)
	subRouter.HandleFunc("/key", KeyList).Methods("GET")    // List keys
	subRouter.HandleFunc("/key", KeyCreate).Methods("POST") // Create key

	subRouter.HandleFunc("/key/{name}", Key).Methods("GET")        // Gets key info
	subRouter.HandleFunc("/key/{name}", KeyUpdate).Methods("POST") // Updates key info

	return nil
}
