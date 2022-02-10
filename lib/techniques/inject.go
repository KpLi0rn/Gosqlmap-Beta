package techniques

import (
	error2 "github.com/KpLi0rn/Gosqlmap/lib/techniques/error"
)

func GetValues(query string){
/**
	获取 payload 之后，进行替换 payload ，通过 limit 进行遍历获取结果
	我这里就简化一些好了
 */

	/**
	略去对 payload 的处理
	 */
	error2.ErrorUse(query)
}