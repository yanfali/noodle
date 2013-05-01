package report

import (
	"mindshrub/account"
	"testing"
)

var expected = `Bert: error
Lenny: $500
Sheila: $-93
`

func Test_Summary(t *testing.T) {
	balances := map[string]*account.Balance{
		"Lenny":    &account.Balance{"Lenny", 1000, "4111111111111111", 500},
		"Sheila":   &account.Balance{"Sheila", 3000, "5454545454545454", -93},
		"Bert": &account.Balance{"Bert", 2000, "error", 0},
	}
	res := Summary(balances)
	if res == expected {
		t.Log("ok")
	} else {
		t.Errorf("Unexpected output. Expected %s  got %s", expected, res)
	}
}
