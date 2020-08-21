package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/lifei6671/mindoc/conf"
	"strings"
	"github.com/astaxie/beego"
)

// TODO: 暂未用到
// -------------------------------
// *
// * 区分针对Book和Doc的标签
// *
// -------------------------------

type LabelBook struct {
	LabelId    int    `orm:"column(label_id);pk;auto;unique;" json:"label_id"`
	LabelName  string `orm:"column(label_name);size(50);unique" json:"label_name"`
	BookNumber int    `orm:"column(book_number)" json:"book_number"`
}

// TableName 获取对应数据库表名.
func (m *LabelBook) TableName() string {
	return "label_book"
}

// TableEngine 获取数据使用的引擎.
func (m *LabelBook) TableEngine() string {
	return "INNODB"
}

func (m *LabelBook) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewLabelBook() *LabelBook {
	return &LabelBook{}
}

func (m *LabelBook) FindFirst(field string, value interface{}) (*LabelBook, error) {
	o := orm.NewOrm()

	err := o.QueryTable(m.TableNameWithPrefix()).Filter(field, value).One(m)

	return m, err
}

//插入或更新标签.
func (m *LabelBook) InsertOrUpdate(labelName string) error {
	o := orm.NewOrm()

	err := o.QueryTable(m.TableNameWithPrefix()).Filter("label_name", labelName).One(m)
	if err != nil && err != orm.ErrNoRows {
		return err
	}
	count, _ := o.QueryTable(NewDocument().TableNameWithPrefix()).Filter("labels__icontains", labelName).Count()
	m.BookNumber = int(count)
	m.LabelName = labelName

	if err == orm.ErrNoRows {
		err = nil
		m.LabelName = labelName
		_, err = o.Insert(m)
	} else {
		_, err = o.Update(m)
	}
	return err
}

//批量插入或更新标签.
func (m *LabelBook) InsertOrUpdateMulti(labels string) {
	if labels != "" {
		labelArray := strings.Split(labels, ",")

		for _, label := range labelArray {
			if label != "" {
				NewLabel().InsertOrUpdate(label)
			}
		}
	}
}
//删除标签
func (m *LabelBook) Delete() error {
	o := orm.NewOrm()
	_,err := o.Raw("DELETE FROM " + m.TableNameWithPrefix() + " WHERE label_id= ?",m.LabelId).Exec()

	if err != nil {
		return err
	}
	return nil
}

//分页查找标签.
func (m *LabelBook) FindToPager(pageIndex, pageSize int) (labels []*LabelBook, totalCount int, err error) {
	o := orm.NewOrm()

	count, err := o.QueryTable(m.TableNameWithPrefix()).Count()

	if err != nil {
		return
	}
	totalCount = int(count)

	offset := (pageIndex - 1) * pageSize

	_, err = o.QueryTable(m.TableNameWithPrefix()).OrderBy("-doc_number").Offset(offset).Limit(pageSize).All(&labels)

	if err == orm.ErrNoRows {
		beego.Info("没有查询到标签 ->",err)
		err = nil
		return
	}
	return
}




