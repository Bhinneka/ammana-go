package main

import (
	"fmt"
	"github.com/Bhinneka/ammana-go"
	"time"
)

const (
	BaseURL = "https://paylater-dev.ammana.id/api/v1"
	ClientID = "YOUR-CLIENT-ID"
	ClientSecret = "YOUR-CLIENT-SECRET"
	DefaultTimeout = 10 * time.Second
)

func main() {
	newAmmana := ammana.New(BaseURL, ClientID, ClientSecret, DefaultTimeout)
	getCardDetail(newAmmana)
}

func getCardDetail(amm ammana.AmmanaService) {
	var reqCardDetail ammana.GetCardDetailRequest
	reqCardDetail.PhoneNumber = "085912347789"

	dataCardDetail, errGetCardDetail := amm.GetCardDetail(reqCardDetail)
	if errGetCardDetail != nil {
		fmt.Println(errGetCardDetail.Error())
		return
	}

	fmt.Println(dataCardDetail)
}

