package service

// TaxCalculator represents the logic to calculate tax.
type TaxCalculator struct {
    TaxRate float64
}

// NewTaxCalculator creates a new instance of TaxCalculator.
func NewTaxCalculator(taxRate float64) *TaxCalculator {
    return &TaxCalculator{TaxRate: taxRate}
}

// Calculate applies the tax rate to a given amount.
func (tc *TaxCalculator) Calculate(amount float64) float64 {
    return amount * tc.TaxRate
}
