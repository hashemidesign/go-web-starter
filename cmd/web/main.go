package main

import (
	"fmt"
	"log"
	"github.com/hashemidesign/go-web-starter/pkg/config"
	"github.com/hashemidesign/go-web-starter/pkg/handlers"
	"github.com/hashemidesign/go-web-starter/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	// session management
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // persist after quiting browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	log.Println(fmt.Sprintf("Server starts on http://127.0.0.1%s...", PORT))
	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
