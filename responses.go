package buxfer

type (
	LoginResponseRoot struct {
		Response LoginResponse `json:"response"`
		Error    ErrorResponse `json:"error"`
	}
	LoginResponse struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
	AddTransactionResponseRoot struct {
		Response TransactionResponse `json:"response"`
		Error    ErrorResponse       `json:"error"`
	}
	EditTransactionResponseRoot struct {
		Response TransactionResponse `json:"response"`
		Error    ErrorResponse       `json:"error"`
	}
	DeleteTransactionResponse struct {
		Error ErrorResponse `json:"error"`
	}
	TransactionResponse struct {
		Id            int64    `json:"id"`
		Description   string   `json:"description"`
		Date          string   `json:"date"`
		Type          string   `json:"type"`
		Amount        float64  `json:"amount"`
		ExpenseAmount float64  `json:"expenseAmount"`
		AccountId     int64    `json:"accoundId"`
		AccountName   string   `json:"accountName"`
		Tags          string   `json:"tags"`
		TagNames      []string `json:"tagNames"`
		Status        string   `json:"status"`
		SortDate      string   `json:"sortDate"`
	}
	GetTransactionsResponseRoot struct {
		Response TransactionsResponse `json:"response"`
		Error    ErrorResponse        `json:"error"`
	}
	TransactionsResponse struct {
		Status          string                `json:"status"`
		NumTransactions string                `json:"numTransactions"`
		Transactions    []TransactionResponse `json:"transactions"`
	}
	GetAccountsResponseRoot struct {
		Response AccountsResponse `json:"response"`
		Error    ErrorResponse    `json:"error"`
	}
	AccountsResponse struct {
		Status   string    `json:"status"`
		Accounts []Account `json:"accounts"`
	}
	Account struct {
		Id         int64   `json:"id"`
		Name       string  `json:"name"`
		Bank       string  `json:"bank"`
		Balance    float64 `json:"balance"`
		Currency   string  `json:"currency"`
		LastSynced string  `json:"lastSynced"`
	}
	ErrorResponse struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}
)
