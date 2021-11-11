package lib

import (
	"github.com/binganao/breachkit/config"
	"github.com/binganao/breachkit/pkg/logger"
	"os"
	"strings"
)

func veritypoc(host ParseResult) {
	for _, poc := range config.PocRules {

		var Response response

		fl, _ := os.OpenFile(OutPut, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if host.Scheme == "" {
			host.Scheme = "http://"
			veritypoc(host)
			host.Scheme = "https://"
			veritypoc(host)
			return
		}

		if host.Port == "" {
			Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			if poc.StrVerity.ContainStr != "" {
				if strings.Contains(Response.ResponseBody, poc.StrVerity.ContainStr) {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
				}
			}
		} else {
			Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			if strings.Contains(Response.ResponseBody, poc.StrVerity.ContainStr) {
				logger.Success(host.Ip + " has [" + poc.Name + "]")
				fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
			}
		}

	}
}
