package controller

import (
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/plugin/generic"
)

/**
在确认了 SQL 注入之后进行的后续操作
 */

func Action(){

	if len(data.Configure.Dbms) != 0 {
		data.Configure.ResDbms = generic.GetDbs()
	}

}
