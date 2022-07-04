package buxfer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	host = "https://www.buxfer.com/api"

	endpointLogin          = "/login"
	endpointAddTransaction = "/transaction_add"
)

type (
	RequestParameters interface {
		makeRequest(ctx context.Context) (*http.Request, error)
	}
	LoginParameters struct {
		Email    string
		Password string
	}
	AddTransactionParameters struct {
		Token string

		Description   string // opt
		Amount        float64
		AccountId     string // opt
		FromAccountId string // opt
		ToAccountId   string // opt
		Date          string // YYYY-MM-DD
		Tags          string // opt, comma-separated
		Type          string // income, expense, transfer, refund, sharedBill, paidForFriend, loan
		Status        string // cleared, pending

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
)

func (pars *LoginParameters) makeRequest(ctx context.Context) (*http.Request, error) {
	data := url.Values{}
	data.Add("email", pars.Email)
	data.Add("password", pars.Password)

	return http.NewRequestWithContext(ctx, http.MethodPost, host+endpointLogin, strings.NewReader(data.Encode()))
}

func (pars *AddTransactionParameters) makeRequest(ctx context.Context) (*http.Request, error) {
	data := url.Values{}
	data.Add("token", pars.Token)

	data.Add("description", pars.Description)
	data.Add("amount", strconv.FormatFloat(pars.Amount, 'E', -1, 64))
	data.Add("accountId", pars.AccountId)
	data.Add("fromAccountId", pars.FromAccountId)
	data.Add("toAccountId", pars.ToAccountId)
	data.Add("date", pars.Date)
	data.Add("tags", pars.Tags)
	data.Add("type", pars.Type)
	switch pars.Type {
	case "sharedBill":
		js, err := json.Marshal(pars.Payers)
		if err != nil {
			return nil, err
		}
		data.Add("payers", string(js))
		js, err = json.Marshal(pars.Sharers)
		if err != nil {
			return nil, err
		}
		data.Add("sharers", string(js))
		data.Add("isEvenSplit", strconv.FormatBool(pars.IsEvenSplit))
	case "loan":
		data.Add("loanedBy", pars.LoanedBy)
		data.Add("borrowedBy", pars.BorrowedBy)
	case "paidForFriend":
		data.Add("paidBy", pars.PaidBy)
		data.Add("paidFor", pars.PaidFor)
	}

	return http.NewRequestWithContext(ctx, http.MethodPost, host+endpointAddTransaction, strings.NewReader(data.Encode()))
}
