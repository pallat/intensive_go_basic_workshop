package addstring

import "testing"

func TestAddOneAndTwo(t *testing.T) {
	// arrange
	given := "1,2"
	expected := 3

	// act
	actual := AddString(given)

	// assert
	if expected != actual {
		t.Errorf("given %q expected %d but acctual %v\n", given, expected, actual)
	}
}
