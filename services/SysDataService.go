package services

import (
	"errors"
	"math"
	"port-forward/models"
	"port-forward/utils"
	"time"

	"github.com/astaxie/beego/logs"
)

type SysDataService struct {
}

func (_self *SysDataService) GetSysUserById(userId int) *models.SysUser {

	entity := new(models.SysUser)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("Id", userId)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetSysUserById ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) GetSysUserByName(userName string) *models.SysUser {

	entity := new(models.SysUser)
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
	pwd := utils.GetMd5(password)
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

func (_self *SysDataService) UpdateSysUser(entity *models.SysUser) error {

	_, err := OrmerS.Update(entity)
	return err
}

func (_self *SysDataService) DelSysUsers(ids []int) error {

	//批量删除
	del_num, err := OrmerS.QueryTable(new(models.SysUser)).Filter("Id__in", ids).Delete()
	if err != nil {
		logs.Error("DelSysUsers err：", err)
		return err
	} else {
		logs.Debug("DelSysUsers rows：", del_num)
	}

	return nil

}

func (_self *SysDataService) GetPortForwardById(id int) *models.PortForward {

	entity := new(models.PortForward)
	qs := OrmerS.QueryTable(entity)

	qs = qs.Filter("Id", id)

	err := qs.One(entity)

	if err != nil {
		logs.Error("GetPortForwardById ", err)
		return nil
	}

	return entity

}

func (_self *SysDataService) ChkPortForwardByApi(sourcePort, protocol, targetPort string) *models.PortForward {
	sourceArray := utils.Split(sourcePort, ":")
	targetArray := utils.Split(targetPort, ":")

	addr := sourceArray[0]
	port := utils.ToInt(sourceArray[1])

	toAddr := targetArray[0]
	toPort := utils.ToInt(targetArray[1])

	return _self.GetPortForwardByApi(addr, port, protocol, toAddr, toPort)
}

func (_self *SysDataService) GetPortForwardByApi(addr string, port int, protocol string, targetAddr string, targetPort int) *models.PortForward {

	entity := new(models.PortForward)
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

func (_self *SysDataService) SavePortForwardByApi(sourcePort, protocol, targetPort string) (*models.PortForward, error) {
	if utils.IsEmpty(sourcePort) {
		return nil, errors.New("本地监听端口信息不能为空")
	}

	if utils.IsEmpty(targetPort) {
		return nil, errors.New("目标端口信息不能为空")
	}

	sourceArray := utils.Split(sourcePort, ":")
	targetArray := utils.Split(targetPort, ":")

	entity := &models.PortForward{}
	entity.Name = "API_Forward_" + sourcePort
	entity.Addr = sourceArray[0]
	entity.Port = utils.ToInt(sourceArray[1])
	entity.Protocol = protocol
	entity.TargetAddr = targetArray[0]
	entity.TargetPort = utils.ToInt(targetArray[1])
	entity.CreateTime = time.Now()
	_, err := OrmerS.Insert(entity)
	return entity, err

}

func (_self *SysDataService) GetPortForwardList(query *models.PortForward, pageIndex int64, pageSize int64) models.PageData {

	var entites []*models.PortForward
	entity := new(models.PortForward)
	qs := OrmerS.QueryTable(entity)

	if query.Port > 0 {
		qs = qs.Filter("Port", query.Port)
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

	return models.PageData{PIndex: pageIndex, PSize: pageSize, TotalRows: totals, Pages: int64(pages), Data: entites}

}

func (_self *SysDataService) SavePortForward(entity *models.PortForward) error {

	if entity.Id > 0 {
		update := &models.PortForward{}

		qs := OrmerS.QueryTable(new(models.PortForward))
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

		_, err1 := OrmerS.Update(update)
		return err1
	} else {
		entity.CreateTime = time.Now()
		_, err := OrmerS.Insert(entity)
		return err
	}

}

func (_self *SysDataService) DelPortForwards(ids []int) error {
	//sqlite3在debug模式中，每次操作后会卡住，release环境中没有问题
	//批量删除
	del_num, err := OrmerS.QueryTable(new(models.PortForward)).Filter("Id__in", ids).Delete()
	if err != nil {
		logs.Error("DelPortForwards err：", err)
		return err
	} else {

		logs.Debug("DelPortForwards rows：", del_num)
	}

	return nil

}
