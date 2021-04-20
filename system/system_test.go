package system

import (
	"fmt"
	"testing"
)

func slicesEqual(t *testing.T, one []string, two []string) {
	if (one == nil) || (two == nil) {
		t.Error("an array was nil")
	}
	if len(one) != len(two) {
		t.Error("array lengths not equal")
	}
	for i := range one {
		if one[i] != two[i] {
			t.Errorf("inequality at %d", i)
		}
	}
}

func TestPrependArgument(t *testing.T) {
	firstString := "first_string"
	secondString := "second_string"
	arr := []string{secondString}
	expectedArr := []string{firstString, secondString}
	arr = PrependArgument(firstString, arr)
	slicesEqual(t, arr, expectedArr)
}

func TestGetDeviceID(t *testing.T) {
	echoName := "Echo Plus-CNX"
	expectedID := "38:F7:3D:6C:B6:7A"
	echoID, err := GetDeviceID(echoName)
	fmt.Println(echoID)

	if err != nil {
		t.Error(err)
	}
	if !(echoID == expectedID) {
		t.Error("id's not equal")
	}
}
