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
		Response AddTransactionResponse `json:"response"`
		Error    ErrorResponse          `json:"error"`
	}
	AddTransactionResponse struct {
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
	ErrorResponse struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}
)
