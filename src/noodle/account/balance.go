package account

import (
	"bytes"
	"errors"
	"fmt"
	"luhn"
	"strconv"
)

type Balance struct {
	Name    string
	Limit   int64
	Cardno  string
	Current int64
}

func (b *Balance) String() string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buf, "%s: ", b.Name)
	if b.Cardno == "error" {
		fmt.Fprintf(buf, "%s", b.Cardno)
	} else {
		fmt.Fprintf(buf, "$%d", b.Current)
	}
	return buf.String()
}

type Balances []*Balance

func (b Balances) Len() int      { return len(b) }
func (b Balances) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

type ByName struct{ Balances }

func (b ByName) Less(i, j int) bool { return b.Balances[i].Name < b.Balances[j].Name }

var ERR_UNKNOWN_COMMAND = errors.New("Unknown Command")
var ERR_UNKNOWN_ACCOUNT = errors.New("Unknown Account")

func Add(balances map[string]*Balance, col []string) error {
	if col[0] != "Add" {
		return ERR_UNKNOWN_COMMAND
	}
	b := &Balance{Name: col[1]}
	if err := luhn.ValidLuhn(col[2]); err != nil {
		b.Cardno = "error"
	} else {
		b.Cardno = col[2]
	}
	if value, err := strconv.Atoi(col[3][1:]); err != nil {
		return err
	} else {
		b.Limit = int64(value)
	}
	balances[b.Name] = b
	return nil
}

func Transaction(balances map[string]*Balance, col []string) error {
	if col[0] != "Charge" && col[0] != "Credit" {
		return ERR_UNKNOWN_COMMAND
	}
	var b *Balance
	var ok bool
	if b, ok = balances[col[1]]; !ok {
		return ERR_UNKNOWN_ACCOUNT
	}
	if b.Cardno == "error" {
		// ignore invalid luhn cards
		return nil
	}
	working := b.Current
	var value int64
	var err error
	if value, err = strconv.ParseInt(col[2][1:], 10, 64); err != nil {
		return err
	} else {
		switch col[0] {
		case "Charge":
			if working+value <= b.Limit {
				// ignore declined transactions
				b.Current = working + value
			}
		case "Credit":
			b.Current = working - value
		}
	}
	return nil
}
