package domain

// Invoice represents the structure of an invoice entity.
type Invoice struct {
    ID          string  `json:"id"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}
