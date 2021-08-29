/*
 * Created Date: 2021-04-15 09:51:14
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-08-29 07:52:14
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package pages

import (
	"github.com/Virgil-N/go-app-project-1/pages/components"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.P().Body(
		app.Text("Home "),
		&components.Header{},
	)
}
