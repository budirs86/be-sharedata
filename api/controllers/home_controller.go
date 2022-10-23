package controllers

import (
	"net/http"

	"github.com/budirs86/be-sharedata/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Share Data API")

}
