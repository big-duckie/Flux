package routes

import (
	"fmt"
	"net/http"

	"github.com/foolin/goview"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	err := goview.Render(rw, http.StatusOK, "home", goview.M{
		"PageTitle": "Home",
	})
	if err != nil {
		fmt.Println(err)
	}
}
