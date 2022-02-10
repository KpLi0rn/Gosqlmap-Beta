package data

import "encoding/xml"

/**
XML NODES
 */


// payload xml
type RequestNode struct {
	XMLName xml.Name `xml:"request"`
	Payload string   `xml:"payload"`
	Comment string   `xml:"comment"`
}

type ResponseNode struct {
	XMLName xml.Name `xml:"response"`
	Grep string		 `xml:"grep"`
}

type DetailsNode struct {
	XMLName xml.Name `xml:"details"`
	Dbms string		 `xml:"dbms"`
	Dbms_version string `xml:"dbms_version"`
}

type TestNode struct {
	XMLName 	xml.Name 	`xml:"test"`
	Title  		string 		`xml:"title"`
	Stype 		int    		`xml:"stype"`
	Level  		int    		`xml:"level"`
	Risk  		int		 	`xml:"risk"`
	Clause 		string		 `xml:"clause"`
	Where 		int		 	 `xml:"where"`
	Vector 		string		 `xml:"vector"`
	Request     RequestNode 	`xml:"request"`
	Response	ResponseNode 	`xml:"response"`
	Details     DetailsNode 	`xml:"details"`
}

type Doc struct {
	XMLName xml.Name `xml:"root"`
	Tests  []TestNode  `xml:"test"`
}

// struct
type JsonTest struct {
	Title  		string
	Stype 		int
	Level  		int
	Risk  		int
	Clause 		string
	Where 		int
	Vector 		string
	Request 	JsonRequest
	Response 	JsonResponse
	Details 	JsonDetails
}

type JsonRequest struct {
	Payload	   string
	Comment    string
}

type JsonResponse struct {
	Grep 	   string
}

type JsonDetails struct {
	Dbms			string
	Dbms_version 	string
}


/**
Parse Boundary XML
 */
type BoundaryDoc struct {
	XMLName xml.Name `xml:"root"`
	Boundarys []BoundaryNode `xml:"boundary"`
}

type BoundaryNode struct {
	XMLName xml.Name `xml:"boundary"`
	Level int `xml:"level"`
	Clause string `xml:"clause"`
	Where string `xml:"where"`
	Ptype int `xml:"ptype"`
	Prefix string `xml:"prefix"`
	Suffix string `xml:"suffix"`
}

type BounaryTest struct {
	Level int
	Clause string
	Where string
	Ptype int
	Prefix string
	Suffix string
}


// queries.xml

type QueryNode struct {
	XMLName 	xml.Name 	`xml:"root"`
	Dbms 		[]DbmsNode	`xml:"dbms"`
}

type DbmsNode struct {
	XMLName 	xml.Name 	`xml:"dbms"`
	Cast		NormalNode		`xml:"cast"`
	Length 		NormalNode		`xml:"length"`
}

type NormalNode struct {
	Query 		string 		`xml:"query,attr"`
}

type DbmsTest struct {
	Cast		NormalNode
	Length 		NormalNode
}


