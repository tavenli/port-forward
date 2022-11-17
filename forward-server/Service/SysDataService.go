package Service

import (
	"errors"
	"forward-core/Models"
	"forward-core/Utils"
	"math"
	"time"

	"github.com/astaxie/beego/logs"
)

type SysDataService struct {
}

func (_self *SysDataService) GetSysUserById(userId int) *Models.SysUser {

	entity := new(Models.SysUser)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("Id", userId)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetSysUserById ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) GetSysUserByName(userName string) *Models.SysUser {

	entity := new(Models.SysUser)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("UserName", userName)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetSysUserByName ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) ChangeUserPwd(id int, password string) error {
	pwd := Utils.GetMd5(password)
	res, err := OrmerS.Raw("update t_sys_user SET passWord = ? where id = ?",
		pwd, id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		logs.Debug("ChangeUserPwd", num)

	} else {
		logs.Error("ChangeUserPwd", err)

	}
	return err
}

func (_self *SysDataService) UpdateSysUser(entity *Models.SysUser) error {

	_, err := OrmerS.Update(entity)
	return err
}

func (_self *SysDataService) DelSysUsers(ids []int) error {

	//批量删除
	del_num, err := OrmerS.QueryTable(new(Models.SysUser)).Filter("Id__in", ids).Delete()
	if err != nil {
		logs.Error("DelSysUsers err：", err)
		return err
	} else {
		logs.Debug("DelSysUsers rows：", del_num)
	}

	return nil

}

func (_self *SysDataService) GetPortForwardById(id int) *Models.PortForward {

	entity := new(Models.PortForward)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("Id", id)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetPortForwardById ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) ChkPortForwardByApi(sourcePort, protocol, targetPort string) *Models.PortForward {
	sourceArray := Utils.Split(sourcePort, ":")
	targetArray := Utils.Split(targetPort, ":")

	addr := sourceArray[0]
	port := Utils.ToInt(sourceArray[1])

	toAddr := targetArray[0]
	toPort := Utils.ToInt(targetArray[1])

	return _self.GetPortForwardByApi(addr, port, protocol, toAddr, toPort)
}

func (_self *SysDataService) GetPortForwardByApi(addr string, port int, protocol string, targetAddr string, targetPort int) *Models.PortForward {

	entity := new(Models.PortForward)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("Addr", addr)
	qs = qs.Filter("Port", port)
	qs = qs.Filter("Protocol", protocol)
	qs = qs.Filter("TargetAddr", targetAddr)
	qs = qs.Filter("TargetPort", targetPort)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetPortForwardByApi ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) SavePortForwardByApi(sourcePort, protocol, targetPort string) (*Models.PortForward, error) {
	if Utils.IsEmpty(sourcePort) {
		return nil, errors.New("本地监听端口信息不能为空")
	}

	if Utils.IsEmpty(targetPort) {
		return nil, errors.New("目标端口信息不能为空")
	}

	sourceArray := Utils.Split(sourcePort, ":")
	targetArray := Utils.Split(targetPort, ":")

	entity := &Models.PortForward{}
	entity.Name = "API_Forward_" + sourcePort
	entity.Addr = sourceArray[0]
	entity.Port = Utils.ToInt(sourceArray[1])
	entity.Protocol = protocol
	entity.TargetAddr = targetArray[0]
	entity.TargetPort = Utils.ToInt(targetArray[1])
	entity.CreateTime = time.Now()
	_, err := OrmerS.Insert(entity)
	return entity, err

}

func (_self *SysDataService) GetAllPortForwardList(status int) []*Models.PortForward {

	var entites []*Models.PortForward
	entity := new(Models.PortForward)
	qs := OrmerS.QueryTable(entity)

	if status > -1 {
		qs = qs.Filter("Status", status)
	}

	num, err := qs.All(&entites)
	if err != nil {
		logs.Error("GetAllPortForwardList ", err)
	} else {
		logs.Debug("GetAllPortForwardList rows ", num)
	}

	return entites

}

func (_self *SysDataService) GetPortForwardList(query *Models.PortForward, pageIndex int64, pageSize int64) Models.PageData {

	var entites []*Models.PortForward
	entity := new(Models.PortForward)
	qs := OrmerS.QueryTable(entity)

	if query.Port > 0 {
		qs = qs.Filter("Port", query.Port)
	}

	if len(query.TargetAddr) > 0 {
		qs = qs.Filter("TargetAddr__icontains", query.TargetAddr)
	}

	if query.TargetPort > 0 {
		qs = qs.Filter("TargetPort", query.TargetPort)
	}

	if query.FType > -1 {
		qs = qs.Filter("FType", query.FType)
	}

	totals, _ := qs.Count()
	pages := math.Ceil(float64(totals) / float64(pageSize))

	if pageIndex <= 0 {
		pageIndex = 1
	}
	offset := pageIndex - 1
	num, err := qs.Limit(pageSize, pageSize*offset).OrderBy("-Id").All(&entites)
	if err != nil {
		logs.Error("GetPortForwardList ", err)
	} else {
		logs.Debug("GetPortForwardList rows ", num)
	}

	return Models.PageData{PIndex: pageIndex, PSize: pageSize, TotalRows: totals, Pages: int64(pages), Data: entites}

}

func (_self *SysDataService) SavePortForward(entity *Models.PortForward) error {

	if entity.Id > 0 {
		update := &Models.PortForward{}

		qs := OrmerS.QueryTable(new(Models.PortForward))
		qs = qs.Filter("Id", entity.Id)
		err := qs.One(update)
		if err != nil {
			//如果没查到数据，会抛出 no row found
			return err
		}

		update.Name = entity.Name
		update.Addr = entity.Addr
		update.Port = entity.Port
		update.Protocol = entity.Protocol
		update.TargetAddr = entity.TargetAddr
		update.TargetPort = entity.TargetPort
		update.Others = entity.Others
		update.FType = entity.FType
		update.Status = entity.Status

		_, err1 := OrmerS.Update(update)
		return err1
	} else {
		entity.CreateTime = time.Now()
		res, err := OrmerS.Raw("INSERT INTO t_port_forward(name, status, addr, port, protocol, targetAddr, targetPort, createTime, others, fType) values(?,?,?,?,?,?,?,?,?,?)",
			entity.Name, entity.Status, entity.Addr, entity.Port, entity.Protocol, entity.TargetAddr, entity.TargetPort, entity.CreateTime, entity.Others, entity.FType).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			logs.Debug("AddPortForward", num)

		} else {
			logs.Error("AddPortForward", err)

		}

		return err
	}

}

func (_self *SysDataService) DelPortForwards(ids []int) error {
	//sqlite3在debug模式中，每次操作后会卡住，release环境中没有问题
	//批量删除
	del_num, err := OrmerS.QueryTable(new(Models.PortForward)).Filter("Id__in", ids).Delete()
	if err != nil {
		logs.Error("DelPortForwards err：", err)
		return err
	} else {

		logs.Debug("DelPortForwards rows：", del_num)
	}

	return nil

}

func (_self *SysDataService) ToForwardConfig(entity *Models.PortForward) *Models.ForwardConfig {
	config := new(Models.ForwardConfig)
	config.RuleId = entity.Id
	config.Name = entity.Name
	config.Protocol = entity.Protocol
	config.SrcAddr = entity.Addr
	config.SrcPort = entity.Port
	config.DestAddr = entity.TargetAddr
	config.DestPort = entity.TargetPort
	config.Status = entity.Status
	config.Others = entity.Others

	return config

}

func (_self *SysDataService) GetForwardJob(entity *Models.PortForward) *ForWardJob {

	return ForWardServ.GetForwardJob(_self.ToForwardConfig(entity))

}
