package dto

// InvoiceInputDTO represents the data needed to calculate an invoice's tax.
type InvoiceInputDTO struct {
    ID          string  `json:"id" validate:"required"`
    Amount      float64 `json:"amount" validate:"required,gt=0"`
    Description string  `json:"description" validate:"required"`
}

// InvoiceResponseDTO defines the structure of the invoice response.
type InvoiceResponseDTO struct {
    InvoiceID     string  `json:"invoice_id"`
    InvoiceAmount float64 `json:"invoice_amount"`
    TaxAmount     float64 `json:"tax_amount"`
    TotalAmount   float64 `json:"total_amount"`
    Description   string  `json:"description"`
}