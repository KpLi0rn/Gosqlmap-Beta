package utils

import (
	"fmt"
	"os"
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

func ReadFile(path string) *os.File {
	file,err := os.Open(path)
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()
	return file
}

type Demo interface {

}
// 如果能直接获取到的话那么就直接获取到并进行返回
// 这个函数后期必须要进行修改
func GetInner(m map[string]map[string]string,name string) (map[string]string,map[string]map[string]string) {
	temp := make(map[string]map[string]string)
	tag := name + "$"
	for k,v := range m{
		if strings.Contains(k,tag) {
			newTag := strings.Replace(k,tag,"",-1)
			temp[newTag] = v
		}
	}
	if len(temp) == 0 {
		return m[name],nil
	}
	return nil,temp
}
