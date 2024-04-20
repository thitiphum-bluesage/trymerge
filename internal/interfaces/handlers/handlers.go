package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/thitiphum-bluesage/trymerge/internal/dto"
	"github.com/thitiphum-bluesage/trymerge/internal/service"
)

var validate = validator.New() 

// HelloHandler responds to the "/hello" route
func HelloHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World from TryMerge!")
}

func SecurePage(c echo.Context) error {
    return c.String(http.StatusOK, "You have accessed a secure area!")
}

func HandleInvoice(c echo.Context) error {
    var invoiceInput dto.InvoiceInputDTO
    if err := c.Bind(&invoiceInput); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid invoice data"})
    }

    if err := validate.Struct(invoiceInput); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    // Initialize tax calculator with a sample tax rate of 15%
    taxCalculator := service.NewTaxCalculator(0.15)
    taxAmount := taxCalculator.Calculate(invoiceInput.Amount)

    response := map[string]interface{}{
        "Invoice ID":     invoiceInput.ID,
        "Invoice Amount": invoiceInput.Amount,
        "Tax Amount":     taxAmount,
        "Total Amount":   invoiceInput.Amount + taxAmount,
        "Description":    invoiceInput.Description,
    }

    // Respond with the calculated tax
    return c.JSON(http.StatusOK, response)
}