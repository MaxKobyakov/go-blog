package routes

import (
	"github.com/MaxKobyakov/go-blog/routes"
	"github.com/MaxKobyakov/go-blog/session"
	"github.com/MaxKobyakov/go-blog/utils"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
	"net/http"

	"github.com/MaxKobyakov/go-blog/db/documents"
	"github.com/MaxKobyakov/go-blog/models"
)

func WriteHandler(rnd render.Render) {
	post := models.Post{}
	rnd.HTML(200, "write", post)
}

func EditHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {
	postsCollection := db.C("Posts")
	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)

	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	rnd.HTML(200, "write", post)
}

func SavePostHandler(rnd render.Render, r *http.Request, db *mgo.Database) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{id, title, contentHtml, contentMarkdown}
	postsCollection := db.C("Posts")

	if id != "" {
		postsCollection.UpdateId(id, postDocument)

	} else {
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/")
}
func DeleteHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {

	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}
	postsCollection := db.C("Posts")

	postsCollection.RemoveId(id)

	rnd.Redirect("/")

}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}
