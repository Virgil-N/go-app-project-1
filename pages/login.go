/*
 * Created Date: 2021-08-29 07:48:26
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 03:17:25
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
}

func (h *Login) Render() app.UI {
	return app.Html().Body(
		app.Head().Body(
			app.Title().Body(
				app.Text("Login"),
			),
		),
		app.Body().Body(
			app.P().Text("llll"),
		),
	)
}
