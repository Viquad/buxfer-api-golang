package buxfer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
)

type (
	Client struct {
		*http.Client
	}
)

func NewClient() *Client {
	return &Client{&http.Client{Timeout: defaultTimeout}}
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

	if errmes := res.Error.Message; errmes != "" {
		return nil, errors.New("Buxfer: " + errmes)
	}

	return res, nil
}

func (c *Client) doHTTP(ctx context.Context, pars RequestParameters, res interface{}) error {
	req, err := pars.makeRequest(ctx)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}

	return nil
}
