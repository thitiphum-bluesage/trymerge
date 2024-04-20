package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTaxCalculator_Calculate tests the Calculate method of the TaxCalculator.
func TestTaxCalculator_Calculate(t *testing.T) {
    taxCalculator := NewTaxCalculator(0.15) // 15% tax rate

    tests := []struct {
        name      string
        amount    float64
        wantTax   float64
    }{
        {"Zero amount", 0, 0},
        {"Positive amount", 100, 15},
        {"Negative amount", -100, -15},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotTax := taxCalculator.Calculate(tt.amount)
            assert.Equal(t, tt.wantTax, gotTax, "Tax calculated incorrectly for %s", tt.name)
        })
    }
}
