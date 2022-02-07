package core

import (
	"fmt"
	"github.com/KpLi0rn/Gosqlmap/lib/data"
	"github.com/KpLi0rn/Gosqlmap/lib/log"
	"github.com/KpLi0rn/Gosqlmap/lib/utils"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func _getRootPath() string{
	binary, _ := os.Executable()
	root := filepath.Dir(filepath.Dir(binary))
	return root
}

func SetPath()  {
	//SQLMAP_ROOT_PATH = _getRootPath();
	data.SQLMAP_ROOT_PATH = "/Users/kpli0rn/GolandProjects/Gosqlmap"
	data.SQLMAP_DATA_PATH = filepath.Join(data.SQLMAP_ROOT_PATH,"data")
	data.SQLMAP_XML_PATH = filepath.Join(data.SQLMAP_DATA_PATH,"xml")
	data.SQLMAP_PAYLOADS_PATH = filepath.Join(data.SQLMAP_XML_PATH,"payloads")

}

func RandomInt(length int) string {
	var lenSet int
	if length != 0{
		lenSet = length
	}else {
		lenSet = 4
	}

	temp := make([]rune,lenSet)
	for i:=range temp{
		temp[i] = data.Digits[rand.Intn(len(data.Digits))]
	}
	retVal := string(temp)
	return retVal
}

func RandomStr(length int,lowercase bool, alphabet []rune) string{
	var lenSet int
	var retVal string
	if length != 0 {
		lenSet = length
	}else {
		lenSet = 4
	}
	if alphabet != nil{
		temp := make([]rune,lenSet)
		for i:=range temp{
			temp[i] = data.HEURISTIC_CHECK_ALPHABET[rand.Intn(len(data.HEURISTIC_CHECK_ALPHABET))]
		}
		retVal = string(temp)
	}else {
		temp := make([]rune,lenSet)
		for i:=range temp{
			temp[i] = data.ASCII_LETTERS[rand.Intn(len(data.ASCII_LETTERS))]
		}
		retVal = string(temp)
	}
	return retVal
}

func ParseTargetUrl(uri string){
	/**
	解析输入的 url
	 */
	if uri == "" {
		return
	}
	originUrl := uri

	// 正则判断是否为 ipv6
	ipv6Match,_ := regexp.MatchString("://\\[.+\\]",originUrl)
	if ipv6Match {
		log.Error("ipv6 is not support")
		return
	}

	// 判断是否以 http https 开头
	httpMatch,_ := regexp.MatchString("^(http)s?://",originUrl)
	if !httpMatch {
		portMatch,_ := regexp.MatchString(":443",originUrl)
		if portMatch {
			data.Configure.Url = fmt.Sprintf("https://%s",originUrl)
		}else {
			data.Configure.Url = fmt.Sprintf("http://%s",originUrl)
		}
	}

	if strings.Contains(data.Configure.Url,"?") {
		data.Configure.Url = strings.Replace(data.Configure.Url,"?",data.URI_QUESTION_MARKER,-1)
	}

	parseUrl,_ := url.Parse(data.Configure.Url)
	data.Configure.Scheme = parseUrl.Scheme
	data.Configure.Hostname = parseUrl.Hostname()
	if len(strings.Split(parseUrl.Path,data.URI_QUESTION_MARKER)) > 1 {
		data.Configure.Path = strings.Split(parseUrl.Path,data.URI_QUESTION_MARKER)[0]
		data.Configure.Query = strings.Split(parseUrl.Path,data.URI_QUESTION_MARKER)[1]
	}else {
		data.Configure.Path = parseUrl.Path
	}


	if len(parseUrl.Port()) == 0{
		if data.Configure.Scheme == "https" {
			data.Configure.Port = "443"
		}else {
			data.Configure.Port = "80"
		}
	}else {
		if utils.TransInt(parseUrl.Port()) < 1 || utils.TransInt(parseUrl.Port()) > 65535 {
			log.Error(fmt.Sprintf("invalid target URL port (%s)",parseUrl.Port()))
			return
		}else {
			data.Configure.Port = parseUrl.Port()
		}
	}


	// sqlmap 这里编码了一下
	if len(data.Configure.Query) != 0 {
		data.Configure.Url = fmt.Sprintf("%s://%s:%s%s%s%s",data.Configure.Scheme,data.Configure.Hostname,data.Configure.Port,data.Configure.Path,data.URI_QUESTION_MARKER,data.Configure.Query)
	}else {
		data.Configure.Url = fmt.Sprintf("%s://%s:%s%s",data.Configure.Scheme,data.Configure.Hostname,data.Configure.Port,data.Configure.Path)
	}
	data.Configure.Url = strings.Replace(data.Configure.Url,data.URI_QUESTION_MARKER,"?",-1)

	// map 要初始化
	data.Configure.Params[data.Place.GET] = data.Configure.Query

}

func ExtractRegexResult(check,page string) string{
	retVal := ""
	if len(check) != 0 && strings.Contains(check,"?P<result>") {
		match,_ := regexp.MatchString(check,page)
		if match {
			r,_ := regexp.Compile(check)
			retVal = r.FindString(page)
		}
	}
	return retVal
}