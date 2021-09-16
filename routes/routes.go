/*
 * Created Date: 2021-09-10 03:47:52
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-16 10:29:52
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
			Name: "生活记录",
			Path: "/life",
			// Component:
			Meta: Meta{
				RequireLogin: false,
				Hide:         false,
				Title:        "Life",
				IconSrc:      "",
			},
		},
		{
			Name: "学习笔记",
			Path: "/study",
			Meta: Meta{
				RequireLogin: false,
				Hide:         false,
				Title:        "Study",
				IconSrc:      "",
			},
		},
		{
			Name: "这是个测试",
			Path: "/test",
			Meta: Meta{
				RequireLogin: false,
				Hide:         true,
				Title:        "Test",
				IconSrc:      "",
			},
		},
		{
			Name: "旅行日志",
			Path: "/travel",
			Meta: Meta{
				RequireLogin: false,
				Hide:         false,
				Title:        "Travel",
				IconSrc:      "",
			},
		},
	}
}
