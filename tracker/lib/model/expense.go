package model

import (
	"sync"
	"time"
	"tracker/lib/config"
	"tracker/lib/db"
)

type MyExpense struct {
	ExpenseName string    `json:"expense_name"`
	TotalAmount float64   `json:"group_amount"`
	AmountOwed  float64   `json:"amount_owed"`
	CreatedAt   time.Time `json:"created_at"`
}

type MyExpenses struct {
	Expenses []MyExpense `json:"expenses"`
}

type AllUserExpense struct {
	UserName     string  `json:"name"`
	TotalExpense float64 `json:"total_expense"`
}

type AllUserExpenses struct {
	Expenses []AllUserExpense `json:"expenses"`
}

type Expense struct {
	ExpenseID   int64   `json:"expense_id"`
	ExpenseName string  `json:"expense_name"`
	TotalAmount float64 `json:"total_amount"`
	CreatorID   int64   `json:"creator_id"`
}

type ExpenseParticipant struct {
	ExpenseID  int64   `json:"expense_id"`
	UserID     int64   `json:"user_id"`
	AmountOwed float64 `json:"amount_owed"`
}

func CreateExpenseWithParticipants(expenseName string, totalAmount float64, creatorID int64, participants []Participant) (int64, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return 0, err
	}
	defer dbConn.Close()

	tx, err := dbConn.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var expenseID int64
	err = tx.QueryRow(
		`INSERT INTO expenses (expense_name, total_amount, creator_id)
		VALUES ($1, $2, $3)
		RETURNING expense_id`,
		expenseName, totalAmount, creatorID).Scan(&expenseID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(participants))

	for _, participant := range participants {
		wg.Add(1)
		go func(participant Participant) {
			defer wg.Done()
			_, err := tx.Exec(
				`INSERT INTO expense_tracker (expense_id, user_id, amount_owed)
				VALUES ($1, $2, $3)`,
				expenseID, participant.UserID, participant.AmountOwed)
			if err != nil {
				errChan <- err
			}
		}(participant)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		tx.Rollback()
		return 0, <-errChan
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return expenseID, nil
}

func GetMyExpenses(userID int64) (*MyExpenses, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	query := `
	SELECT e.expense_name, e.total_amount, et.amount_owed, e.created_at
	FROM users u
	JOIN expense_tracker et ON u.user_id = et.user_id
	JOIN expenses e ON et.expense_id = e.expense_id
	WHERE u.user_id = $1
	ORDER BY e.created_at DESC`

	rows, err := dbConn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []MyExpense
	for rows.Next() {
		var expense MyExpense
		err := rows.Scan(&expense.ExpenseName, &expense.TotalAmount, &expense.AmountOwed, &expense.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &MyExpenses{expenses}, nil
}

func GetAllUsersExpenses(userID int64) (*AllUserExpenses, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	query := `
	SELECT u.name, COALESCE(SUM(et.amount_owed), 0) AS total_expense
	FROM users u
	JOIN expense_tracker et ON u.user_id = et.user_id
	JOIN expenses e ON et.expense_id = e.expense_id
	JOIN users creator ON e.creator_id = creator.user_id
	WHERE creator.user_id = $1
	GROUP BY u.user_id, u.name
	ORDER BY total_expense DESC;`

	rows, err := dbConn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []AllUserExpense
	for rows.Next() {
		var expense AllUserExpense
		err := rows.Scan(&expense.UserName, &expense.TotalExpense)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &AllUserExpenses{expenses}, nil
}

func GetAdminAllExpenses() (*AllUserExpenses, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	query := `
	SELECT u.name, COALESCE(SUM(et.amount_owed), 0) AS total_expense
	FROM users u
	LEFT JOIN expense_tracker et ON u.user_id = et.user_id
	GROUP BY u.user_id, u.name
	ORDER BY total_expense DESC;`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []AllUserExpense
	for rows.Next() {
		var expense AllUserExpense
		err := rows.Scan(&expense.UserName, &expense.TotalExpense)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &AllUserExpenses{expenses}, nil
}
