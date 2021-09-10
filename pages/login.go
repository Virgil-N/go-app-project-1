/*
 * Created Date: 2021-08-29 07:48:26
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 04:06:59
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

	return app.Div().Body(
		app.Nav().Class("nav").Body(),
		app.Div().Class("main").Body(
			app.Aside().Class("aside").Body(),
			app.Section().Class("section").Body(),
		),
	)
}
