package govincheck_test

import (
	"fmt"
	"testing"

	"github.com/gpascual2/govincheck"
)

func init() {
	fmt.Println("\n>>>> govincheck_test : Init")
}

// Test VIN validity
func TestValidateVIN(t *testing.T) {
	fmt.Println("\n>> govincheck : TestValidateVIN")

	var testsValues = []struct {
		vin     string // input
		valid   bool   // expected validity flag
		correct string // expected valid vin
	}{
		{"8T6BS48F6B7170309", true, "8T6BS48F6B7170309"},
		{"8T6BS48F0B7170309", false, "8T6BS48F6B7170309"},
		{"8T6BS48F0B7170300", false, "8T6BS48FXB7170300"},
		{"8T6BS48FXB7170306", false, "8T6BS48F0B7170306"},
		{"8T6BS48F0B7170301", false, "8T6BS48F1B7170301"},
		{"8T6BS48F0B7170307", false, "8T6BS48F2B7170307"},
		{"8T6BS48F0B7170302", false, "8T6BS48F3B7170302"},
		{"8T6BS48F0B7170308", false, "8T6BS48F4B7170308"},
		{"8T6BS48F0B7170303", false, "8T6BS48F5B7170303"},
		{"8T6BS48F0B7170304", false, "8T6BS48F7B7170304"},
		{"8T6BS48F0B7170330", false, "8T6BS48F8B7170330"},
		{"8T6BS48F0B7170305", false, "8T6BS48F9B7170305"},
	}
	for _, tt := range testsValues {
		isValid, correctVin := govincheck.VinCheck(tt.vin)
		if isValid != tt.valid {
			t.Errorf("  - Debug :: ValidateVIN(%v): expected %t, actual %t", tt.vin, tt.valid, isValid)
		}
		if correctVin != tt.correct {
			t.Errorf("  - Debug :: ValidateVIN(%v): expected %v, actual %v", tt.vin, tt.correct, correctVin)
		}
	}
}
