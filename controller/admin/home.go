package admin

import (
	"github.com/ZJGSU-ACM/GoOnlineJudge/class"

	"github.com/ZJGSU-ACM/restweb"
)

type AdminHome struct {
	class.Controller
} //@Controller

//@URL: /admin/ @method: GET
func (hc *AdminHome) Home() {
	restweb.Logger.Debug("Admin Home")

	hc.Output["IsHome"] = true
	hc.Output["Title"] = "Admin - Home"
	hc.RenderTemplate("view/admin/layout.tpl", "view/admin/home.tpl")
}
