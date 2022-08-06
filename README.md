# Buxfer API Golang SDK

Implementation SDK for [Buxfer API](https://www.buxfer.com/help/api) with Golang.

## Install 

```shell
go get -u github.com/Viquad/buxfer-api-golang
```

## Examples
### [/login](https://www.buxfer.com/help/api#login)
```go
	ctx := context.Background()
	client := buxfer.NewClient()
	loginpars := buxfer.LoginParameters{
		Email:    "youremail@gmail.com",
		Password: "yourpassword",
	}
	respLogin, err := client.Login(ctx, loginpars)
	if err != nil {
		fmt.Println("Login failed with error: ", err)
	}
```

### [/transaction_add](https://www.buxfer.com/help/api#transaction_add)
```go
	addpars := buxfer.AddTransactionParameters{
		Token:       respLogin.Response.Token,
		Transaction: buxfer.Transaction{
			Description: "100",
			Amount:      100,
			Type:        "expense",
		},
	}
	respAdd, err := client.AddTransaction(ctx, addpars)
	if err != nil {
		fmt.Println("AddTransaction failed with error: ", err)
	}
	fmt.Printf("AddTransaction success. id: %d\n", respAdd.Response.Id)
```

### [/transaction_edit](https://www.buxfer.com/help/api#transaction_edit)
```go
	editpars := buxfer.EditTransactionParameters{
		Token:       respLogin.Response.Token,
		Transaction: buxfer.Transaction{
			Description: "2",
			Amount:      200,
			Type:        "expense",
		},
	}
	respEdit, err := client.EditTransaction(ctx, editpars)
	if err != nil {
		fmt.Println("EditTransaction failed with error: ", err)
	}
	fmt.Printf("EditTransaction success. id: %d\n", respEdit.Response.Id)
```

### [/transaction_delete](https://www.buxfer.com/help/api#transaction_delete)
```go
deletepars := buxfer.DeleteTransactionParameters{
		Token: respLogin.Response.Token,
		Id:    respEdit.Response.Id,
	}
	_, err = client.DeleteTransaction(ctx, deletepars)
	if err != nil {
		fmt.Println("DelTransaction failed with error: ", err)
	}
	fmt.Printf("DelTransaction success. id: %d\n", respEdit.Response.Id)
```

### [/transactions](https://www.buxfer.com/help/api#transactions)
```go
	transactionsPars := buxfer.GetTransactionsParameters{
		Token: token,
	}

	respTransactions, err := client.GetTransactions(ctx, transactionsPars)
	if err != nil {
		log.Fatal("Get transactions failed with error: ", err)
	}
	for _, tr := range respTransactions.Response.Transactions {
		fmt.Printf("ID: %d | Description: %s | AMOUNT: %.2f | TYPE %s\n",
			tr.Id,
			tr.Description,
			tr.Amount,
			tr.Type,
		)
	}
```

### [/accounts](https://www.buxfer.com/help/api#accounts)
```go
	accountsPars := buxfer.GetAccountsParameters{
		Token: respLogin.Response.Token,
	}
	respAccounts, err := client.GetAccounts(ctx, accountsPars)
	if err != nil {
		fmt.Println("Get accounts failed with error: ", err)
	}
	for _, acc := range respAccounts.Response.Accounts {
		fmt.Printf("ID: %d | NAME: %s | BANK: %s | BALANCE: %f %s\n",
			acc.Id,
			acc.Name,
			acc.Bank,
			acc.Balance,
			acc.Currency,
		)
	}
```