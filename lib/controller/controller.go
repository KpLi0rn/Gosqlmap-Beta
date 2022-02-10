package controller

import (
	"github.com/KpLi0rn/Gosqlmap/lib/core"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/log"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
)

func Start()  {
	if data.Configure.Url != ""{
		data.Kb.Targets = append(data.Kb.Targets, data.Configure.Url)
	}

	for _,uri := range data.Kb.Targets {
		core.ParseTargetUrl(uri)
		core.SetTargetEnv()

		if !CheckConnection(){
			continue
		}
		CheckWaf()

		place := data.Configure.Method
		params := data.Configure.Params[place]
		data.Configure.ParamsDict = utils.DictParams(place,params)

		// 遍历参数进行sql注入测试
		for k,v := range data.Configure.ParamsDict[place]{
			parameter := k
			value := v
			var testSqlInj = true
			// 检测参数是否为动态的 h还没做
			check := CheckDynParam(place,parameter,value)
			if !check {
				log.Error("parameter does not appear to be dynamic")
				testSqlInj = false
			}
			// 开始 SQL 注入
			if testSqlInj{
				// 启发式测试是否存在 sql 注入
				check = HeuristicCheckSqlInjection(place,parameter,value)

				// SQL 注入真正测试
				check = CheckSqlInjection(place,parameter,value)

				// 目标存在 sql 注入
				if check {

					// 进行数据库名注入
					Action()
					return
				}
			}
		}
	}
}
