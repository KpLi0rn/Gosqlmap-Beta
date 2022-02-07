package parse

import (
	"flag"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
)

var (
	url string
	timeout int
	help bool
	useragent string
)
func ParserInput() {
	flag.StringVar(&url,"u","","target url")
	flag.IntVar(&timeout,"t",5,"timeout")
	flag.StringVar(&useragent,"ua","","UserAgent")
	flag.BoolVar(&help, "help", false, "help info")

	flag.Parse()

	if help {
		flag.Usage()
	}
	/**
		针对初值进行赋值
	 */

	data.Configure.Url = url
	data.Configure.Timeout = timeout
	data.Configure.UserAgent = useragent

}
