package controllers


type HomeController struct {
	BaseController
}


func (this *HomeController) Get(){
	this.Data["name"] = "Jonathan"
	this.Layout = "public/layout.html"
	this.TplNames = "public/home.html"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["Navbar"] = "public/_navbar.html"
    this.LayoutSections["Menubar"] = "public/_menubar.html"
}