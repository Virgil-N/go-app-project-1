/*
 * Created Date: 2021-04-15 09:56:28
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-04-15 09:57:23
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package pages

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type header struct {
	app.Compo
}

func (h *header) Render() app.UI {
	return app.Text("header!")
}
