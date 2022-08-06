package buxfer

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const defaultTimeout = 5 * time.Second

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
			Transport: loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}
}

func (c *Client) Login(ctx context.Context, pars LoginParameters) (*LoginResponseRoot, error) {
	res := new(LoginResponseRoot)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmes := res.Error.Message; errmes != "" {
		return nil, errors.New("Buxfer: " + errmes)
	}

	if res.Response.Status != "OK" {
		return nil, errors.New("can't get token in API response")
	}

	if res.Response.Token == "" {
		return nil, errors.New("empty token in API response")
	}

	return res, nil
}

func (c *Client) AddTransaction(ctx context.Context, pars AddTransactionParameters) (*AddTransactionResponseRoot, error) {
	res := new(AddTransactionResponseRoot)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmsg := res.Error.Message; errmsg != "" {
		return nil, errors.New("Buxfer: " + errmsg)
	}

	return res, nil
}

func (c *Client) EditTransaction(ctx context.Context, pars EditTransactionParameters) (*EditTransactionResponseRoot, error) {
	res := new(EditTransactionResponseRoot)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmsg := res.Error.Message; errmsg != "" {
		return nil, errors.New("Buxfer: " + errmsg)
	}

	return res, nil
}

func (c *Client) DeleteTransaction(ctx context.Context, pars DeleteTransactionParameters) (*DeleteTransactionResponse, error) {
	res := new(DeleteTransactionResponse)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmsg := res.Error.Message; errmsg != "" {
		return nil, errors.New("Buxfer: " + errmsg)
	}

	return res, nil
}

func (c *Client) GetTransactions(ctx context.Context, pars GetTransactionsParameters) (*GetTransactionsResponseRoot, error) {
	res := new(GetTransactionsResponseRoot)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmsg := res.Error.Message; errmsg != "" {
		return nil, errors.New("Buxfer: " + errmsg)
	}

	if res.Response.Status != "OK" {
		return nil, errors.New("can't get accounts in API response")
	}

	return res, nil
}

func (c *Client) GetAccounts(ctx context.Context, pars GetAccountsParameters) (*GetAccountsResponseRoot, error) {
	res := new(GetAccountsResponseRoot)
	err := c.doHTTP(ctx, &pars, res)
	if err != nil {
		return nil, err
	}

	if errmsg := res.Error.Message; errmsg != "" {
		return nil, errors.New("Buxfer: " + errmsg)
	}

	if res.Response.Status != "OK" {
		return nil, errors.New("can't get accounts in API response")
	}

	return res, nil
}

func (c *Client) doHTTP(ctx context.Context, pars RequestParameters, res interface{}) error {
	req, err := pars.makeRequest()
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}

	return nil
}
