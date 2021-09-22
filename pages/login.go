/*
 * Created Date: 2021-08-29 07:48:26
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-22 05:32:44
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package pages

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Login struct {
	app.Compo
	Name     string
	password string
}

func (login *Login) Render() app.UI {
	return app.Div().Class("ui grid container center middle aligned").Body(
		app.Div().Class("login-wrap").Body(
			app.H2().Class("ui header").Body(
				app.Div().Class("content").Body(
					app.Text("Login"),
				),
			),
			app.Div().Class("ui large form").Body(),
			app.Div().Class("ui message").Body(
				app.Text("新用户？"),
				app.A().Href("#").Text("注册"),
			),
			app.Button().Body(
				app.Text("click me"),
			).OnClick(login.withClick),
		),
	)
}

func (login *Login) withClick(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue("click-me", "from login")
}
