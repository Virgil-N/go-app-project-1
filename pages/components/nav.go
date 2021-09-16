/*
 * Created Date: 2021-09-10 04:06:49
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-16 11:09:54
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
	return app.Nav().Class("nav bg-white").Body(
		app.Div().Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8").Body(
			app.Div().Class("relative flex items-center justify-between h-16").Body(
				app.Div().Class("flex-1 flex items-center justify-center sm:items-stretch").Body(
					app.Div().Class("hidden sm:block sm:ml-6").Body(
						app.Div().Class("flex space-x-4").Body(
							app.Range(routes.Routes).Slice(func(i int) app.UI {
								if !routes.Routes[i].Meta.Hide && !routes.Routes[i].Meta.RequireLogin {
									return app.A().Class("text-xl text-black font-bold bg-transparent hover:text-purple-600 focus:text-purple-600 px-3 py-2").Href("#").Text(routes.Routes[i].Name)
								} else {
									return nil
								}
							}),
						),
					),
				),
			),
		),
	)
}
