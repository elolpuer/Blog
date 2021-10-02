package usercontroller

import (
	"database/sql"
	"log"

	dbPackage "github.com/elolpuer/blog/pkg/db"

	"github.com/elolpuer/blog/pkg/models"
	"github.com/elolpuer/blog/pkg/service/auth"
	"github.com/elolpuer/blog/pkg/service/user"
	tml2 "github.com/elolpuer/blog/pkg/tml"
	"github.com/gin-gonic/gin"
)

var tml = tml2.GetTemplates()
var db *sql.DB

func init() {
	database, err := dbPackage.ConnectionToDB()
	if err != nil {
		log.Fatal(err)
	}

	db = database
}

//UserGet ...
func UserGet(ctx *gin.Context) {
	user, err := auth.GetSessionUser(ctx.Request, "session-name")
	if err != nil {
		log.Fatal(err)
	}
	tml.ExecuteTemplate(ctx.Writer, "user.gohtml", struct {
		Title string
		H1    string
		User  *models.SessionUser
	}{
		Title: "Me",
		H1:    "Me",
		User:  user,
	})
}

//UsersGet ...
func UsersGet(ctx *gin.Context) {
	var sessionUser bool
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		sessionUser = true
	}
	users, err := user.GetAll(db)
	if err != nil {
		log.Fatal(err)
	}
	tml.ExecuteTemplate(ctx.Writer, "users.gohtml", struct {
		Title string
		H1    string
		User  bool
		Users []*models.SessionUser
	}{
		Title: "Users",
		H1:    "Users",
		User:  sessionUser,
		Users: users,
	})
}
