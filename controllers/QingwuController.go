package controllers

import (
	"net/url"
	"github.com/lifei6671/mindoc/conf"
	"os/exec"
)

type QingwuController struct {
	BaseController
}

func (c *QingwuController) Index() {
	c.Prepare()
	c.TplName = "qingwu/index.tpl"
	c.Data["JSON_NAME"] = "book_list"

	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login")+"?url="+url.PathEscape(conf.BaseUrl+c.Ctx.Request.URL.RequestURI()), 302)
	}
}

func (c *QingwuController) ViewKnowledgeGraph() {
	c.Prepare()
	c.TplName = "qingwu/index.tpl"
	var book_id = c.GetString(":book_id")
	c.Data["JSON_NAME"] = "book_" + book_id

	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login")+"?url="+url.PathEscape(conf.BaseUrl+c.Ctx.Request.URL.RequestURI()), 302)
	}
}

func (c *QingwuController) ViewSubKnowledgeGraph() {
	c.Prepare()
	c.TplName = "qingwu/index.tpl"
	var doc_id = c.GetString(":doc_id")
	c.Data["JSON_NAME"] = "doc_" + doc_id

	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login")+"?url="+url.PathEscape(conf.BaseUrl+c.Ctx.Request.URL.RequestURI()), 302)
	}
}

func (c *QingwuController) UpdateKnowledgeGraph() {
	c.Prepare()
	c.TplName = "qingwu/update_kg.tpl"
	exec.Command("C:\\Programs\\Python\\python.exe",
		"D:/Documents/Programs/QwMinderServer/scripts.py").Run()
	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login")+"?url="+url.PathEscape(conf.BaseUrl+c.Ctx.Request.URL.RequestURI()), 302)
	}
}