package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type App struct {
	app.Compo
}

func (a *App) Render() app.UI {
  return app.Div().Class("flex flex-col justify-center py-32 px-8 items-center font-mono text-primary").Body(
    app.H1().Class("font-sans text-orange text-4xl").Text("Ask Me Anything"),
    app.P().Class("mt-4").Text("Just for fun! Questions will be visible after I’ve answered."),
  )
}
