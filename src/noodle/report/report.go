package report

import (
	"noodle/account"
	"bytes"
	"fmt"
	"sort"
)

func Summary(balances map[string]*account.Balance) string {
	var sorted account.Balances = make(account.Balances, len(balances))
	index := 0
	for key := range balances {
		sorted[index] = balances[key]
		index += 1
	}
	sort.Sort(account.ByName{sorted})
	buf := bytes.NewBuffer([]byte{})
	for _, b := range sorted {
		fmt.Fprintf(buf, "%s\n", b)
	}
	return buf.String()
}
