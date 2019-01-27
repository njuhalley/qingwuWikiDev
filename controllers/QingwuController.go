package controllers

import (
	"github.com/lifei6671/mindoc/conf"
	"github.com/lifei6671/mindoc/models"
	"github.com/lifei6671/mindoc/utils"
	"github.com/lifei6671/mindoc/utils/pagination"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
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

//RECENT
func (c *QingwuController) List() {
	c.Prepare()
	c.TplName = "qingwu/recent.tpl"

	//如果没有开启你们访问则跳转到登录
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login"), 302)
		return
	}

	// keyword := c.GetString("keyword")
	var PageSize = 5
	keyword := `""`
	pageIndex, _ := c.GetInt("page", 1)

	c.Data["BaseUrl"] = c.BaseUrl()

	memberId := 0
	if c.Member != nil {
		memberId = c.Member.MemberId
	}
	searchResult, totalCount, err := models.RecentDocumentResult().FindToPager(keyword, pageIndex, PageSize, memberId)
	
	totalCount = 5

	if err != nil {
		return
	}
	if totalCount > 0 {
		pager := pagination.NewPagination(c.Ctx.Request, totalCount, PageSize,c.BaseUrl())
		c.Data["PageHtml"] = pager.HtmlPages()
	} else {
		c.Data["PageHtml"] = ""
	}
	if len(searchResult) > 0 {
		keywords := strings.Split(keyword," ")

		for _, item := range searchResult {
			for _,word := range keywords {
				item.DocumentName = strings.Replace(item.DocumentName, word, "<em>"+word+"</em>", -1)
				if item.Description != "" {
					src := item.Description

					r := []rune(utils.StripTags(item.Description))

					if len(r) > 100 {
						src = string(r[:100])
					} else {
						src = string(r)
					}
					item.Description = strings.Replace(src, word, "<em>"+word+"</em>", -1)
				}
			}
			if item.Identify == "" {
				item.Identify = strconv.Itoa(item.DocumentId)
			}
			if item.ModifyTime.IsZero() {
				item.ModifyTime = item.CreateTime
			}
		}
	}
	c.Data["Lists"] = searchResult
}