package memory

import "testing"

func TestNewRegisterPair(t *testing.T) {
	rp := NewRegisterPair(8, 16)
	if rp.High.data != 8 {
		t.Errorf("Expected 8 but got %d", rp.High.data)
	}
	if rp.Low.data != 16 {
		t.Errorf("Expected 16 but got %d", rp.Low.data)
	}
}

func TestRegisterPairWrite(t *testing.T) {
	rp := new(RegisterPair)
	buf := []byte{8, 16}
	rp.Write(buf)

	if rp.Low.data != 8 {
		t.Errorf("Expected 8 but got %d", rp.Low.data)
	}

	if rp.High.data != 16 {
		t.Errorf("Expected 16 but got %d", rp.High.data)
	}
}

func TestRegisterPairWriteLow(t *testing.T) {
	rp := new(RegisterPair)
	rp.WriteLow(8)

	if rp.Low.data != 8 {
		t.Errorf("Expected 8 but got %d", rp.Low.data)
	}
}

func TestRegisterPairWriteHigh(t *testing.T) {
	rp := new(RegisterPair)
	rp.WriteHigh(16)

	if rp.High.data != 16 {
		t.Errorf("Expected 16 but got %d", rp.High.data)
	}
}

func TestRegisterPairWrite16(t *testing.T) {
	rp := new(RegisterPair)
	rp.Write16(17544)

	if rp.High.data != 68 {
		t.Errorf("Expected 136 but got %d", rp.High.data)
	}

	if rp.Low.data != 136 {
		t.Errorf("Expected ")
	}
}

func TestRegisterPairRead(t *testing.T) {
	rp := &RegisterPair{
		Low:  *NewRegister(8),
		High: *NewRegister(16),
	}

	buf := make([]byte, 2)
	rp.Read(buf)

	if buf[0] != 8 {
		t.Errorf("Expected 8 but got %d", buf[0])
	}

	if buf[1] != 16 {
		t.Errorf("Expected 16 but got %d", buf[1])
	}
}

func TestRegisterPairReadLow(t *testing.T) {
	rp := &RegisterPair{
		Low: *NewRegister(8),
	}

	var buf uint8
	rp.ReadLow(&buf)

	if buf != 8 {
		t.Errorf("Expected 8 but got %d", buf)
	}
}

func TestRegisterPairReadHigh(t *testing.T) {
	rp := &RegisterPair{
		High: *NewRegister(8),
	}

	var buf uint8
	rp.ReadHigh(&buf)

	if buf != 8 {
		t.Errorf("Expected 8 but got %d", buf)
	}
}

func TestRegisterPairRead16(t *testing.T) {
	rp := &RegisterPair{
		Low:  *NewRegister(136),
		High: *NewRegister(68),
	}

	var buf uint16
	rp.Read16(&buf)

	if buf != 17544 {
		t.Errorf("Expected 17544 but got %d", buf)
	}
}
