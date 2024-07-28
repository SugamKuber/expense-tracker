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

func createExcelForTrackAll(expenses []Expense, buffer *bytes.Buffer) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(1), sheet)

	headers := []string{"Name", "Total Expense"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheet, cell, header)
	}

	for i, expense := range expenses {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), expense.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), expense.TotalExpense)
	}

	if err := f.Write(buffer); err != nil {
		return err
	}
	return nil
}

func createExcelForTrackMe(expenses []MyExpense, buffer *bytes.Buffer) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(1), sheet)

	headers := []string{"Expense Name", "Group Amount", "Amount Owed", "Created At"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheet, cell, header)
	}

	for i, expense := range expenses {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), expense.ExpenseName)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), expense.TotalAmount)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), expense.AmountOwed)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), expense.CreatedAt.Format("2006-01-02 15:04"))
	}

	if err := f.Write(buffer); err != nil {
		return err
	}
	return nil
}
