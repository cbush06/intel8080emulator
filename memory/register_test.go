package memory

import "testing"

func TestNewRegister(t *testing.T) {
	r := NewRegister(8)
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegister_Write(t *testing.T) {
	r := new(Register)
	r.Write([]byte{8})
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegister_Write8(t *testing.T) {
	r := new(Register)
	r.Write8(8)
	if r.data != 8 {
		t.Errorf("Expected 8 but got %d", r.data)
	}
}

func TestRegister_Read(t *testing.T) {
	r := &Register{
		data: 8,
	}

	buf := make([]byte, 1)
	r.Read(buf)
	if buf[0] != 8 {
		t.Errorf("Expected 8 but got %d", buf[0])
	}
}

func TestRegister_Read8(t *testing.T) {
	r := &Register{
		data: 8,
	}

	var buf uint8
	r.Read8(&buf)
	if buf != 8 {
		t.Errorf("Expected 8 but got %d", buf)
	}
}
