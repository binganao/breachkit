package lib

import (
	"regexp"
	"strings"
)

type ParseResult struct {
	Scheme string
	Ip     string
	Port   string
}

func Parse(target string) ParseResult {
	var ParseData ParseResult
	var SchemeResult string = ""
	var IpResult string = ""
	var PortResult string = ""

	ReScheme := regexp.MustCompile(`(https?://)`)
	ReIp := regexp.MustCompile(`[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?`)
	RePort := regexp.MustCompile(`:(\d+)/?`)

	if len(ReScheme.FindAllStringSubmatch(target, -1)) != 0 {
		SchemeResult = ReScheme.FindAllStringSubmatch(target, -1)[0][0]
	}

	if len(ReIp.FindAllStringSubmatch(target, -1)) != 0 {
		IpResult = ReIp.FindAllStringSubmatch(target, -1)[0][0]
	}

	if len(RePort.FindAllStringSubmatch(target, -1)) != 0 {
		PortResult = RePort.FindAllStringSubmatch(target, -1)[0][1]
	} else {
		if strings.Contains(SchemeResult, "https") {
			PortResult = "443"
		} else {
			PortResult = "80"
		}
	}

	//PathResult := RePath.FindAllStringSubmatch(target,-1)[0][0]

	ParseData.Scheme = SchemeResult
	ParseData.Ip = IpResult
	ParseData.Port = PortResult

	return ParseData
}
