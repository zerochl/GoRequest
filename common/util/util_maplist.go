package util

import (
	"container/list"
	"sort"
	"bytes"
	"strings"
	"GoRequest/common/util/maplist"
)

type Keyer interface {
	GetKey() string
}

type MapList struct {
	dataMap  map[string]*list.Element
	dataList *list.List
}

type Elements struct {
	Key   string
	Value string
}

func (e Elements) GetKey() string {
	return e.Key
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (mapList *MapList) GetDataMap() map[string]*list.Element {
	return mapList.dataMap
}

func (mapList *MapList) Exists(data Keyer) bool {
	_, exists := mapList.dataMap[string(data.GetKey())]
	return exists
}

func (mapList *MapList) Push(data Keyer) bool {
	if mapList.Exists(data) {
		return false
	}
	elem := mapList.dataList.PushBack(data)
	mapList.dataMap[data.GetKey()] = elem
	return true
}

func (mapList *MapList) Remove(data Keyer) {
	if !mapList.Exists(data) {
		return
	}
	mapList.dataList.Remove(mapList.dataMap[data.GetKey()])
	delete(mapList.dataMap, data.GetKey())
}

func (mapList *MapList) Size() int {
	return mapList.dataList.Len()
}

func (mapList *MapList) Walk(cb func(data Keyer)) {
	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		cb(elem.Value.(Keyer))
	}
}

func (mapList *MapList) WalkByKey(cb func(data Keyer), keys []string) {
	//for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
	//	cb(elem.Value.(Keyer))
	//}
	for i := 0; i < len(keys); i++ {
		cb(mapList.GetDataMap()[keys[i]].Value.(*maplist.Elements))
	}
}

func (mapList *MapList) GetSignParam() string {
	// 先执行key sort进行排序
	var keys []string
	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(Keyer).GetKey())
	}
	sort.Strings(keys)

	var paramBuffer bytes.Buffer
	cb := func(data Keyer) {
		element := mapList.GetDataMap()[data.GetKey()].Value.(*maplist.Elements)
		paramBuffer.WriteString(element.Key)
		paramBuffer.WriteString("=")
		switch element.Value.(type) {
		case string:
			paramBuffer.WriteString(element.Value.(string))
		case float64:
			paramBuffer.WriteString(float64ToString(element.Value.(float64)))
		}
	}
	mapList.WalkByKey(cb, keys)
	return paramBuffer.String()
}

func (mapList *MapList) GetRequestParam() string {
	var paramBuffer bytes.Buffer
	paramBuffer.WriteString("?")
	cb := func(data Keyer) {
		element := mapList.GetDataMap()[data.GetKey()].Value.(*maplist.Elements)
		paramBuffer.WriteString(element.Key)
		paramBuffer.WriteString("=")
		switch element.Value.(type) {
		case string:
			paramBuffer.WriteString(element.Value.(string))
		case float64:
			paramBuffer.WriteString(float64ToString(element.Value.(float64)))
		}
		paramBuffer.WriteString("&")
	}
	mapList.Walk(cb)
	param := paramBuffer.String()
	param = strings.TrimRight(param, "&")
	return param
}
