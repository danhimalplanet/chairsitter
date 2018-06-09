package event

import (
	"testing"
)

func TestGotTwoArguments(t *testing.T) {
	if GotTwoArguments("two arguments") != true {
		t.Fatal("gave 2 arguments")
	}
	if GotTwoArguments("one") != false {
		t.Fatal("gave 1 arguments")
	}
	if GotTwoArguments("one two three") != false {
		t.Fatal("gave 3 arguments")
	}
	if GotTwoArguments("") != false {
		t.Fatal("gave 0 arguments")
	}
}

func TestIsNum(t *testing.T) {
	if IsNum("10000") != true {
		t.Fatal("fed it 10000")
	}
	if IsNum("twelve") != false {
		t.Fatal("fed it twelve")
	}
	if IsNum("100.00") != false {
		t.Fatal("fed it 100.00")
	}
}
