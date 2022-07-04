# Buxfer API Golang SDK
## https://www.buxfer.com/help/api

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/Viquad/buxfer-api-golang"
)

func main() {
	ctx := context.Background()
	client := buxfer.NewClient()
	loginpars := buxfer.LoginParameters{
		Email:    "youremail@gmail.com",
		Password: "yourpassword",
	}
	respLogin, err := client.Login(ctx, loginpars)
	token := respLogin.Response.Token
	if err != nil {
		fmt.Println("Login failed with error: ", err)
		return
	}
	fmt.Printf("Token: %s\n", token)
	addpars := buxfer.AddTransactionParameters{
		Token:       token,
		Description: "test",
		Amount:      100,
		Date:        "2022/07/04",
		Type:        "expense",
		Status:      "cleared",
	}
	respAdd, err := client.AddTransaction(ctx, addpars)
	if err != nil {
		fmt.Println("AddTransaction failed with error: ", err)
		return
	}
	fmt.Printf("AddTransaction success. id: %d\n", respAdd.Response.Id)
}
```