package grgitscanner

import "github.com/kazekim/backend-engineer-challenge/grlib/grenums"

var rules = []ScannerRule{
	{
		Type:        "begin_word",
		RuleId:      "word01",
		Word:        "public_key",
		Description: "warning! secret 'public_key' has been detect",
		Severity:    grenums.SeverityMedium,
	},
	{
		Type:        "begin_word",
		RuleId:      "word02",
		Word:        "private_key",
		Description: "warning! secret 'private_key' has been detect",
		Severity:    grenums.SeverityHigh,
	},
}
