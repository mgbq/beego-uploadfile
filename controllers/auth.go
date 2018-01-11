package controllers

import ()

type AuthController struct {
	BaseController
}

func (this *AuthController) Prepare() {
	v := this.GetSession("userid")
	if v == nil {
		this.Redirect("/index/index", 302)
	}
}
