package ammana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// ammana entry struick for ammana
type ammana struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	client       *ammanaHttpClient
	*logger
}

/*New Function, create ammana pointer
Required parameter :
1. Your ClientID (this from Team Ammana)
2. Your ClientSecret (this from Team Ammana)
3. BaseURL (hit to endpoint ex: https://paylater-dev.ammana.id/api/v1 for sandbox or https://paylater.ammana.id/api/v1 for production.
this value based on https://paylater-dev.ammana.id/docs/1.x/overview)
*/
func New(baseUrl string, clientId string, clientSecret string, timeout time.Duration) *ammana {
	httpRequest := newRequest(timeout)
	return &ammana{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		BaseURL:      baseUrl,
		client:       httpRequest,
		logger:       newLogger(),
	}
}

func (a *ammana) call(method string, path string, body io.Reader, v interface{}, headers map[string]string) error {
	a.info().Println("Starting http call..")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = fmt.Sprintf("%s%s", a.BaseURL, path)
	return a.client.exec(method, path, body, v, headers)
}

func (a *ammana) GetAuth() (DataAuth GenerateAccessTokenResponse, errs error) {
	a.info().Println("Starting Get Authentication Ammana")
	var (
		generateAccessTokenRequest  GenerateAccessTokenRequest
		generateAccessTokenResponse GenerateAccessTokenResponse
	)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	generateAccessTokenRequest.ClientSecret = a.ClientSecret
	payload, err := json.Marshal(generateAccessTokenRequest)
	if err != nil {
		a.error().Println(err.Error())
		return generateAccessTokenResponse, err
	}

	pathGenerateAccessToken := fmt.Sprintf("clients/%s/auth", a.ClientID)
	err = a.call("POST", pathGenerateAccessToken, bytes.NewBuffer(payload), &generateAccessTokenResponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return generateAccessTokenResponse, err
	}
	if len(generateAccessTokenResponse.Message) > 0 {
		err = errors.New(generateAccessTokenResponse.Message)
		return generateAccessTokenResponse, err
	}

	return generateAccessTokenResponse, nil
}

func (a *ammana) GetCardDetail(request GetCardDetailRequest) ServiceResult {
	a.info().Println("Starting Get Card Detail Ammana")

	// get auth data
	auth, err := a.GetAuth()
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("%s %s", auth.Data.TokenType, auth.Data.AccessToken)

	// set path
	queryStringGetCardDetail := ""
	if len(request.PhoneNumber) > 0 {
		queryStringGetCardDetail = fmt.Sprintf("phone_number=%s", request.PhoneNumber)
	} else if len(request.IdentityNumber) > 0 {
		queryStringGetCardDetail = fmt.Sprintf("identity_number=%s", request.IdentityNumber)
	} else {
		err = errors.New("invalid query string")
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	pathGetCardDetail := fmt.Sprintf("clients/%s/card?%s", a.ClientID, queryStringGetCardDetail)

	var getCardDetailResponse GetCardDetailResponse
	err = a.call("GET", pathGetCardDetail, nil, &getCardDetailResponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}
	if len(getCardDetailResponse.Message) > 0 {
		err = errors.New(getCardDetailResponse.Message)
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: getCardDetailResponse}
}

func (a *ammana) GetOrderURL(request OrderRequest) ServiceResult {
	a.info().Println("Starting Get Order URL Ammana")

	// get auth data
	auth, err := a.GetAuth()
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("%s %s", auth.Data.TokenType, auth.Data.AccessToken)

	pathGetOrderURL := fmt.Sprintf("clients/%s/cards/%s/orders", a.ClientID, request.CardID)
	//Marshal Order
	payload, errPayload := json.Marshal(request)
	if errPayload != nil {
		a.error().Println(errPayload.Error())
		return ServiceResult{Error: errPayload}
	}

	var getOrderURLResponse OrderResponse
	err = a.call("POST", pathGetOrderURL, bytes.NewBuffer(payload), &getOrderURLResponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}
	if len(getOrderURLResponse.Message) > 0 {
		err = errors.New(getOrderURLResponse.Message)
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: getOrderURLResponse}
}

func (a *ammana) GetPaymentDetail(invoiceCode string) ServiceResult {
	a.info().Println("Starting Get Payment Detail Ammana")

	// get auth data
	auth, err := a.GetAuth()
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("%s %s", auth.Data.TokenType, auth.Data.AccessToken)

	pathGetPaymentDetail := fmt.Sprintf("clients/%s/orders/%s", a.ClientID, invoiceCode)

	var getPaymentDetailReponse GetPaymentDetailResponse
	err = a.call("GET", pathGetPaymentDetail, nil, &getPaymentDetailReponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}
	if len(getPaymentDetailReponse.Message) > 0 {
		err = errors.New(getPaymentDetailReponse.Message)
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: getPaymentDetailReponse}
}

func (a *ammana) SetSettleOrder(invoiceCode string) ServiceResult {
	a.info().Println("Starting Set Settle Order Ammana")

	// get auth data
	auth, err := a.GetAuth()
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("%s %s", auth.Data.TokenType, auth.Data.AccessToken)

	pathSetSettleOrder := fmt.Sprintf("clients/%s/orders/%s/settle", a.ClientID, invoiceCode)

	var setSettleOrderResponse SetSettlePaymentResponse
	err = a.call("PATCH", pathSetSettleOrder, nil, &setSettleOrderResponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	if len(setSettleOrderResponse.Message) > 0 {
		err = errors.New(setSettleOrderResponse.Message)
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: setSettleOrderResponse}
}

func (a *ammana) SetVoidOrder(invoiceCode string) ServiceResult {
	a.info().Println("Starting Set Void Order Ammana")

	// get auth data
	auth, err := a.GetAuth()
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("%s %s", auth.Data.TokenType, auth.Data.AccessToken)

	pathSetVoidOrder := fmt.Sprintf("clients/%s/orders/%s", a.ClientID, invoiceCode)

	var setVoidOrderReponse SetVoidPaymentResponse
	err = a.call("DELETE", pathSetVoidOrder, nil, &setVoidOrderReponse, headers)
	if err != nil {
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	if len(setVoidOrderReponse.Message) > 0 {
		err = errors.New(setVoidOrderReponse.Message)
		a.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: setVoidOrderReponse}
}
