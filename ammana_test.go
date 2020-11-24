package ammana

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	BaseURL = "https://paylater-dev.ammana.id/api/v1"
	ClientID = "da89901f-4b2e-438b-b707-f93645b39631"
	ClientSecret = "3b5b8d854e53662d91a615d7705d61fe"
	DefaultTimeout = 10 * time.Second
)

func TestAmmanaCreation(t *testing.T) {
	t.Run("Test Ammana Creation", func(t *testing.T) {

		//construct ammana
		//and add required parameters
		amm := New(BaseURL, ClientID, ClientSecret, DefaultTimeout)

		if amm == nil {
			t.Error("Cannot Call Ammana Constructor")
		}
	})
}

func  TestAmmanaGetAuth(t *testing.T) {
	t.Run("Test Ammana Get Auth", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/auth" {
				t.Errorf("Expected request to ‘/auth, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
								"data": {
									"token_type": "Bearer",
									"expires_in": 3600,
									"access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzUxMiJ9.eyJzdWIiOiJkYTg5OTAxZi00YjJlLTQzOGItYjcwNy1mOTM2NDViMzk2MzEiLCJtZXJjaGFudCI6ImIzYWJkNWI1LWUwMTMtNGNkYS1hNmE4LWQwODU3MGY1NTdmZSIsImV4cCI6MTYwNjEyMzI2MX0.Y4TxmRj0S0hXvEaXS_njqLdPh4XCRbVMOcrM3PnDqEcSlBckHcbhOJZ_9t2NXWNSo4l2VFkbwVoI1SQqCwOqTR8kgrVEmieIEuXcVCjWWAvkU1_-xIJTtv4gbpdxzdekWwJx4X4FaGjrtFJ7iE8gRy18b3BOceXmu-h2QtUgyk_DDNKr4QAqImw4C2b5NB2VwHEzV2DNTCDJpnIrewEYu_WpvyzoG9JQEXLMJgmLlV0jg4MKWs7RClQiHCBG3Jl91k3t0RuCG79okVX4rYuGrhVe5DhGQrMFL91Ud-SqrbBThRRfVDCuYUzwhCf4SFsgzP1f3gTFujYHwoa-UY7yrPVJPapC6VoW1BBft46lMP7nHc8FgTQCrOcldUQRhV9DKpd_BPIddOcGCgrCjmwGgJ63QkWgLArVmDCPUzEaps_r294Por6OpiV4o8eWaBq7aurLwj9E-6TlsOXRaFrRlr4K3iL3GGtoO79KjVttPX54HgPb-xhkjHIA22pfIlztOJAuai2yV3Ws8aDhZMYhhP_T8kYg48bFv2BMYu43mfFuOGlDSY4ZyxuLHXfdT_X2hnJIvuo7DZA9EkqzEt35qOqEk-2CZ4Q0HiwsIWYOIaRFEUz0B4EdfGQmhZxOka0559hHj5c99akP2pDkrK9ZLv5gyzOdrQ28KfDZBEdVWtA"
								},
								"meta": {
									"hostname": "paylater-dev",
									"client_ip": "180.244.234.91"
								}
							}`))

		}))
		//close server
		defer ts.Close()

		//construct ammana
		amm := New(BaseURL, ClientID, ClientSecret, DefaultTimeout)
		getAuth, errGetAuth := amm.GetAuth()
		if errGetAuth != nil {
			t.Errorf("GetAuth() returned an error: %s", errGetAuth.Error())
		}
		if len(getAuth.Message) > 0 {
			t.Errorf("GetAuth() returned an error: %s", getAuth.Message)
		}
	})
}

func  TestAmmanaGetCardDetail(t *testing.T) {
	t.Run("Test Ammana Get Card Detail", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "GET" {
				t.Errorf("Expected ‘GET’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/card" {
				t.Errorf("Expected request to ‘/card, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
								"data": {
									"id": "7474f079-5869-4e40-9989-8439b2bc34c5",
									"status": "activated",
									"limit_amount": 500000,
									"usage_amount": 0,
									"available_amount": 500000
								},
								"meta": {
									"hostname": "paylater-dev",
									"client_ip": "180.244.234.91"
								}
							}`))

		}))
		//close server
		defer ts.Close()

		//construct ammana
		amm := New(BaseURL, ClientID, ClientSecret, DefaultTimeout)
		var reqCardDetail GetCardDetailRequest
		reqCardDetail.PhoneNumber = "085912347789"

		getCardDetail := amm.GetCardDetail(reqCardDetail)
		if getCardDetail.Error != nil {
			t.Errorf("GetCardDetail() returned an error: %s", getCardDetail.Error.Error())
		}
		getCardDetailData, ok := getCardDetail.Result.(GetCardDetailResponse)
		if !ok {
			t.Error("Response GetCardDetail() not GetCardDetailResponse.")
		}
		if len(getCardDetailData.Message) > 0 {
			t.Errorf("GetCardDetail() returned an error: %s", getCardDetailData.Message)
		}
	})
}

func  TestAmmanaGetOrderURL(t *testing.T) {
	orderData := []byte(`{
						"card_id":"7474f079-5869-4e40-9989-8439b2bc34c5",
						"response_type": "microsite",
						"invoice_code": "00123455678",
						"callback_url": "http://bhinneka.com/paymentv2/api/notification/ammana-finish",
						"cancel_url": "http://bhinneka.com/paymentv2/api/notification/ammana-finish",
						"items": [{
							"id" : "SKU00001",
							"name" : "SAMSUNG A51 6/64",
							"qty" : 1,
							"price": 100000,
							"total": 100000
						}]
					}`)

	t.Run("Test Ammana Get Order URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/cards" {
				t.Errorf("Expected request to ‘/card, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
								"data": {
									"url": "https://paylater-dev.ammana.id/merchants/b3abd5b5-e013-4cda-a6a8-d08570f557fe/orders/c1193c32-e4ce-46b4-a415-7e3db4e77000/payment?_token=eyJpdiI6ImJLeGVBVFdsMExFT01VYWF1QjlkcWc9PSIsInZhbHVlIjoicXZJS2VOdklqMnVBWHlhVGdBOEpwSU9FSXIvdmttYnk3K3ZUa0NIT1lYcUZyTFNKb0xYY2tIQWY1YThXL1hWZ01adWtVT2txaHBid2VtbXV5eXJsTnlHVDFUTnU2dVdKdnBQeHdGSGpZTEt3dW1hWFJYZjlJSmZxTG1lejBNSWFIeEZZbEtsSXRJS1RjN2RYeWpPY2YyZVFId3JqRHEyS3RPZWJnKzdlQnV0bitrUnZrV3FNNmRNOUJydXo5TTNFeXgranBLQjlVNHNWVVdDdERVWXg1L29ldUpXOTY5bTBsZlZJeko2emMrTDdXTzNhWTgvVjcwU3VxNjhKL0ExR2I1b2ZlV1BHZjRuWXRmTlB6QnpzTnZYdmFkZWUzN2ZYMmNRaWk0QkVPQkU9IiwibWFjIjoiMjcwZDZjZjU5OTllMDg2ZTY3YTczOTJiYWI0NThjNDM4M2Q0OGRlMTYzY2IzN2IyMzFjZGQ4NjY0Mzg5NzNlYiJ9&callback_url=http%3A%2F%2Fbhinneka.com%2Fpaymentv2%2Fapi%2Fnotification%2Fammana-finish&cancel_url=http%3A%2F%2Fbhinneka.com%2Fpaymentv2%2Fapi%2Fnotification%2Fammana-finish",
									"expires_in": 3600
								},
								"meta": {
									"hostname": "paylater-dev",
									"client_ip": "180.244.234.91"
								}
							}`))

		}))
		//close server
		defer ts.Close()

		//construct ammana
		amm := New(BaseURL, ClientID, ClientSecret, DefaultTimeout)
		var order OrderRequest
		err := json.Unmarshal(orderData, &order)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		getOrderURL := amm.GetOrderURL(order)
		if getOrderURL.Error != nil {
			t.Errorf("GetOrderURL() returned an error: %s", getOrderURL.Error.Error())
		}
		getOrderURLData, ok := getOrderURL.Result.(OrderResponse)
		if !ok {
			t.Error("Response GetOrderURL() not OrderResponse.")
		}
		if len(getOrderURLData.Message) > 0 {
			t.Errorf("GetOrderURL() returned an error: %s", getOrderURLData.Message)
		}
	})
}

