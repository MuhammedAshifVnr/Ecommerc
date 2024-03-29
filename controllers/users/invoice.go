package users

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoice(c *gin.Context) {
	userId := c.GetUint("userID")
	ID := c.Param("ID")
	var user database.User
	var items []database.OrderItems
	helper.DB.Where("id=?", userId).First(&user)
	helper.DB.Preload("Product").Preload("Order").Where("order_id=? AND status!=?", ID, "Cancelled").Find(&items)
	var order database.Order
	helper.DB.Preload("Coupon").Preload("Address").Where("id=?", ID).First(&order)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Invoice")
	pdf.Ln(10)
	pdf.Image("./assets/images.png", 140, 15, 50, 0, false, "", 0, "")
	pdf.SetXY(140, 35)
	pdf.SetXY(10, 20)
	pdf.CellFormat(160, 60, "E-COM", "", 0, "R", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Invoice No: INV-"+ID)
	pdf.Ln(5)
	pdf.Cell(40, 10, "Date: "+order.UpdatedAt.Format("2006-01-02"))
	pdf.Ln(5)

	pdf.Cell(40, 10, "User: "+user.Name)
	pdf.Ln(5)
	pdf.Cell(40, 10, "Email: "+user.Email)
	pdf.Ln(10)

	pdf.Cell(40, 10, "Billing Address:")
	pdf.Ln(5)
	pdf.Cell(40, 10, order.Address.Street)
	pdf.Ln(5)
	pdf.Cell(40, 10, order.Address.City)
	pdf.Ln(5)
	pdf.Cell(40, 10, order.Address.State+", "+fmt.Sprintf("%.d", order.Address.ZipCode))
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(70, 10, "Product Name", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Unit Price", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Total", "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	for _, item := range items {
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(70, 10, item.Product.ProductName, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, strconv.Itoa(int(item.Quantity)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", item.Product.ProductPrice), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", item.Amount), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(140, 10, "Sub-Total:", "", 0, "R", false, 0, "")
	pdf.Cell(140, 10, fmt.Sprintf("%.2f", order.Amount+order.Coupon.Amount))
	pdf.Ln(5)

	discount := order.Coupon.Amount
	pdf.CellFormat(140, 10, "Discount:", "", 0, "R", false, 0, "")
	pdf.Cell(140, 10, fmt.Sprintf("%.2f", discount))
	pdf.Ln(5)

	pdf.CellFormat(140, 10, "Grand Total:", "", 0, "R", false, 0, "")
	pdf.Cell(140, 10, fmt.Sprintf("%.2f", order.Amount))

	pdfPath := "/home/ashif/sales/invoce.pdf"

	if err := pdf.OutputFileAndClose(pdfPath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate PDF file"})
		return
	}

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", pdfPath))
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.File(pdfPath)

	c.JSON(200, gin.H{
		"message": "PDF file generated and sent successfully",
	})

}
