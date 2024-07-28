package handler

import (
	"bytes"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

type Expense struct {
	Name         string `json:"name"`
	TotalExpense int    `json:"total_expense"`
}

type ExpensesResponse struct {
	Expenses []Expense `json:"expenses"`
}

type MyExpense struct {
	ExpenseName string    `json:"expense_name"`
	TotalAmount float64   `json:"group_amount"`
	AmountOwed  float64   `json:"amount_owed"`
	CreatedAt   time.Time `json:"created_at"`
}

type MyExpenses struct {
	Expenses []MyExpense `json:"expenses"`
}

func createExcel(headers []string, rows [][]interface{}, buffer *bytes.Buffer) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(1), sheet)

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheet, cell, header)
	}

	for i, row := range rows {
		for j, value := range row {
			cell := fmt.Sprintf("%s%d", string(rune('A'+j)), i+2)
			f.SetCellValue(sheet, cell, value)
		}
	}

	if err := f.Write(buffer); err != nil {
		return err
	}
	return nil
}

func createExcelForTrackAll(expenses []Expense, buffer *bytes.Buffer) error {
	headers := []string{"Name", "Total Expense"}
	rows := make([][]interface{}, len(expenses))
	for i, expense := range expenses {
		rows[i] = []interface{}{expense.Name, expense.TotalExpense}
	}
	return createExcel(headers, rows, buffer)
}

func createExcelForTrackMe(expenses []MyExpense, buffer *bytes.Buffer) error {
	headers := []string{"Expense Name", "Group Amount", "Amount Owed", "Created At"}
	rows := make([][]interface{}, len(expenses))
	for i, expense := range expenses {
		rows[i] = []interface{}{
			expense.ExpenseName,
			expense.TotalAmount,
			expense.AmountOwed,
			expense.CreatedAt.Format("2006-01-02 15:04"),
		}
	}
	return createExcel(headers, rows, buffer)
}
