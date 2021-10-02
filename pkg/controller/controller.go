package controller

import (
	"github.com/elolpuer/blog/pkg/service/auth"
	tml2 "github.com/elolpuer/blog/pkg/tml"
	"github.com/gin-gonic/gin"
)

var tml = tml2.GetTemplates()

//IndexGet ...
func IndexGet(ctx *gin.Context) {
	var user bool
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		user = true
	}
	tml.ExecuteTemplate(ctx.Writer, "index.gohtml", struct {
		Title string
		H1    string
		User  bool
	}{
		Title: "Index Page",
		H1:    "Index Page",
		User:  user,
	})
}
