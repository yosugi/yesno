package main

import "testing"

func TestCurriedDetermineYesNo(t *testing.T) {
	var expected int
	var actual int

	determineYesNoNumeric := curriedDetermineYesNo("^[0-9]$")

    actual = determineYesNoNumeric("5")
    expected = 0
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

    actual = determineYesNoNumeric("a")
    expected = 1
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

	determineYesNoLowerCase := curriedDetermineYesNo("^[a-z]$")

    actual = determineYesNoLowerCase("a")
    expected = 0
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

    actual = determineYesNoLowerCase("A")
    expected = 1
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

}

func TestCreateDetermineYesNo(t *testing.T) {
	var expected int
	var actual int

    yesNoFunc := createDetermineYesNo(true)
    actual = yesNoFunc("")
    expected = 0
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

    yesNoFunc = createDetermineYesNo(false)
    actual = yesNoFunc("")
    expected = 1
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }
}

func TestCreateShowMessage(t *testing.T) {
	var expected string
	var actual string

    actual = createShowMessage("test", false);
    expected = "test [y/N]: "
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }

    actual = createShowMessage("test", true);
    expected = "test [Y/n]: "
    if actual != expected {
        t.Errorf("actual %v\nwant %v", actual, expected)
    }
}
