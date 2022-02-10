package parse

import (
	"encoding/xml"
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/beevik/etree"
	"io/ioutil"
	"os"
)

func ParseXML(path string)  {
	/**
	解析xml并把xml中的数据都存放到 json 里
	*/
	file,err := os.Open(path)
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()

	XMLdata, _ := ioutil.ReadAll(file)
	doc := data.Doc{}
	xml.Unmarshal(XMLdata, &doc)
	//fmt.Println(doc)
	var jsTest data.JsonTest
	//var jsTests []data.JsonTest
	// 耦合严重
	for _,test := range doc.Tests{
		jsTest.Title = test.Title
		jsTest.Stype = test.Stype
		jsTest.Level = test.Level
		jsTest.Risk  = test.Risk
		jsTest.Clause = test.Clause
		jsTest.Where = test.Where
		jsTest.Vector = test.Vector

		// 处理嵌套结构
		var jsRequest data.JsonRequest
		jsRequest.Payload = test.Request.Payload
		jsRequest.Comment = test.Request.Comment
		jsTest.Request = jsRequest

		var jsResponse data.JsonResponse
		jsResponse.Grep = test.Response.Grep
		jsTest.Response = jsResponse

		var jsDetails data.JsonDetails
		jsDetails.Dbms = test.Details.Dbms
		jsDetails.Dbms_version = test.Details.Dbms_version
		jsTest.Details = jsDetails

		//jsTests = append(jsTests,jsTest)
		data.Configure.Tests = append(data.Configure.Tests, jsTest)
	}
	//for _,value := range jsTests{
	//	data.Configure.Tests = append(data.Configure.Tests, value)
	//}
}

func ParseBoundaryXML(path string){
	//file := utils.ReadFile(path)
	file,err := os.Open(path)
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()

	XMLdata, _ := ioutil.ReadAll(file)
	doc := data.BoundaryDoc{}
	xml.Unmarshal(XMLdata, &doc)

	var boundaryTest data.BounaryTest
	for _,test := range doc.Boundarys{
		boundaryTest.Level = test.Level
		boundaryTest.Clause = test.Clause
		boundaryTest.Where = test.Where
		boundaryTest.Ptype = test.Ptype
		boundaryTest.Prefix = test.Prefix
		boundaryTest.Suffix = test.Suffix
		data.Configure.Boundaries = append(data.Configure.Boundaries,boundaryTest)
	}
}

func ParseQueryXML(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}
	root := doc.SelectElement("root")

	var dbs string
	queryData := make(map[string]map[string]string)
	for _, dbms := range root.SelectElements("dbms") {
		dbs = dbms.Attr[0].Value
		for _,d := range dbms.FindElements("*"){
			if len(d.Child) != 0 {
				for _,child := range d.FindElements("*"){
					// 处理 <users><inbind query=xxxx> 多层级情况
					queryData[d.Tag +"$"+ child.Tag] = _iterate(child)
				}
			}else {
				queryData[d.Tag] = _iterate(d)
			}
		}
	}
	data.Configure.Queries[dbs] = queryData
}


func _iterate(e *etree.Element) map[string]string {
	var temp = make(map[string]string)
	for _,w := range e.Attr{
		temp[w.Key] = w.Value
	}
	return temp
}
