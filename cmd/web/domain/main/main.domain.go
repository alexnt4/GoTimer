package mainDomain

import (
	"GoTimer/internals/models"
	"GoTimer/internals/utils"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesIntance.MAIN)
}
