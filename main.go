/*
 * Created Date: 2021-04-12 09:34:32
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 03:25:58
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package main

import (
	"log"
	"net/http"

	"github.com/Virgil-N/go-app-project-1/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
	name string
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Text(h.name),
		app.H1().Text("------"),
		app.Input().
			Value(h.name). // The name field used as current input value
			OnInput(h.OnInputChange),
		app.Button().Body(
			app.Text("go"),
		).OnClick(h.Leap),
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.name = ctx.JSSrc().Get("value").String()
	h.ValueTo(&h.name)
	h.Update()
	ctx.Navigate("/home")
}

func (h *hello) Leap(ctx app.Context, e app.Event) {
	ctx.Navigate("/home")
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})
	app.Route("/login", &pages.Login{})
	app.Route("/home", &pages.Home{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// On the server-side, RunWhenOnBrowser() does nothing, which allows the
	// writing of server logic without needing precompiling instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/styles/index.css", // Loads hello.css file.
		},
	})

	// err := app.GenerateStaticWebsite("./dist", &app.Handler{
	// 	Name:        "Hello",
	// 	Description: "An Hello World! example",
	// 	Resources:   app.LocalDir("./dist"),
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := http.ListenAndServe(":3400", nil); err != nil {
		log.Fatal(err)
	}
}
