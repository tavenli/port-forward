package Models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id       int    `orm:"column(id);pk;auto"`
	UserName string `orm:"column(userName);null"`
	PassWord string `orm:"column(passWord);null"`
	// 0:禁用,1:启用
	Status     int       `orm:"column(status);null"`
	CreateTime time.Time `orm:"column(createTime);type(datetime)"`
}

func (t *SysUser) TableName() string {
	return "t_sys_user"
}

func init() {
	orm.RegisterModel(new(SysUser))

}
