package luhn

import (
	"testing"
)

func Test_ValidLuhn(t *testing.T) {
	cc := "4111111111111111"
	if err := ValidLuhn(cc); err != nil {
		t.Errorf("Invalid Credit card(%s) - Error: '%s'", cc, err)
	} else {
		t.Log("ok")
	}
}

func Test_ValidLuhnReal(t *testing.T) {
	// I'm using generated locked credit card numbers from shopsafe
	// these are valid and unusable by anyone but the vendor they
	// are locked to.
	cc := []string{"5466365312093969", "5466365564479213", "5466365177922096"}
	for _, value := range cc {
		if err := ValidLuhn(value); err != nil {
			t.Errorf("Invalid Credit card(%s) - Error: '%s'", cc, err)
		} else {
			t.Log("ok")
		}
	}
}

func Benchmark_LuhnSpeed(b *testing.B) {
	cc := "5466365312093969"
	for i := 0; i < b.N; i++ {
		ValidLuhn(cc)
	}
}

func Test_InvalidLuhn(t *testing.T) {
	cc := "4111111111111112"
	if err := ValidLuhn(cc); err != nil {
		if err == ERR_INVALID_LUHN {
			t.Log("ok")
		} else {
			t.Error("Unexpected error %s", err)
		}
	} else {
		t.Errorf("Expected Invalid Credit card(%s) - Error: '%s'", cc, err)
	}
}

func Test_LengthShort(t *testing.T) {
	cc := "123456789012345"
	if err := ValidLuhn(cc); err != nil {
		if err == ERR_INVALID_LENGTH_SHORT {
			t.Log("ok - expected error too short")
		} else {
			t.Errorf("Invalid Credit card %s - Error: '%s'", cc, err)
		}
	} else {
		t.Errorf("expected error CC(%s) too short", cc)
	}
}

func Test_LengthLong(t *testing.T) {
	cc := "12345678901234567890"
	if err := ValidLuhn(cc); err != nil {
		if err == ERR_INVALID_LENGTH_LONG {
			t.Log("ok - expected error too short")
		} else {
			t.Errorf("Invalid Credit card(%s) - Error: '%s'", cc, err)
		}
	} else {
		t.Errorf("expected error CC(%s) too long", cc)
	}
}

func Test_Numeric(t *testing.T) {
	cc := "l234567890123456789"
	if err := ValidLuhn(cc); err != nil {
		if err == ERR_INVALID_NUMBER {
			t.Log("ok - expected invalid number")
		} else {
			t.Errorf("Invalid Credit card(%s) - Error: '%s'", cc, err)
		}
	} else {
		t.Errorf("expected errr CC(%s) invalid number error", cc)
	}
}
