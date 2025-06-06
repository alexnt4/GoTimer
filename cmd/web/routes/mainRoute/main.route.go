package mainRoute

import (
	"GoTimer/cmd/web/routes"
	"GoTimer/internals/models"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesIntance.MAIN, func())
}
