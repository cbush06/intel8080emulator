package memory

import "testing"

func TestNewRegister(t *testing.T) {
	r := NewRegister(8)
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegisterWrite(t *testing.T) {
	r := new(Register)
	r.Write([]byte{8})
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegisterWrite8(t *testing.T) {
	r := new(Register)
	r.Write8(8)
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegisterRead(t *testing.T) {
	r := &Register{
		data: 8,
	}

	buf := make([]byte, 1)
	r.Read(buf)
	if buf[0] != 8 {
		t.Errorf("Expected 8 but got %d", buf[0])
	}
}

func TestRegisterRead8(t *testing.T) {
	r := &Register{
		data: 8,
	}

	var buf uint8
	r.Read8(&buf)
	if buf != 8 {
		t.Errorf("Expected 8 but got %d", buf)
	}
}
