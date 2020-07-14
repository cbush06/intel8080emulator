package alu

import "testing"

func TestConditionFlagsImpl_ClearFlags(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Zero:           true,
		Sign:           true,
		Parity:         true,
		Carry:          true,
		AuxillaryCarry: true,
	}

	cndFlags.ClearFlags()

	if cndFlags.Zero || cndFlags.Sign || cndFlags.Parity || cndFlags.Carry || cndFlags.AuxillaryCarry {
		t.Errorf("Expected all flags cleared but got %+v", cndFlags)
	}
}

func TestCreateStatusWord(t *testing.T) {
	var (
		allFlagsSet uint8 = 0xD7 // 11010111
		noFlagsSet  uint8 = 0x02 // 00000010
	)

	cndFlags := &ConditionFlagsImpl{
		Carry:          true,
		Parity:         true,
		AuxillaryCarry: true,
		Zero:           true,
		Sign:           true,
	}

	if status := cndFlags.CreateStatusWord(); status != allFlagsSet {
		t.Errorf("Expected %bb but got %bb", allFlagsSet, status)
	}

	cndFlags.ClearFlags()

	if status := cndFlags.CreateStatusWord(); status != noFlagsSet {
		t.Errorf("Expected %bb but got %bb", noFlagsSet, status)
	}
}

func TestApplyStatusWord(t *testing.T) {
	var (
		allFlagsSet uint8 = 0xD7 // 11010111
		noFlagsSet  uint8 = 0x02 // 00000010
	)

	cndFlags := new(ConditionFlagsImpl)
	cndNoFlagsSet := &ConditionFlagsImpl{
		Zero:           false,
		Sign:           false,
		Parity:         false,
		Carry:          false,
		AuxillaryCarry: false,
	}
	cndAllFlagsSet := &ConditionFlagsImpl{
		Zero:           true,
		Sign:           true,
		Parity:         true,
		Carry:          true,
		AuxillaryCarry: true,
	}

	cndFlags.ApplyStatusWord(allFlagsSet)
	if *cndFlags != *cndAllFlagsSet {
		t.Errorf("Expected %+v but got %+v", cndAllFlagsSet, cndFlags)
	}

	cndFlags.ApplyStatusWord(noFlagsSet)
	if *cndFlags != *cndNoFlagsSet {
		t.Errorf("Expected %+v but got %+v", cndNoFlagsSet, cndFlags)
	}
}

func TestIsZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}

	cndFlags.Zero = true
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}
}

func TestSetZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetZero()
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}
}

func TestUpdateZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	cndFlags.UpdateZero(0)
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}

	cndFlags.UpdateZero(1)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}
}

func TestClearZero(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Zero: true,
	}

	cndFlags.ClearZero()
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}
}

func TestIsSign(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}

	cndFlags.Sign = true
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}
}

func TestSetSign(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetSign()
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}
}

func TestUpdateSign(t *testing.T) {
	var (
		signedResult   uint8 = 0x80
		unsignedResult uint8 = 0x40
	)

	cndFlags := &ConditionFlagsImpl{
		Sign: false,
	}

	cndFlags.UpdateSign(signedResult)
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}

	cndFlags.UpdateSign(unsignedResult)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}
}

func TestClearSign(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Sign: true,
	}

	cndFlags.ClearSign()
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}
}

func TestIsParity(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.Parity = true
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestSetParity(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetParity()
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestUpdateParity(t *testing.T) {
	var (
		oddParity  uint8 = 0x57
		evenParity uint8 = 0x55
	)

	cndFlags := new(ConditionFlagsImpl)

	cndFlags.UpdateParity(oddParity)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.UpdateParity(evenParity)
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestClearParity(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Parity: true,
	}

	cndFlags.ClearParity()
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}
}

func TestIsCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.Carry = true
	if !cndFlags.IsCarry() {
		t.Error("Expected true but got false")
	}
}

func TestSetCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetCarry()
	if !cndFlags.IsCarry() {
		t.Error("Expected true but got false")
	}
}

func TestUpdateCarry(t *testing.T) {
	var (
		addend      uint8 = 0x80
		addend2     uint8 = 0x08
		sum         uint8 = 0x00
		sum2        uint8 = 0x10
		minuend     uint8 = 0x80
		minuend2    uint8 = 0x81
		difference  uint8 = 0x40
		difference2 uint8 = 0x80
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("Add w/ Carry", func(t *testing.T) {
		cndFlags.UpdateCarry(addend, sum) // 0x80 + 0x80 = 0x00 // overflow occurs as a carry out
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Add w/o Carry", func(t *testing.T) {
		cndFlags.UpdateCarry(addend2, sum2) // 0x08 + 0x08 = 0x10 // this qualifies as an axuillary carry, but not a carry
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})

	t.Run("Subtract w/ Borrow", func(t *testing.T) {
		cndFlags.UpdateCarry(minuend, difference) // 0x80 - 0x40 = 0x40 // borrow occurs from bit 7 to bit 6
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Subtract w/o Borrow", func(t *testing.T) {
		cndFlags.UpdateCarry(minuend2, difference2) // 0x81 - 0x01 = 0x80 // no borrow
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestUpdateCarryDoublePrecision(t *testing.T) {
	var (
		addend      uint16 = 0x8000
		addend2     uint16 = 0x0800
		sum         uint16 = 0x0000
		sum2        uint16 = 0x1000
		minuend     uint16 = 0x8000
		minuend2    uint16 = 0x8100
		difference  uint16 = 0x4000
		difference2 uint16 = 0x8000
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("Add with Carry", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(addend, sum) // 0x80 + 0x80 = 0x00 // overflow occurs as a carry out
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Add without Carry", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(addend2, sum2) // 0x08 + 0x08 = 0x10 // this qualifies as an auxillary carry, but not a carry
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})

	t.Run("Subtract with Borrow", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(minuend, difference) // 0x80 - 0x40 = 0x40 // borrow occurs from bit 7 to bit 6
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Subtract without Borrow", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(minuend2, difference2) // 0x81 - 0x01 = 0x80 // no borrow
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestClearCarry(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Carry: true,
	}
	cndFlags.ClearCarry()

	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}
}

func TestIsAuxillaryCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)

	if cndFlags.IsAuxillaryCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetAuxillaryCarry()
	if !cndFlags.IsAuxillaryCarry() {
		t.Error("Expected true but got false")
	}
}

func TestUpdateAuxillaryCarry(t *testing.T) {
	var (
		addend      uint8 = 0x08
		addend2     uint8 = 0x01
		sum         uint8 = 0x10
		sum2        uint8 = 0x09
		minuend     uint8 = 0x08
		minuend2    uint8 = 0x09
		difference  uint8 = 0x04
		difference2 uint8 = 0x08
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("Add with Auxillary Carry", func(t *testing.T) {
		cndFlags.UpdateAuxillaryCarry(addend, sum) // 0x08 + 0x08 = 0x10 // addition causes a carry from bit 3 into bit 4
		if !cndFlags.IsAuxillaryCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Add without Auxillary Carry", func(t *testing.T) {
		cndFlags.UpdateAuxillaryCarry(addend2, sum2) // 0x08 + 0x01 = 0x09 // this does not cause an auxillary carry
		if cndFlags.IsAuxillaryCarry() {
			t.Error("Expected false but got true")
		}
	})

	t.Run("Subtract with Auxillary Borrow", func(t *testing.T) {
		cndFlags.UpdateAuxillaryCarry(minuend, difference) // 0x10 - 0x08 = 0x08 // subtraction causing a borrow from bit 4 to bit 3
		if !cndFlags.IsAuxillaryCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Subtract without Auxillary Carry", func(t *testing.T) {
		cndFlags.UpdateAuxillaryCarry(minuend2, difference2) // 0x09 - 0x01 = 0x08 // this does not cause a borrow
		if cndFlags.IsAuxillaryCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestSetAuxillaryCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)

	if cndFlags.IsAuxillaryCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetAuxillaryCarry()
	if !cndFlags.IsAuxillaryCarry() {
		t.Error("Expected true but got false")
	}
}

func TestClearAuxillaryCarry(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		AuxillaryCarry: true,
	}

	cndFlags.ClearAuxillaryCarry()
	if cndFlags.IsAuxillaryCarry() {
		t.Error("Expected false but go true")
	}
}
