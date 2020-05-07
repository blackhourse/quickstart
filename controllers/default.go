package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "register.html"
	//c.Ctx.WriteString("hello,world")
	beego.AppConfig.String("mysqluser")

}

func (c *MainController) Post() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		c.Redirect("/register", 302)
		return
	}
	c.Ctx.WriteString("注册成功")
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}
func (c *MainController) HandleLogin() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/login.html", 302)
		return
	}

	user := models.User{Name: userName}
	o := orm.NewOrm()
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("没有此用户")
		c.TplName = "/login.html"
		return
	}
	c.Ctx.WriteString("登录成功")

}

// 首页
func (c *MainController) ShowIndex() {
	c.TplName = "/index.html"
}
