package buxfer

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const (
	host                      = "https://www.buxfer.com/api"
	endpointLogin             = host + "/login?"
	endpointAddTransaction    = host + "/transaction_add?"
	endpointEditTransaction   = host + "/transaction_edit?"
	endpointDeleteTransaction = host + "/transaction_delete?"
	endpointTransactions      = host + "/transactions?"
	endpointAccounts          = host + "/accounts?"
)

type (
	RequestParameters interface {
		makeRequest() (*http.Request, error)
	}
	LoginParameters struct {
		Email    string
		Password string
	}
	AddTransactionParameters struct {
		Token string
		Transaction
	}
	EditTransactionParameters struct {
		Token string
		Id    int64
		Transaction
	}
	DeleteTransactionParameters struct {
		Token string
		Id    int64
	}
	Transaction struct {
		Description string
		Amount      float64
		// Optional field
		AccountId string
		// Optional field
		FromAccountId string
		// Optional field
		ToAccountId string
		// YYYY/MM/DD
		//
		// Optional field
		Date string
		// Optional field
		Tags string
		// Comma-separated
		//
		// Optional field
		Type string
		// income, expense, transfer, refund, sharedBill, paidForFriend, loan
		Status string
		// cleared, pending
		//
		// Optional field

		// type=sharedBill only
		Payers      []Member
		Sharers     []Member
		IsEvenSplit bool

		// type=loan only
		LoanedBy   string
		BorrowedBy string

		// type=paidForFriend only
		PaidBy  string
		PaidFor string
	}
	Member struct {
		Email  string `json:"email"`
		Amount int    `json:"amount"`
	}
	GetTransactionsParameters struct {
		Token string
		// Optional field
		AccountId int64
		// Optional field
		AccountName string
		// Optional field
		TagId int64
		// Optional field
		TagName string
		// StartDate can be specified as "10 feb 2008", or "2008-02-10".
		//
		// Optional field
		StartDate string
		// EndDate can be specified as "10 feb 2008", or "2008-02-10".
		//
		// Optional field
		EndDate string
		// month can be specified as "feb08", "feb 08", or "feb 2008".
		//
		// Optional field
		Month string
		// Optional field
		BudgetId int64
		// Optional field
		BudgetName string
		// Optional field
		ContactId int64
		// Optional field
		ContactName string
		// Optional field
		GroupId int64
		// Optional field
		GroupName string
		// pending | reconciled | cleared
		Status string
		// paginate between results
		//
		// Optional field
		Page int64
	}
	GetAccountsParameters struct {
		Token string
	}
)

func (pars *LoginParameters) makeRequest() (*http.Request, error) {
	data := url.Values{}
	data.Add("email", pars.Email)
	data.Add("password", pars.Password)

	return http.NewRequest(http.MethodPost, endpointLogin+data.Encode(), nil)
}

func (t Transaction) genValues() url.Values {
	data := url.Values{}
	data.Add("description", t.Description)
	data.Add("amount", strconv.FormatFloat(t.Amount, 'f', 2, 64))
	if len(t.AccountId) > 0 {
		data.Add("accountId", t.AccountId)
	}
	if len(t.FromAccountId) > 0 {
		data.Add("fromAccountId", t.FromAccountId)
	}
	if len(t.ToAccountId) > 0 {
		data.Add("toAccountId", t.ToAccountId)
	}
	if len(t.Tags) > 0 {
		data.Add("tags", t.Tags)
	}
	if len(t.Tags) > 0 {
		data.Add("date", t.Date)
	}
	data.Add("type", t.Type)
	switch t.Type {
	case "sharedBill":
		if len(t.Payers) > 0 {
			if js, err := json.Marshal(t.Payers); err == nil {
				data.Add("payers", string(js))
			}
		}
		if len(t.Sharers) > 0 {
			if js, err := json.Marshal(t.Sharers); err == nil {
				data.Add("sharers", string(js))
			}
		}
		data.Add("isEvenSplit", strconv.FormatBool(t.IsEvenSplit))
	case "loan":
		if len(t.LoanedBy) > 0 {
			data.Add("loanedBy", t.LoanedBy)
		}
		if len(t.BorrowedBy) > 0 {
			data.Add("borrowedBy", t.BorrowedBy)
		}
	case "paidForFriend":
		if len(t.PaidBy) > 0 {
			data.Add("paidBy", t.PaidBy)
		}
		if len(t.PaidFor) > 0 {
			data.Add("paidFor", t.PaidFor)
		}
	}

	return data
}

func (pars *AddTransactionParameters) makeRequest() (*http.Request, error) {
	data := pars.genValues()
	data.Add("token", pars.Token)

	return http.NewRequest(http.MethodPost, endpointAddTransaction+data.Encode(), nil)
}

func (pars *EditTransactionParameters) makeRequest() (*http.Request, error) {
	data := pars.genValues()
	data.Add("token", pars.Token)
	data.Add("id", strconv.FormatInt(pars.Id, 10))

	return http.NewRequest(http.MethodPost, endpointEditTransaction+data.Encode(), nil)
}

func (pars *DeleteTransactionParameters) makeRequest() (*http.Request, error) {
	data := url.Values{}
	data.Add("token", pars.Token)
	data.Add("id", strconv.FormatInt(pars.Id, 10))

	return http.NewRequest(http.MethodPost, endpointDeleteTransaction+data.Encode(), nil)
}

func (pars *GetTransactionsParameters) makeRequest() (*http.Request, error) {
	data := url.Values{}
	data.Add("token", pars.Token)

	if pars.AccountId != 0 {
		data.Add("accountId", strconv.FormatInt(pars.AccountId, 10))
	}

	if len(pars.AccountName) > 0 {
		data.Add("accountName", pars.AccountName)
	}

	if pars.TagId != 0 {
		data.Add("tagId", strconv.FormatInt(pars.TagId, 10))
	}

	if len(pars.TagName) > 0 {
		data.Add("accountName", pars.TagName)
	}

	if len(pars.StartDate) > 0 {
		data.Add("startDate", pars.StartDate)
	}

	if len(pars.EndDate) > 0 {
		data.Add("endDate", pars.EndDate)
	}

	if len(pars.Month) > 0 {
		data.Add("month", pars.Month)
	}

	if pars.BudgetId != 0 {
		data.Add("budgetId", strconv.FormatInt(pars.BudgetId, 10))
	}

	if len(pars.BudgetName) > 0 {
		data.Add("budgetName", pars.BudgetName)
	}

	if pars.ContactId != 0 {
		data.Add("contactId", strconv.FormatInt(pars.ContactId, 10))
	}

	if len(pars.ContactName) > 0 {
		data.Add("contactName", pars.ContactName)
	}

	if pars.GroupId != 0 {
		data.Add("groupId", strconv.FormatInt(pars.GroupId, 10))
	}

	if len(pars.GroupName) > 0 {
		data.Add("groupName", pars.GroupName)
	}

	if len(pars.Status) > 0 {
		data.Add("status", pars.Status)
	}

	if pars.Page != 0 {
		data.Add("page", strconv.FormatInt(pars.Page, 10))
	}

	return http.NewRequest(http.MethodGet, endpointTransactions+data.Encode(), nil)
}

func (pars *GetAccountsParameters) makeRequest() (*http.Request, error) {
	data := url.Values{}
	data.Add("token", pars.Token)

	return http.NewRequest(http.MethodGet, endpointAccounts+data.Encode(), nil)
}
