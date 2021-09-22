/*
 * Created Date: 2021-08-29 07:48:26
 * Author: Virgil-N
 * Description:
 * -----
 * Last Modified: 2021-09-22 05:34:37
 * Modified By: Virgil-N (lieut9011@126.com)
 * -----
 * Copyright (c) 2019 - 2021 ⚐
 * Virgil-N will save your soul!
 * -----
 */

package pages

import (
	"fmt"

	"github.com/Virgil-N/go-app-project-1/pages/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
	name string
}

func (home *Home) Render() app.UI {

	return app.Div().Body(
		&components.Nav{},
		&Login{},
		app.Div().Class("main").Body(
			app.Aside().Class("aside").Body(),
			app.Section().Class("section").Body(
				app.Input().
					Value(home.name). // The name field used as current input value
					OnInput(home.OnInputChange),
				app.Br(),
				app.Text(home.name),
				app.Br(),
				app.Button().Class("waves-effect waves-light btn").Body(
					app.Text("go"),
				).OnClick(home.Go)),
		),
	)
}

func (home *Home) OnInputChange(ctx app.Context, e app.Event) {
	home.name = ctx.JSSrc().Get("value").String()
	home.ValueTo(&home.name)
	home.Update()
}

func (l *Home) Go(ctx app.Context, e app.Event) {
	ctx.Navigate("/home")
}

func (home *Home) OnMount(ctx app.Context) {
	fmt.Println("mounted")
	ctx.Handle("click-me", home.handleMsg)
}

func (home *Home) handleMsg(ctx app.Context, a app.Action) {
	fmt.Println(a.Value)
	if msg, ok := a.Value.(string); ok {
		fmt.Println("get msg: ", msg)
	} else {
		fmt.Println("error")
	}
}
