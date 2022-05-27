package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Base(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(
		map[string]string{
			"api":     strings.ToLower(Context.App.Name),
			"version": Context.App.Version,
		},
	)
}
