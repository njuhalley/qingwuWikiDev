package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lifei6671/mindoc/conf"
	"github.com/lifei6671/mindoc/models"
	"github.com/lifei6671/mindoc/utils/pagination"
	"strconv"
)

type DocumentStarController struct {
	BaseController
}

func (c *DocumentStarController) Prepare() {
	c.BaseController.Prepare()

	//如果没有开启你们访问则跳转到登录
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login"), 302)
		return
	}
}

//查看星标的文档列表.
func (c *DocumentStarController) Index() {
	c.Prepare()
	c.TplName = "doc_labels/document_star.tpl"

	pageIndex, _ := c.GetInt("page", 1)
	bookId, err := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
	book, _ := models.NewBook().Find(bookId)

	memberId := 0
	if c.Member != nil {
		memberId = c.Member.MemberId
	}
	// TODO: 查找包含某个标签的所有文档
	searchResult, totalCount, err := models.NewDocument().FindForStarToPager(bookId, pageIndex, conf.PageSize, memberId)

	if err != nil && err != orm.ErrNoRows {
		beego.Error("查询星标文档时出错 ->", err)
		c.ShowErrorPage(500, "查询星标文档列表时出错")
	}
	if totalCount > 0 {
		pager := pagination.NewPagination(c.Ctx.Request, totalCount, conf.PageSize, c.BaseUrl())
		c.Data["PageHtml"] = pager.HtmlPages()
	} else {
		c.Data["PageHtml"] = ""
	}
	searchParentsResult, err := models.NewDocument().FindForItemParents(searchResult)  // 查找parents
	c.Data["Lists"] = searchParentsResult
	c.Data["Model"] = book
	c.Data["Remarks"] = "is_star"
}

//查看 为公文 的文档列表.
func (c *DocumentStarController) IsDoc() {
	c.Prepare()
	c.TplName = "doc_labels/document_star.tpl"

	pageIndex, _ := c.GetInt("page", 1)
	bookId, err := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
	book, _ := models.NewBook().Find(bookId)

	memberId := 0
	if c.Member != nil {
		memberId = c.Member.MemberId
	}
	// TODO: 查找包含某个标签的所有文档
	searchResult, totalCount, err := models.NewDocument().FindForIsDocToPager(bookId, pageIndex, conf.PageSize, memberId)

	if err != nil && err != orm.ErrNoRows {
		beego.Error("查询星标文档时出错 ->", err)
		c.ShowErrorPage(500, "查询星标文档列表时出错")
	}
	if totalCount > 0 {
		pager := pagination.NewPagination(c.Ctx.Request, totalCount, conf.PageSize, c.BaseUrl())
		c.Data["PageHtml"] = pager.HtmlPages()
	} else {
		c.Data["PageHtml"] = ""
	}
	searchParentsResult, err := models.NewDocument().FindForItemParents(searchResult)  // 查找parents
	c.Data["Lists"] = searchParentsResult
	c.Data["Model"] = book
	c.Data["Remarks"] = "is_doc"
}

//查看 为简历 的文档列表.
func (c *DocumentStarController) IsResume() {
	c.Prepare()
	c.TplName = "doc_labels/document_star.tpl"

	pageIndex, _ := c.GetInt("page", 1)
	bookId, err := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
	book, _ := models.NewBook().Find(bookId)

	memberId := 0
	if c.Member != nil {
		memberId = c.Member.MemberId
	}
	// TODO: 查找包含某个标签的所有文档
	searchResult, totalCount, err := models.NewDocument().FindForIsResumeToPager(bookId, pageIndex, conf.PageSize, memberId)

	if err != nil && err != orm.ErrNoRows {
		beego.Error("查询星标文档时出错 ->", err)
		c.ShowErrorPage(500, "查询星标文档列表时出错")
	}
	if totalCount > 0 {
		pager := pagination.NewPagination(c.Ctx.Request, totalCount, conf.PageSize, c.BaseUrl())
		c.Data["PageHtml"] = pager.HtmlPages()
	} else {
		c.Data["PageHtml"] = ""
	}
	searchParentsResult, err := models.NewDocument().FindForItemParents(searchResult)  // 查找parents
	c.Data["Lists"] = searchParentsResult
	c.Data["Model"] = book
	c.Data["Remarks"] = "is_resume"
}