package common

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"strings"
)

func Reverse(input *context.BeegoInput,module string)(string,error){
	proxyFullUrl := getFullUrl(module,input.Param(":splat"),input.URI())
	if proxyFullUrl == ""{
		return "",errors.New(fmt.Sprintf("module %s config not found",module))
	}
	Logger.Info(map[string]interface{}{
		"url":proxyFullUrl,
		"params":input.Params(),
	},"debug")
	request :=httplib.NewBeegoRequest(proxyFullUrl,input.Method())
	if input.IsPost() {
		request.Body(input.RequestBody)
		request.Header("Content-Type",input.Header("Content-Type"))
	}
	resp,err := request.Response()
	if err != nil{
		return "",err
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "",err
	}
	return string(body),nil
}

func getFullUrl(module string,path string,uri string)string{
	proxyUrl := beego.AppConfig.String(module)
	if proxyUrl == ""{
		return ""
	}
	if path == ""{
		return fmt.Sprintf("%s%s",proxyUrl,getAllQuery(uri,"?"))
	}
	if proxyUrl[len(proxyUrl)-1:] == "/"{
		return fmt.Sprintf("%s%s/%s",proxyUrl,path,getAllQuery(uri,"?"))
	}

	return fmt.Sprintf("%s/%s%s",proxyUrl,path,getAllQuery(uri,"?"))
}

func getAllQuery(url string,substr string) string{
	pos := strings.Index(url,substr)
	if pos > 0{
		return url[pos:]
	}
	return ""
}
