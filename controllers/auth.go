package controllers

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	v := c.GetSession("userid")
	if v == nil {
		c.Redirect("/index/index", 302)
		return
	}

}
