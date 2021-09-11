/*
 * Created Date: 2021-08-29 07:48:26
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-11 07:17:14
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package pages

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Login struct {
	app.Compo
	name string
}

func (l *Login) Render() app.UI {

	return app.Div().Body(
		app.Nav().Class("nav").Body(),
		app.Div().Class("main").Body(
			app.Aside().Class("aside").Body(),
			app.Section().Class("section").Body(
				app.Input().
					Value(l.name). // The name field used as current input value
					OnInput(l.OnInputChange),
				app.Br(),
				app.Text(l.name),
				app.Br(),
				app.Button().Body(
					app.Text("go"),
				).OnClick(l.Go)),
		),
	)
}

func (l *Login) OnInputChange(ctx app.Context, e app.Event) {
	l.name = ctx.JSSrc().Get("value").String()
	l.ValueTo(&l.name)
	l.Update()
}

func (l *Login) Go(ctx app.Context, e app.Event) {
	ctx.Navigate("/home")
}
