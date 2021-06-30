package main

import (
	"fmt"
	"net/http"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/logs"
)

type options struct {
	Port int `env:"PORT" help:"The port used to run the server."`
}

func main() {
	app.Route("/", &App{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Author:          "Marcus Lee",
		BackgroundColor: "#24283b",
		ThemeColor:      "#24283b",
		Name:            "Marcus Lee",
		Title:           "Marcus Lee",
		Description:     "A personal website",
		LoadingLabel:    "Welcome to my site!",
		Styles: []string{
			"/web/tailwind.css",
		},
	})

	opts := options{Port: 8001}

	app.Logf("%s", logs.New("Starting server").Tag("port", opts.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%v", opts.Port), nil); err != nil {
		panic(err)
	}
}

