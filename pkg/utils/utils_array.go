package utils

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sArray struct {
	Data interface{}
}

// ArrayColumns 根据字段取出新切片数据
func (a sArray) ArrayColumns(field string) ([]interface{}, error) {
	// 统一转换为二维数组MAP
	slicesList := gconv.Maps(a.Data)
	if len(slicesList) == 0 {
		return nil, gerror.New(`数据为空`)
	}
	// 处理数据
	slices := make([]interface{}, 0)
	for _, slice := range slicesList {
		if _, ok := slice[field]; !ok {
			return nil, gerror.New(`数据字段` + field + `不存在`)
		}
		slices = append(slices, slice[field])
	}
	return slices, nil
}

// ColumnMap 二维数据字段合集Map
func (a sArray) ColumnMap(field string) (map[interface{}]interface{}, error) {
	// 统一转换为二维数组MAP
	slicesList := gconv.Maps(a.Data)
	if len(slicesList) == 0 {
		return nil, gerror.New(`数据为空`)
	}
	// 转换数据
	sMap := make(map[interface{}]interface{})
	for _, slice := range slicesList{
		if _, ok := slice[field]; !ok {
			return nil, gerror.New(`数据字段` + field + `不存在`)
		}
		sMap[slice[field]] = slice
	}
	return sMap, nil
}

// ArrayColumnsUnique 二维数组去重并获取字段集合
func (a sArray) ArrayColumnsUnique(field string) []interface{} {
	slices, err := a.ArrayColumns(field)
	if err != nil {
		return nil
	}
	// 数据为空
	if len(slices) == 0 {
		return nil
	}
	var (
		slicesUni = make([]interface{}, 0)
		dedup = make(map[interface{}]bool)
	)
	// 处理去重
	for _, slice := range slices{
		if _, ok := dedup[slice]; ok {
			continue
		}
		slicesUni = append(slicesUni, slice)
	}
	return slicesUni
}