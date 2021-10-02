package postcontroller

import (
	"database/sql"
	"log"
	"net/http"

	dbPackage "github.com/elolpuer/blog/pkg/db"
	"github.com/elolpuer/blog/pkg/models"
	"github.com/elolpuer/blog/pkg/service/auth"
	"github.com/elolpuer/blog/pkg/service/post"
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

//PostsGet ...
func PostsGet(ctx *gin.Context) {
	var userYet bool
	if isNew, err := auth.SessionIsNew(ctx.Request, "session-name"); !isNew || err != nil {
		userYet = true
	}
	user, err := auth.GetSessionUser(ctx.Request, "session-name")
	if err != nil {
		log.Fatal(err)
	}
	posts, err := post.Get(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	tml.ExecuteTemplate(ctx.Writer, "posts.gohtml", struct {
		Title string
		H1    string
		User  bool
		Posts []*models.Post
	}{
		Title: "Posts",
		H1:    "Add",
		User:  userYet,
		Posts: posts,
	})
}

//AddPost ...
func AddPost(ctx *gin.Context) {
	user, err := auth.GetSessionUser(ctx.Request, "session-name")
	if err != nil {
		log.Fatal(err)
	}
	text := ctx.PostForm("body")
	err = post.Add(db, user.ID, text)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/posts")
}

//DeletePost ...
func DeletePost(ctx *gin.Context) {
	session, err := auth.GetSessionStore(ctx.Request, "session-name")
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
	userID := session.Values["userID"].(int)
	id := ctx.Query("id")
	err = post.Delete(db, id, userID)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/posts")
}
