package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"

	"strings"
)
var (
	proxyEntryType = "Proxy"
)

// SendMessage 发送钉钉消息
func (p *ProxyController) Reverse() {
	module := p.Ctx.Input.Param(":module")
	proxyUrl := beego.AppConfig.String(fmt.Sprintf("proxy::%s",module))
	if proxyUrl == ""{
		p.JsonError(proxyEntryType,fmt.Sprintf("proxy module %s not exists",module),StringMap{"result":"error"},true)
	}
	urlFormatter := "%s/%s%s"
	if proxyUrl[len(proxyUrl)-1:] == "/"{
		urlFormatter = "%s%s/%s"
	}
	proxyFullUrl := fmt.Sprintf(urlFormatter,proxyUrl,p.Ctx.Input.Param(":splat"),p.getAllQuery(p.Ctx.Input.URI(),"?"))
	request :=httplib.NewBeegoRequest(proxyFullUrl,p.Ctx.Request.Method)
	if p.Ctx.Input.IsPost() {
		contentType := "application/x-www-form-urlencoded"
		if p.Ctx.Input.RequestBody == nil{
			for k,v := range p.Ctx.Request.Form{
				request.Param(k,v[0])
			}
		}else{
			contentType = "application/json"
			request.Body(p.Ctx.Input.RequestBody)
		}
		request.Header("Content-Type",contentType)
	}

	resp,err := request.Response()
	if err != nil{
		p.JsonError(proxyEntryType,"proxy err",StringMap{"result":err.Error()},true)
	}
	body,err := request.String()
	if err != nil{
		p.JsonError(proxyEntryType,"proxy err",StringMap{"result":err.Error()},true)
	}
	p.LogInfo(proxyEntryType,StringMap{"proxy_url":proxyFullUrl,"method":p.Ctx.Request.Method})
	p.JsonOK(proxyEntryType,StringMap{"result":StringMap{"http_code":resp.StatusCode,"body":body}},true)
}

func (p *ProxyController) getAllQuery(url string,substr string) string{
	pos := strings.Index(url,substr)
	if pos > 0{
		return url[pos:]
	}
	return ""
}

