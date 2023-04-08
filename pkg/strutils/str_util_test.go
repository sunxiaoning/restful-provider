package strutils

import (
	"fmt"
	"testing"
)

func TestTruncate(t *testing.T) {
	fmt.Println(Truncate("1", 2))
	fmt.Println(Truncate("1", 1))
	fmt.Println(Truncate("1", 0))
}

func TestStrs(t *testing.T) {
	if IsValidMoney("0") {
		t.Error("0 none Pass!")
	}
	if IsValidMoney("0.00") {
		t.Error("0.00 none Pass!")
	}
	if IsValidMoney("-1") {
		t.Error("0.00 none Pass!")
	}
	if IsValidMoney("aba") {
		t.Error("aba none Pass!")
	}
	if IsValidMoney("1") {
		t.Log("1 Pass!")
	}
	if IsValidMoney("1.00") {
		t.Log("1.00 pass!")
	}
}
