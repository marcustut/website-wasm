package main

import (
	"fmt"

	"github.com/marcustut/go-pwa/pkg/constants"
	"github.com/marcustut/go-pwa/pkg/fetchers"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type App struct {
	app.Compo
}

func (a *App) Render() app.UI {
	all := fetchers.FetchAllCountries()

	my := fetchers.FetchCountry("Malaysia")

	return app.Html().
		Style("color", constants.DefaultColor().Gray100).
		Body().
		Class("w-9/12 my-0 mx-auto").
		Style("background-color", constants.DefaultColor().Gray900).
		Body(
			app.Div().
				Class("py-5").
				Body(
					app.H1().
						Class("text-5xl font-bold text-center").
						Text("COVID-19 Tracker"),
				),
			app.Div().
				Class("flex items-center w-full rounded-lg").
				Style("background-color", constants.DefaultColor().Gray700).
				Style("color", constants.DefaultColor().Gray300).
				Body(
					app.P().
						Class("h-full mx-4 flex items-center pointer-events-none").
						Text("Malaysia"),
					app.Input().
						Class("w-full py-2 px-4 rounded-lg focus:outline-none pl-0").
						Style("background-color", constants.DefaultColor().Gray700).
						AutoComplete(false).
						Placeholder("Search for country names"),
				),
			app.Div().
				Class("flex w-full rounded-lg my-5 py-2 px-4").
				Style("background-color", constants.DefaultColor().Gray700).
				Body(
					app.Div().Body(
						app.P().Text(fmt.Sprintf("Worldwide: %d cases, %d deaths, %d recovered", int(all.Cases), int(all.Deaths), int(all.Recovered))),
						app.P().Text(fmt.Sprintf("Malaysia: %d cases, %d deaths, %d recovered", int(my.Cases), int(my.Deaths), int(my.Recovered))),
					),
				),
		)
}
