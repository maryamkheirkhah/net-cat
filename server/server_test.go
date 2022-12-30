package server

import (
	"testing"
)

func TestCheckPort(t *testing.T) {
	correctPort, err1 := CheckPort("14567")
	falsePort1, err2 := CheckPort("-45f6")
	falsePort2, err3 := CheckPort("10")
	if err1 != nil {
		t.Errorf("CheckPort failing on valid port \ngot: %v \nexpected: %v", correctPort, "14567")
	} else if err2 == nil {
		t.Errorf("CheckPort not failing on invalid input integer: %v", falsePort1)
	} else if err3 == nil {
		t.Errorf("CheckPort not failing on port number not being within valid range 1024 - 65535: %v", falsePort2)
	}
}
