package utils

import (
	"reflect"
	"strconv"
	"strings"
)

//type ParseXML interface {
//	ReadXML()
//}

func TransInt(s string) int {
	i,_ := strconv.Atoi(s)
	return i
}

func DictParams(place,params string) map[string]map[string]string {
	//var dictParam []map[string]string
	var dictParam = make(map[string]map[string]string)

	if strings.Contains(params,"&") {
		temp := make(map[string]string)
		for _,param := range strings.Split(params,"&"){
			key := strings.Split(param,"=")[0]
			value := strings.Split(param,"=")[1]
			temp[key] = value
		}
		dictParam[place] = temp

	}else {
		temp := make(map[string]string)
		key := strings.Split(params,"=")[0]
		value := strings.Split(params,"=")[1]
		temp[key] = value
		dictParam[place] = temp
	}
	return dictParam
}

func GetItems(o interface{}) (string, string) {
	object := reflect.TypeOf(o)
	fieldType := object.Field(object.NumMethod())
	fieldValue := reflect.ValueOf(o).FieldByName(fieldType.Name)
	return fieldType.Name,fieldValue.Interface().(string)
}

func ContainsKey(m map[string]string,key string) bool {
	for k,_ := range m{
		if k == key {
			return true
		}
	}
	return false
}