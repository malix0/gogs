// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routers

import (
	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/base"
	"github.com/gogits/gogs/modules/middleware"
	"github.com/gogits/gogs/routers/user"
)

func Home(ctx *middleware.Context) {
	if ctx.IsSigned {
		user.Dashboard(ctx)
		return
	}

	// Check auto-login.
	userName := ctx.GetCookie(base.CookieUserName)
	if len(userName) != 0 {
		ctx.Redirect("/user/login")
		return
	}

	// Show recent updated repositoires for new visiters.
	repos, err := models.GetRecentUpdatedRepositories()
	if err != nil {
		ctx.Handle(500, "dashboard.Home(GetRecentUpdatedRepositories)", err)
		return
	}

	for _, repo := range repos {
		repo.Owner, err = models.GetUserById(repo.OwnerId)
		if err != nil {
			ctx.Handle(500, "dashboard.Home(GetUserById)", err)
			return
		}
	}
	ctx.Data["Repos"] = repos
	ctx.Data["PageIsHome"] = true
	ctx.HTML(200, "home")
}

func Help(ctx *middleware.Context) {
	ctx.Data["PageIsHelp"] = true
	ctx.Data["Title"] = "Help"
	ctx.HTML(200, "help")
}

func NotFound(ctx *middleware.Context) {
	ctx.Data["PageIsNotFound"] = true
	ctx.Data["Title"] = "Page Not Found"
	ctx.Handle(404, "home.NotFound", nil)
}
