package algorithms

import "testing"

func TestSpongetext(t *testing.T) {
	t.Run("I need healthcare", testSpongetextFunc("I need healthcare", "i NeEd HeAlThCaRe"))
	t.Run("Don't use that weird spongebob mocking meme", testSpongetextFunc("Don't use that weird spongebob mocking meme", "dOn'T uSe ThAt WeIrD sPoNgEbOb MoCkInG mEmE"))
	t.Run("As a man, my viewpoint is—", testSpongetextFunc("As a man, my viewpoint is—", "aS a MaN, mY vIeWpOiNt Is"))
	t.Run("All lives matter", testSpongetextFunc("All lives matter", "aLl LiVeS mAtTeR"))
}

/* Spongetext will, given a string, turn it into sPoNgEtExT and return the result. 
Example:
	testSpongeText("I need healthcare") => "i NeEd hEaLtHcArE"*/
func testSpongetextFunc(input string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := Spongetext(input)
		if actual != expected {
			t.Errorf("Spongetext(\"%s\")\n\tactual:   '%s'\n\texpected: '%s'", input, actual, expected)
		}
	}
}