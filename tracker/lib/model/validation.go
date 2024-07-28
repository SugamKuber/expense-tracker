package model

import "errors"

type SplitMethod string

const (
	SplitEqual      SplitMethod = "equal"
	SplitExact      SplitMethod = "exact"
	SplitPercentage SplitMethod = "percentage"
)

type ExpenseRequest struct {
	ExpenseName  string        `json:"expense_name"`
	TotalAmount  float64       `json:"total_amount"`
	Participants []Participant `json:"participants"`
	SplitMethod  SplitMethod   `json:"split_method"`
}

type Participant struct {
	UserID     int64   `json:"user_id"`
	AmountOwed float64 `json:"amount_owed,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

func ValidateAndCalculateAmounts(expenseReq *ExpenseRequest) error {
	switch expenseReq.SplitMethod {
	case SplitEqual:
		equalAmount := expenseReq.TotalAmount / float64(len(expenseReq.Participants))
		for i := range expenseReq.Participants {
			expenseReq.Participants[i].AmountOwed = equalAmount
		}
	case SplitExact:
		var total float64
		for _, participant := range expenseReq.Participants {
			total += participant.AmountOwed
		}
		if total != expenseReq.TotalAmount {
			return errors.New("total amount does not match the sum of exact amounts")
		}
	case SplitPercentage:
		var totalPercentage float64
		for _, participant := range expenseReq.Participants {
			totalPercentage += participant.Percentage
		}
		if totalPercentage != 100 {
			return errors.New("percentages do not add up to 100")
		}
		for i := range expenseReq.Participants {
			expenseReq.Participants[i].AmountOwed = (expenseReq.TotalAmount * expenseReq.Participants[i].Percentage) / 100
		}
	default:
		return errors.New("invalid split method")
	}
	return nil
}
