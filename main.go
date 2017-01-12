package main

import (
	"crypto/rand"
	"fmt"
	"github.com/MaxKobyakov/go-blog/models"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
	"net/http"
)

var posts map[string]*models.Post
var counter int

func indexHandler(rnd render.Render) {
	fmt.Println(counter)

	rnd.HTML(200, "index", posts)

}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	post, found := posts[id]
	if !found {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "write", post)
}

func savePostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
<<<<<<< HEAD
	contentHtml := r.FormValue("content")

	contentMarkdown := string(blackfriday.MarkdownBasic([]byte(contentHtml)))
=======
	content := r.FormValue("content")
>>>>>>> ea53ff83f9f3c861201f592567797b186b11c3d5

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
<<<<<<< HEAD
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, contentHtml, contentMarkdown)
=======
		post.Content = content
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, content)
>>>>>>> ea53ff83f9f3c861201f592567797b186b11c3d5
		posts[post.Id] = post

	}

	rnd.Redirect("/")

}

func deleteHandler(rnd render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	delete(posts, id)

	rnd.Redirect("/")

}

func getHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))

	rnd.JSON(200, map[string]interface{}{"html": string(htmlBytes)})
}

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func main() {
	fmt.Println("Слушаем порт: 3000")

	posts = make(map[string]*models.Post, 0)

	counter = 0

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:           []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
		IndentXML:  true,    // Output human readable XML
		//HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
	}))

	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))
	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Get("/edit/:id", editHandler)
	m.Get("/DeletePost/:id", deleteHandler)
	m.Post("/SavePost", savePostHandler)
	m.Post("/gethtml", getHtmlHandler)

	m.Run()
}
