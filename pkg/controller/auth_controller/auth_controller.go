package authcontroller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbPackage "github.com/elolpuer/blog/pkg/db"

	"github.com/elolpuer/blog/pkg/models"
	"github.com/elolpuer/blog/pkg/service/auth"
	tml2 "github.com/elolpuer/blog/pkg/tml"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

//SignInGet ...
func SignInGet(ctx *gin.Context) {
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		ctx.Redirect(http.StatusSeeOther, "/posts")
	}
	user := false
	tml.ExecuteTemplate(ctx.Writer, "signin.gohtml", struct {
		Title string
		H1    string
		User  bool
	}{
		Title: "Sign In",
		H1:    "Sign In",
		User:  user,
	})
}

//SignUpGet ...
func SignUpGet(ctx *gin.Context) {
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		ctx.Redirect(http.StatusSeeOther, "/posts")
	}
	user := false
	tml.ExecuteTemplate(ctx.Writer, "signup.gohtml", struct {
		Title string
		H1    string
		User  bool
	}{
		Title: "Sign Up",
		H1:    "Sign Up",
		User:  user,
	})
}

//SignUpPost ...
func SignUpPost(ctx *gin.Context) {
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	}
	var NewUser = new(models.User)
	NewUser.Username = ctx.PostForm("username")
	NewUser.Email = ctx.PostForm("email")
	NewUser.Password = ctx.PostForm("password")
	err := auth.SignUp(db, NewUser)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/")
}

//SignInPost ...
func SignInPost(ctx *gin.Context) {
	fmt.Println(db)
	var signUser = new(models.User)
	signUser.Email = ctx.PostForm("email")
	signUser.Password = ctx.PostForm("password")
	sessionUser, err := auth.SignIn(db, signUser)
	if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
		http.Error(ctx.Writer, "Invalid Data", 400)
		return
	}
	err = auth.CreateSessionUser(ctx.Writer, ctx.Request, sessionUser, "session-name")
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/posts")
}

//LogOutPost ...
func LogOutPost(ctx *gin.Context) {
	err := auth.Logout(ctx.Writer, ctx.Request, "session-name")
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/")
}
