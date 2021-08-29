/*
 * Created Date: 2021-04-15 09:56:28
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-08-29 07:52:03
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Text("header!")
}
