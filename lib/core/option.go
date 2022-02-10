package core

import (
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/parse"
	"path/filepath"
	"strings"
)

func Init(){
	createHomeDirectories()
	if data.Configure.Url != ""{
		_setHostname()
		_setRequestMethod()
		_setTimeout()
		_setUserAgent()
	}
	_setThreads()
	SetPath()
	loadBoundaries()  // payload 加载
	loadPayloads()
	loadQueries()
}

func createHomeDirectories(){
	/**
		创建 output 文件夹
	 */
}

func _setHostname(){
	/**
	设置hostname需要进行切割端口部分
	https://studygolang.com/articles/2876
	 */
	//data.Configure.Hostname,_ = url.Parse(data.Configure.Url)
}

func _setRequestMethod(){

	data.Configure.Method = "GET"
}

func _setTimeout(){
	/**
	设置超时时间，timeout 不能小于 3
	 */
	if data.Configure.Timeout < 3 {
		data.Configure.Timeout = 3
	}
}

func _setUserAgent()  {
	
}

func _setThreads()  {
	data.Configure.Thread = 1
}

func loadBoundaries()  {

	for _,boundariesFile := range strings.Split(data.BOUNDARIES_XML_FILES,","){
		boundariesFilePath := filepath.Join(data.SQLMAP_XML_PATH,boundariesFile)
		parse.ParseBoundaryXML(boundariesFilePath)
	}
}

func loadPayloads()  {
	/**
	https://studygolang.com/articles/5328 xml 解析
	 */
	for _,payloadFile := range strings.Split(data.PAYLOAD_XML_FILES,","){
		payloadFilePath := filepath.Join(data.SQLMAP_PAYLOADS_PATH,payloadFile)
		parse.ParseXML(payloadFilePath)
	}
}

func loadQueries(){
	for _,queriesFile := range strings.Split(data.QUERIES_XML_FILES,","){
		queriesFilePath := filepath.Join(data.SQLMAP_XML_PATH,queriesFile)
		parse.ParseQueryXML(queriesFilePath)
	}
}




