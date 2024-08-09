package tempconv

import (
    "testing"
)

func TestCToK(t *testing.T) {
    if AbsoluteZeroK != CToK(AbsoluteZeroC) {
        t.Fatalf("Conversion from Absolute Zero C to Kelvin failed")
    }

    if FreezingK != CToK(FreezingC) {
        t.Fatalf("Conversion from 0°C to Kelvin failed")
    }

    if BoilingK != CToK(BoilingC) {
        t.Fatalf("Conversion from 100°C to Kelvin failed")
    }
}

