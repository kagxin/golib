package function

import "testing"

func Test_iff(t *testing.T) {
	r := IIF(1 > 2, 1, 2).(int)
	if r != 2 {
		t.Errorf("iff error")
	}
}
