package Models

import (
	"forward-core/Utils"
)

type PageParam struct {
	PIndex int64
	PSize  int64
	// 要排序的字段名
	Sort string
	// ASC 或 DESC
	Direction string
}

func (this *PageParam) PageSize() int {

	return int(this.PSize)

}

// 分页, 排序处理
func (this *PageParam) SkipRows() int {
	this.PIndex = Utils.If(this.PIndex < 1, 1, this.PIndex).(int64)
	this.PSize = Utils.If(this.PSize < 1, 5, this.PSize).(int64)

	skipRows := (this.PIndex - 1) * this.PSize
	return int(skipRows)
}

func (this *PageParam) SortField() string {
	var sortField string
	if !Utils.IsEmpty(this.Sort) {
		if !Utils.IsEmpty(this.Direction) && this.Direction == "DESC" {
			//降序
			sortField = "-" + this.Sort
		} else {
			//升序
			sortField = this.Sort

		}
	}

	return sortField
}
