package data

//var PayloadKey = []string{"Title","Stype","Level","Risk","Clause","Where","Vector"}

var (
	Configure = Conf{
		Params: make(map[string]string),
		Headers: make(map[string]string),
		Queries: make(map[string]map[string]map[string]string),
	}
	Kb = KbObject{
		Chars: KbChars,
	}

	KbChars = Chars{
		Start: "qpbvq",
		Stop: "qpkqq",
		At: "acp",
		Space: "qzq",
		Dollar: "qgq",
		Hash_: "qoq",
	}
)

// 记录原始配置信息
type Conf struct {
	// url 相关
	Method string
	BaseUrl string
	Url string
	Scheme string
	Path   string
	Hostname string
	Port string
	Params map[string]string // get -> xxx post -> xxx
	Query string
	UserAgent string
	ParamsDict map[string]map[string]string


	// request setting
	//Hostname *url.URL
	Thread int
	Data   string
	Headers map[string]string
	Cookie string
	Timeout int

	// payloads
	Tests []JsonTest
	Boundaries []BounaryTest
	Queries map[string]map[string]map[string]string

	// 目标版本
	Dbms string

	// SQL 注入之后的结果
	ResDbms []string
}


// 记录运行过程中解析到的结果
type KbObject struct {
	Targets []string
	ReduceTests string
	TestType int

	Chars Chars

	// 保存注入成功之后的数据
	Vector string

	Prefix string
	Suffix string


	//
	Place string
	Parameter string

}
type Chars struct {
	Start string
	Stop string
	At string
	Space string
	Dollar string
	Hash_ string
}
