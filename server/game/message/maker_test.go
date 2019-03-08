package message

import "testing"

func TestMakeLoginResult(t *testing.T) {
	var id int64 = 100

	body := MakeLoginResult(id)
	loginResult := NewLoginResult(body)

	if id != loginResult.Id() {
		t.Errorf("exprected: %v, got: %v", id, loginResult.Id())
	}
}
