package controllers

import (
	"devops-api/common"
	"fmt"
)
var (
	proxyEntryType = "Proxy"
)

//代理所有path透传
func (p *ProxyController) Reverse() {
	module := fmt.Sprintf("proxy::%s",p.Ctx.Input.Param(":module"))
	body,err := common.Reverse(p.Ctx.Input,module)
	if err != nil{
		p.JsonError(proxyEntryType,err.Error(),StringMap{"success":"fail"},true)
	}
	p.JsonOK(proxyEntryType,StringMap{"success":"success","body":body},true)
}

func (p *ProxyController) ReverseConfig() {
	p.Ctx.Input.SetParam(":splat","")
	module := fmt.Sprintf("self::%s",p.Ctx.Input.Param(":module"))
	body,err := common.Reverse(p.Ctx.Input,module)
	if err != nil{
		p.JsonError(proxyEntryType,err.Error(),StringMap{"success":"fail"},true)
	}
	p.JsonOK(proxyEntryType,StringMap{"success":"success","body":body},true)
}

