package controller

import (
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/core"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/log"
	"github.com/KpLi0rn/Gosqlmap/lib/request"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
	"net"
	"net/url"
	"regexp"
	"strings"
)

func CheckConnection() bool {
	// 如果不是ip，那么就通过 socket 检测域名的连通性，不连通就报错
	// 测试 DNS查询是否正常
	IpMatch,_ := regexp.MatchString("\\A\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\Z",data.Configure.Hostname)
	// 没有匹配到那么就是域名
	if !IpMatch {
		_,err := net.LookupHost(data.Configure.Hostname)
		if err != nil {
			log.Error(fmt.Sprintf("resolving hostname '%s' Error",data.Configure.Hostname))
			return false
		}
	}
	// 测试网络连通性
	_,_,code := request.QueryPage("","")
	// 根据状态码来进行判断
	if code < 200 || code >= 300 {
		return false
	}
	return true
}

func CheckWaf(){

}

func CheckDynParam(place,parameter,value string) bool {
	return true
}

func HeuristicCheckSqlInjection(place,parameter,value string) bool {

	/**
	启发式扫描SQL注入
	 */

	//originValue := value  // value
	//paramType := place  // GET
	//
	prefix := ""
	suffix := ""
	randStr := ""

	// 同时满足的情况需要用 or
	for {
		randStr = core.RandomStr(10,false,data.HEURISTIC_CHECK_ALPHABET)
		if strings.Count(randStr,"'") == 1 && strings.Count(randStr,"\"") == 1 {
			break
		}
	}

	//payload := fmt.Sprintf("%s%s%s",prefix,suffix,randStr)
	fmt.Sprintf("%s%s%s",prefix,suffix,randStr)

	//    payload = agent.payload(place, parameter, newValue=payload)  id=1 => id=___payload___1)'".,)(..)___payload___


	return true
}

func CheckSqlInjection(place,parameter,value string) bool {

	var inject core.Agent

	tests := data.Configure.Tests
	boundaries := data.Configure.Boundaries

	for _,test := range tests{

		if len(data.Configure.Dbms) == 0 {
			// 还未缺认数据库
			// 还没做完，关注 sqlmap kb.htmlFp 参数
			data.Configure.Dbms = "MYSQL"
			data.Kb.ReduceTests = "MYSQL"  // 减少 payload 测试，只做mysql的sql注入
		}

		fastPayload := inject.CleanupPayload(test.Request.Payload,value)

		// 闭合 + payload
		for _,boundary := range boundaries{
			var prefix string
			var suffix string
			//ptype := boundary.Ptype

			if len(boundary.Prefix) != 0 {
				prefix = boundary.Prefix
			}else {
				prefix = ""
			}

			if len(boundary.Suffix) != 0 {
				suffix = boundary.Suffix
			}else {
				suffix = ""
			}

			if len(fastPayload) != 0 {
				boundPayload := inject.PrefixQuery(fastPayload,prefix,test.Where,test.Clause)
				boundPayload = inject.SuffixQuery(boundPayload,"1",suffix,test.Where)
				reqPayload := inject.Payload(place,parameter,boundPayload,1)

				method,check  := utils.GetItems(test.Response)
				check = inject.CleanupPayload(check,"")
				// 如果是报错注入
				if method == data.PAYLOAD_METHOD.GREP{
					page,_,_ := request.QueryPage(reqPayload,place)
					output := core.ExtractRegexResult(check,page)
					// 说明报错注入有结果了
					if len(output) != 0 {
						log.Info(fmt.Sprintf("\"parameter '%s' is '%s' injectable \"",parameter,test.Title))
						payload,_ := url.QueryUnescape(reqPayload)
						log.Warn(fmt.Sprintf("payload: %s",payload))
						return true
					}
				}
			}
		}

		//title := test.Title
		data.Kb.TestType = test.Stype
		//clause := test.Clause

		// union 联合注入
		if data.PAYLOAD_TECHNIQUE.UNION == test.Stype {

		}
		// 报错注入
	}
	return false
}