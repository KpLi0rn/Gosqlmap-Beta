package request

import (
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/log"
	"io/ioutil"
	"net/http"
)



// 设置请求参数，然后进行发包
func getPage(uri ,get ,post ,method string) (string,map[string][]string,int) {

	if len(data.Configure.UserAgent) == 0 {
		data.Configure.UserAgent = data.DefaultArgs.UserAgent
	}

	client := &http.Client{}
	var resp *http.Response
	var err error
	var page string
	var code int
	var header map[string][]string
	if method == "GET" {
		req,_ := http.NewRequest(method,uri,nil)
		req.Header.Add(data.HttpHeader.UserAgent,data.Configure.UserAgent)
		resp,err = client.Do(req)
		if err != nil {
			log.Error("error occurred while sending Request")
		}
		defer resp.Body.Close()

		body,_ := ioutil.ReadAll(resp.Body)
		page = string(body)
		headers := make(map[string][]string)
		for key,value := range resp.Header{
			headers[key] = value
		}
		code = resp.StatusCode
	}

	// page 需要根据返回头中的编码进行解码,
	return page,header,code
}

func QueryPage(payload,place string) (string,map[string][]string,int) {

	var pageLength int
	uri := data.Configure.BaseUrl
	method := data.Configure.Method
	var page string
	var code int
	var header map[string][]string

	//payload := core.SQLAgent.ExtractPayload()
	if len(payload) != 0 {
		if place == data.Place.GET{
			if pageLength == 0{
				uri = fmt.Sprintf("%s?%s",uri,payload)
				page, header, code = getPage(uri,"","",place)
				return page,header,code
			}
		}
	}

	var get,post string
	// 获取 get 请求中的请求参数
	if data.Configure.Params[data.Place.GET] != "" {
		_ = data.Configure.Params[data.Place.GET]
	}

	if data.Configure.Params[data.Place.POST] != "" {
		_ = data.Configure.Params[data.Place.POST]
	}

	if pageLength == 0 {
		page, header, code = getPage(uri,get,post,method)
	}
	return page,header,code
}



