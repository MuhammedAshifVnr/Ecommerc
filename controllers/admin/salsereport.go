package admin

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

// @Summary Download Sales Reoprt
// @Description Admin can download Sales Reoprt
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param filter query string true "Filter"
// @Router /admin/salesreport [get]
func DownloadReport(c *gin.Context) {
	filter := c.Query("filter")
	var startTime, endTime time.Time

	switch filter {
	case "daily":
		startTime = time.Now().AddDate(0, 0, -1)
		endTime = time.Now()
	case "weekly":
		startTime = time.Now().AddDate(0, 0, -7)
		endTime = time.Now()
	case "montly":
		startTime = time.Now().AddDate(0, -1, 0)
		endTime = time.Now()
	case "yearly":
		startTime = time.Now().AddDate(-1, 0, 0)
		endTime = time.Now()
	case "custom":
		startStr := c.Request.FormValue("start_date")
		endStr := c.Request.FormValue("end_date")
		var err error
		startTime, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "Failed", "message": "Invalid start date format. Use YYYY-MM-DD"})
			return
		}
		endTime, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "Failed", "message": "Invalid end date format. Use YYYY-MM-DD"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "Failed", "message": "Invalid filter parameter"})
		return
	}

	reportData, total, count := generateReportData(startTime, endTime)
	pdfPath, err := generatePDF(reportData, total, count)
	if err != nil {
		c.JSON(500, gin.H{"code":500,"status":"Failed","message": "Failed to generate PDF file"})
		return
	}

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", pdfPath))
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.File(pdfPath)

}

func generateReportData(starting, ending time.Time) ([]database.OrderItems, float64, int) {
	var orders []database.OrderItems
	var totalAmount float64
	var count int
	helper.DB.Preload("Order").Preload("Order.User").Where("status!=? AND created_at BETWEEN ? AND ?", "Cancelled", starting, ending).Find(&orders)
	for _, item := range orders {
		totalAmount += item.Amount
		count++
	}
	return orders, totalAmount, count
}

func generatePDF(orders []database.OrderItems, totalAmount float64, count int) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, "Sales Report")
	pdf.Ln(10)

	headers := []string{"Order ID", "Payment Method", "Total Amount", "Status", "Order Date"}
	for _, header := range headers {
		pdf.Cell(32, 10, header)
	}
	pdf.Ln(-1)

	for _, order := range orders {
		pdf.Cell(32, 8, strconv.Itoa(int(order.OrderID)))
		pdf.Cell(32, 8, order.Order.PaymentMethod)
		pdf.Cell(32, 8, strconv.FormatFloat(order.Amount, 'f', 2, 64))
		pdf.Cell(32, 8, order.Status)
		pdf.Cell(32, 8, order.CreatedAt.Format("2006-01-02"))
		pdf.Ln(-1)
	}

	pdf.Cell(0, 10, "Total Sales Count: "+strconv.Itoa(count))
	pdf.Ln(-1)
	pdf.Cell(0, 10, "Total Amount: "+strconv.FormatFloat(totalAmount, 'f', 2, 64))

	tempFilePath := "/var/www/ashif.online/html/sales_report.pdf"


	err := pdf.OutputFileAndClose(tempFilePath)
	if err != nil {
		return "", err
	}

	return tempFilePath, nil
}
