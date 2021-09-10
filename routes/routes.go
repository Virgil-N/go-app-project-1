/*
 * Created Date: 2021-09-10 03:47:52
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-10 04:13:15
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package routes

type Route struct {
	Name string
	Path string
	// Component app.Compo
	Meta     Meta
	Children []Route
}

type Meta struct {
	RequireLogin bool
	Hide         bool
	Title        string
	IconSrc      string
}

var Routes []Route

func init() {
	Routes = []Route{
		{
			Name: "Login",
			Path: "/login",
			// Component:
			Meta: Meta{
				RequireLogin: false,
				Hide:         false,
				Title:        "Login",
				IconSrc:      "",
			},
		},
		{
			Name: "Home",
			Path: "/home",
			Meta: Meta{
				RequireLogin: false,
				Hide:         false,
				Title:        "Home",
				IconSrc:      "",
			},
		},
	}
}
