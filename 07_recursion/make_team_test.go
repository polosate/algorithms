package _7_recursion

import (
	"reflect"
	"testing"
)

func TestGroup_SetTeamSize(t *testing.T) {
	expected := []string{"ABC", "ABD", "ABE", "ACD", "ACE", "ADE", "BCD", "BCE", "BDE", "CDE"}
	res := makeTeams("", "ABCDE", 3)
	if !reflect.DeepEqual(res, expected) {
		t.Error("Want", expected, "have", res)
	}
}
