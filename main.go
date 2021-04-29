package main

import (
	"fmt"
	"net/http"

	"github.com/marcustut/go-pwa/pkg/constants"
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
		BackgroundColor: constants.DefaultColor().Gray900,
		ThemeColor:      constants.DefaultColor().Gray900,
		Name:            "COVID-19 Tracker",
		Title:           "COVID-19 Tracker",
		Description:     "Track COVID-19 cases worldwide",
		LoadingLabel:    "Track COVID-19 cases worldwide.",
		Styles: []string{
			"https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css",
			"/web/styles.css",
		},
	})

	opts := options{Port: 8001}

	app.Logf("%s", logs.New("Starting COVID-19 tracker server").Tag("port", opts.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%v", opts.Port), nil); err != nil {
		panic(err)
	}
}
