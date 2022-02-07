package data

//
//var Place = map[Enums]string{
//	GET: 	"GET",
//}

type PlaceEnum struct {
	GET string
	POST string
}

var Place = PlaceEnum{
	GET: "GET",
	POST: "POST",

}


type HTTPHEADER struct {
	UserAgent string
	REFERER   string
	HOST string
	CONTENT_ENCODING string
	CONTENT_TYPE string
}

var HttpHeader = HTTPHEADER{
	UserAgent: "User-Agent",
	REFERER:   "Referer",
	HOST:      "HOST",
	CONTENT_ENCODING: "Content-Encoding",
	CONTENT_TYPE : "Content-Type",
}

var PAYLOAD_TECHNIQUE = TECHNIQUE{
	BOOLEAN: 1,
	ERROR: 2,
	QUERY: 3,
	STACKED: 4,
	TIME: 5,
	UNION: 6,
}


type TECHNIQUE struct {
	BOOLEAN  int
	ERROR int
	QUERY int
	STACKED int
	TIME int
	UNION int
}

var PAYLOAD_METHOD = METHOD{
	GREP: "Grep",
}

type METHOD struct {
	GREP string
}

