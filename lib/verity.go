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

		if poc.Logic == "bodyor" {
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			if len(poc.StrVerity.BodyStr) >= 2 {
				for _, bodystr := range poc.StrVerity.BodyStr {
					if strings.Contains(Response.ResponseBody, bodystr) {
						logger.Success(host.Ip + " has [" + poc.Name + "]")
						fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
						break
					}
				}
			}
		} else if poc.Logic == "bodyand" {
			var flag = true
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			if len(poc.StrVerity.BodyStr) >= 2 {
				for _, bodystr := range poc.StrVerity.BodyStr {
					if !strings.Contains(Response.ResponseBody, bodystr) {
						flag = false
						break
					}
				}
				if flag {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
				}
			}
		} else if poc.Logic == "headeror" {
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			for key, _ := range poc.StrVerity.HeaderStr {
				if strings.Contains(Response.ResponseHeader[key][0], poc.StrVerity.HeaderStr[key]) {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
					break
				}
			}
		} else if poc.Logic == "headerand" {
			var flag = true
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			for key, _ := range poc.StrVerity.HeaderStr {
				if !strings.Contains(Response.ResponseHeader[key][0], poc.StrVerity.HeaderStr[key]) {
					flag = false
					break
				}
			}
			if flag {
				logger.Success(host.Ip + " has [" + poc.Name + "]")
				fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
			}
		} else if poc.Logic == "bodyheaderand" {
			var flag = true
			var bodyflag = true
			var headerflag = true
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			for key, _ := range poc.StrVerity.HeaderStr {
				if !strings.Contains(Response.ResponseHeader[key][0], poc.StrVerity.HeaderStr[key]) {
					headerflag = false
					break
				}
			}
			if len(poc.StrVerity.BodyStr) >= 2 {
				for _, bodystr := range poc.StrVerity.BodyStr {
					if !strings.Contains(Response.ResponseBody, bodystr) {
						bodyflag = false
						break
					}
				}
			}
			flag = bodyflag && headerflag

			if flag {
				logger.Success(host.Ip + " has [" + poc.Name + "]")
				fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
			}
		} else if poc.Logic == "body" {
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)

			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			if poc.StrVerity.BodyStr[0] != "" {
				if strings.Contains(Response.ResponseBody, poc.StrVerity.BodyStr[0]) {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
				}
			}
		} else if poc.Logic == "header" {
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)
			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			for key, _ := range poc.StrVerity.HeaderStr {
				if strings.Contains(Response.ResponseHeader[key][0], poc.StrVerity.HeaderStr[key]) {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
					break
				}
			}
		} else {
			if host.Port == "" {
				Response = requests(poc.Method, host.Scheme+host.Ip+poc.Path, poc.Header, poc.Body)

			} else {
				Response = requests(poc.Method, host.Scheme+host.Ip+":"+host.Port+poc.Path, poc.Header, poc.Body)
			}
			if poc.StrVerity.BodyStr[0] != "" {
				if strings.Contains(Response.ResponseBody, poc.StrVerity.BodyStr[0]) {
					logger.Success(host.Ip + " has [" + poc.Name + "]")
					fl.WriteString(host.Ip + " has [" + poc.Name + "]\n")
				}
			}
		}

	}
}
