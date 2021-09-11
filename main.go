/*
 * Created Date: 2021-04-12 09:34:32
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-11 07:43:53
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

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &pages.Login{})
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
		Name:        "Go-App",
		Description: "Go-App!",
		Styles: []string{
			"/web/styles/index.css", // Loads hello.css file.
		},
		Scripts: []string{
			"/web/scripts/materialize.min.js",
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
