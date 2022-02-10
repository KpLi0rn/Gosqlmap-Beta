package generic

import (
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/techniques"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
)

func GetDbs() []string{
	dbms := data.Configure.Dbms
	_,rootQuery := utils.GetInner(data.Configure.Queries[dbms],"dbs")
	// 有 information 表
	values := rootQuery["inband"]["query"]
	// 没有 information 表
	//values = rootQuery["inband"]["query1"]

	techniques.GetValues(values)

	return []string{"1"}
}

