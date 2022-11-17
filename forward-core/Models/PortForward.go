package Models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type PortForward struct {
	Id   int    `orm:"column(id);pk;auto"`
	Name string `orm:"column(name);size(256);null"`
	// 0:禁用,1:启用
	Status int    `orm:"column(status);null"`
	Addr   string `orm:"column(addr);size(256);null"`
	// 端口号
	Port int `orm:"column(port);null"`
	//协议
	Protocol   string `orm:"column(protocol);size(32);null"`
	TargetAddr string `orm:"column(targetAddr);size(256);null"`
	// 端口号
	TargetPort int       `orm:"column(targetPort);null"`
	CreateTime time.Time `orm:"column(createTime);type(datetime)"`
	//暂时用来存放端口分发配置，后续版本再调整
	Others string `orm:"column(others);size(500);null"`
	// 0:普通映射,1:内网穿透映射（Server->Client）,2:内网穿透反向映射（Client->Server）
	FType int `orm:"column(fType);null"`
}

func (t *PortForward) TableName() string {
	return "t_port_forward"
}

func init() {
	orm.RegisterModel(new(PortForward))
}
