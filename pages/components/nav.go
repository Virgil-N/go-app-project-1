/*
 * Created Date: 2021-09-10 04:06:49
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 04:29:06
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package components

import (
	"github.com/Virgil-N/go-app-project-1/routes"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Nav struct {
	app.Compo
}

func (n *Nav) Render() app.UI {
	return app.Ul().Body(
		app.Range(routes.Routes).Slice(func(i int) app.UI {
			return app.Li().Text(routes.Routes[i].Name)
		}),
	)
}
