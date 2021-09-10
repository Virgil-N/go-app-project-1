/*
 * Created Date: 2021-04-15 09:56:28
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 03:17:15
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Text("header!")
}
