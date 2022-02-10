package core

import (
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
	"net/url"
	"regexp"
	"strings"
)

var SQLAgent Agent

type Agent struct {

}

func (a *Agent) ExtractPayload() string {
	return ""
}

func (a *Agent) AdjustLateValues()  {

}

// AND (SELECT 2*(IF((SELECT * FROM (SELECT CONCAT('[DELIMITER_START]',(SELECT (ELT([RANDNUM]=[RANDNUM],1))),'[DELIMITER_STOP]','x'))s), 8446744073709551610, 8446744073709551610)))
func (a *Agent) CleanupPayload(payload,value string) string {
	var originValue string
	if len(value) != 0 {
		originValue = value
	} else {
		originValue = ""
	}

	var replacements = make(map[string]string,7)
	replacements["[DELIMITER_START]"] = data.Kb.Chars.Start
	replacements["[DELIMITER_STOP]"] = data.Kb.Chars.Stop
	replacements["[AT_REPLACE]"] = data.Kb.Chars.At
	replacements["[SPACE_REPLACE]"] = data.Kb.Chars.Space
	replacements["[DOLLAR_REPLACE]"] = data.Kb.Chars.Dollar
	replacements["[HASH_REPLACE]"] = data.Kb.Chars.Hash_
	replacements["[GENERIC_SQL_COMMENT]"] = data.GENERIC_SQL_COMMENT


	r, _ := regexp.Compile("\\[[A-Z_]+\\]")
	for _,value := range r.FindAllString(payload,-1){
		if utils.ContainsKey(replacements,value)  {
			payload = strings.Replace(payload,value,replacements[value],-1)
		}
	}

	r, _ = regexp.Compile("(?i)\\[RANDNUM(?:\\d+)?\\]")
	for _,value := range r.FindAllString(payload,-1){
		payload = strings.Replace(payload,value,RandomInt(0),-1)
	}

	r, _ = regexp.Compile("(?i)\\[RANDSTR(?:\\d+)?\\]")
	for _,value := range r.FindAllString(payload,-1){
		payload = strings.Replace(payload,value,RandomStr(0,false,nil),-1)
	}

	if strings.Contains(payload,"[ORIGINAL]") {
		payload = strings.Replace(payload,"[ORIGINAL]",originValue,-1)
	}
	if strings.Contains(payload,"[ORIGVALUE]") {
		payload = strings.Replace(payload,"[ORIGVALUE]",originValue,-1)
	}


	return payload
}

// 添加 payload 前缀
func (a *Agent) PrefixQuery(payload,prefix string,where int,clause string) string{
	prefix = a.CleanupPayload(prefix,"")
	return fmt.Sprintf("%s %s",prefix,payload)
}

// 添加 payload 后缀
func (a *Agent) SuffixQuery(payload,comment,suffix string, where int) string{
	suffix = a.CleanupPayload(suffix,"")
	return fmt.Sprintf("%s%s",payload,suffix)
}

// 生成请求的 payload 即合并参数
func (a *Agent) Payload(place,parameter,boundPayload string, where int) string{

	if len(data.Kb.Place) == 0 || len(data.Kb.Parameter) ==0 {
		data.Kb.Place = place
		data.Kb.Parameter = parameter
	}
	/**
	parameter -> id
	 */
	paramString := data.Configure.Params[place]
	paramDict := data.Configure.ParamsDict[place]
	originValue := paramDict[parameter]


	before := fmt.Sprintf("%s=%s",parameter,originValue)
	payload := fmt.Sprintf("%s=%s%s",parameter,originValue,url.QueryEscape(boundPayload))
	//payload := fmt.Sprintf("%s=%s%s",parameter,originValue,boundPayload)
	reqPayload := strings.Replace(paramString,before,payload,-1)
	return reqPayload
}


// 正则提取payload中的各个部分
func (a *Agent) GetFields(query string) string{
	fieldsNoSelect := query
	fieldsToCastStr := fieldsNoSelect

	// 获取 select from 之间的参数
	prefixRegex := "(?:\\s+(?:FIRST|SKIP|LIMIT(?: \\d+)?)\\s+\\d+)*"
	fieldsSelectFrom, _ := regexp.Compile(fmt.Sprintf("\\ASELECT%s\\s+(.+?)\\s+FROM ",prefixRegex))

	for _,value := range fieldsSelectFrom.FindAllStringSubmatch(query,-1){
		fieldsToCastStr = value[1]
	}
	return fieldsToCastStr
}


//
func (a *Agent) NullAndCastField(field string) string {
	var nulledCastedField string
	if len(field) != 0 && len(data.Configure.Dbms) != 0{
		rootQuery := data.Configure.Queries[data.Configure.Dbms]
		if strings.HasPrefix(field,"(CASE") || strings.HasPrefix(field,"(IIF"){
			nulledCastedField = field
		}
		query,_ := utils.GetInner(rootQuery,"cast")
		nulledCastedField = fmt.Sprintf(query["query"],field)
	}
	return nulledCastedField
}


func (a *Agent)LimitQuery(num int, query,field string) string{
	limitedQuery := query
	Str,_ := utils.GetInner(data.Configure.Queries[data.Configure.Dbms],"limit")
	limitStr := Str["query"]
	if data.Configure.Dbms == "MySQL" {
		limitStr = fmt.Sprintf(limitStr,num,1)
		limitedQuery += " " + limitStr
	}

	return limitedQuery
}