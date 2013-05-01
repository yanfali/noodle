package account

import (
	"testing"
)

func Test_Add(t *testing.T) {
	balances := map[string]*Balance{}
	line := []string{"Add", "Lenny", "4111111111111111", "$1000"}
	if err := Add(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Limit == 1000 && b.Cardno == "4111111111111111" && b.Current == 0 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}

func Test_Charge(t *testing.T) {
	balances := map[string]*Balance{"Lenny": &Balance{"Lenny", 1000, "4111111111111111", 0}}
	line := []string{"Charge", "Lenny", "$100"}
	if err := Transaction(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Current == 100 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}

func Test_ChargeIgnored(t *testing.T) {
	balances := map[string]*Balance{"Lenny": &Balance{"Lenny", 500, "4111111111111111", 400}}
	line := []string{"Charge", "Lenny", "$101"}
	if err := Transaction(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Current == 400 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}

func Test_ChargeEdgeCase(t *testing.T) {
	balances := map[string]*Balance{"Lenny": &Balance{"Lenny", 500, "4111111111111111", 400}}
	line := []string{"Charge", "Lenny", "$100"}
	if err := Transaction(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Current == 500 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}

func Test_Credit(t *testing.T) {
	balances := map[string]*Balance{"Lenny": &Balance{"Lenny", 500, "4111111111111111", 400}}
	line := []string{"Credit", "Lenny", "$100"}
	if err := Transaction(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Current == 300 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}

func Test_CreditNegative(t *testing.T) {
	balances := map[string]*Balance{"Lenny": &Balance{"Lenny", 500, "4111111111111111", 0}}
	line := []string{"Credit", "Lenny", "$100"}
	if err := Transaction(balances, line); err != nil {
		t.Errorf("Unexpected error: ", err)
	} else {
		b, ok := balances["Lenny"]
		if !ok {
			t.Errorf("No balance for Lenny!")
			return
		}
		if b.Name == "Lenny" && b.Current == -100 {
			t.Log("ok")
		} else {
			t.Errorf("Unexpected values: %s", b)
		}
	}
}
