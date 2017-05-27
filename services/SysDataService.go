package services

import (
	"math"
	"port-forward/models"

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
