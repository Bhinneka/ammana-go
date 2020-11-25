package ammana

import (
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

const (
    BaseURL        = "https://paylater-dev.ammana.id/api/v1"
    ClientID       = "YOUR_CLIENT_ID"
    ClientSecret   = "YOUR_CLIENT_SECRET"
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

func TestAmmanaGetAuth(t *testing.T) {
    t.Run("Test Ammana Get Auth", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            if r.Method != "POST" {
                t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
            }

            w.Header().Set("Content-Type", "application-json; charset=utf-8")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))

        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        getAuth, errGetAuth := amm.GetAuth()
        if errGetAuth != nil {
            t.Errorf("GetAuth() returned an error: %s", errGetAuth.Error())
        }
        if len(getAuth.Message) > 0 {
            t.Errorf("GetAuth() returned an error: %s", getAuth.Message)
        }
    })
}

func TestAmmanaGetCardDetail(t *testing.T) {
    t.Run("Test Ammana Get Card Detail by Phone Number", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "GET":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","limit_amount": 500000,"usage_amount": 0,"available_amount": 500000},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))

            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        var reqCardDetail GetCardDetailRequest
        reqCardDetail.PhoneNumber = mock.Anything

        _, errGetCardDetail := amm.GetCardDetail(reqCardDetail)
        if errGetCardDetail != nil {
            t.Errorf("GetCardDetail() returned an error: %s", errGetCardDetail.Error())
        }
    })
    t.Run("Test Ammana Get Card Detail by Identity Number", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "GET":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","limit_amount": 500000,"usage_amount": 0,"available_amount": 500000},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        var reqCardDetail GetCardDetailRequest
        reqCardDetail.IdentityNumber = mock.Anything

        _, errGetCardDetail := amm.GetCardDetail(reqCardDetail)
        if errGetCardDetail != nil {
            t.Errorf("GetCardDetail() returned an error: %s", errGetCardDetail.Error())
        }
    })
    t.Run("Negative Case Query String Empty", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "GET":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","limit_amount": 500000,"usage_amount": 0,"available_amount": 500000},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        var reqCardDetail GetCardDetailRequest

        _, errGetCardDetail := amm.GetCardDetail(reqCardDetail)
        if errGetCardDetail != nil {
            assert.Error(t, errGetCardDetail)
        }
    })
}

func TestAmmanaGetOrderURL(t *testing.T) {
    orderData := []byte(`{"card_id":"7474f079-5869-4e40-9989-8439b2bc34c5","response_type": "microsite","invoice_code": "00123455678","callback_url": "http://dummy.test/ammana-finish","cancel_url": "http://dummy.test/ammana-finish","items": [{"id" : "SKU00001","name" : "SAMSUNG A51 6/64","qty" : 1,"price": 100000,"total": 100000}]}`)

    t.Run("Test Ammana Get Order URL", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            if r.Method != "POST" {
                t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
            }

            w.Header().Set("Content-Type", "application-json; charset=utf-8")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"data": {"url": "https://paylater-dev.ammana.id/merchants/b3abd5b5-e013-4cda-a6a8-d08570f557fe/orders/c1193c32-e4ce-46b4-a415-7e3db4e77000/payment?_token=eyJpdiI6ImJLeGVBVFdsMExFT01VYWF1QjlkcWc9PSIsInZhbHVlIjoicXZJS2VOdklqMnVBWHlhVGdBOEpwSU9FSXIvdmttYnk3K3ZUa0NIT1lYcUZyTFNKb0xYY2tIQWY1YThXL1hWZ01adWtVT2txaHBid2VtbXV5eXJsTnlHVDFUTnU2dVdKdnBQeHdGSGpZTEt3dW1hWFJYZjlJSmZxTG1lejBNSWFIeEZZbEtsSXRJS1RjN2RYeWpPY2YyZVFId3JqRHEyS3RPZWJnKzdlQnV0bitrUnZrV3FNNmRNOUJydXo5TTNFeXgranBLQjlVNHNWVVdDdERVWXg1L29ldUpXOTY5bTBsZlZJeko2emMrTDdXTzNhWTgvVjcwU3VxNjhKL0ExR2I1b2ZlV1BHZjRuWXRmTlB6QnpzTnZYdmFkZWUzN2ZYMmNRaWk0QkVPQkU9IiwibWFjIjoiMjcwZDZjZjU5OTllMDg2ZTY3YTczOTJiYWI0NThjNDM4M2Q0OGRlMTYzY2IzN2IyMzFjZGQ4NjY0Mzg5NzNlYiJ9&callback_url=http%3A%2F%2Fbhinneka.com%2Fpaymentv2%2Fapi%2Fnotification%2Fammana-finish&cancel_url=http%3A%2F%2Fbhinneka.com%2Fpaymentv2%2Fapi%2Fnotification%2Fammana-finish","expires_in": 3600},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        var order OrderRequest
        err := json.Unmarshal(orderData, &order)
        if err != nil {
            t.Error("Cannot Unmarshal Order JSON data")
        }

        _, errGetOrderURL := amm.GetOrderURL(order)
        if errGetOrderURL != nil {
            t.Errorf("GetOrderURL() returned an error: %s", errGetOrderURL.Error())
        }
    })
    t.Run("Test Negative Case Ammana Get Order URL", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            if r.Method != "POST" {
                t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
            }

            w.Header().Set("Content-Type", "application-json; charset=utf-8")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"message":"invalid parameter on order request"}`))

        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        var order OrderRequest
        err := json.Unmarshal(orderData, &order)
        if err != nil {
            t.Error("Cannot Unmarshal Order JSON data")
        }

        _, errGetOrderURL := amm.GetOrderURL(order)
        if errGetOrderURL != nil {
            assert.Error(t, errGetOrderURL)
        }
    })

}

func TestAmmanaGetPaymentDetail(t *testing.T) {
    t.Run("Test Ammana Get Payment Detail", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "GET":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"transaction_id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","invoice_code": "TEST","total": 0,"create_at": "2020-11-20 00:00:00","paid_at": "2020-11-20 00:00:00"},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))

            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        invoiceNumber := mock.Anything
        _, errGetPaymentDetai := amm.GetPaymentDetail(invoiceNumber)
        if errGetPaymentDetai != nil {
            t.Errorf("GetPaymentDetail() returned an error: %s", errGetPaymentDetai.Error())
        }
    })
}
func TestAmmanaSetVoid(t *testing.T) {
    t.Run("Test Ammana Set Void", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "DELETE":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"transaction_id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","invoice_code": "TEST","total": 0,"create_at": "2020-11-20 00:00:00","void_at": "2020-11-20 00:00:00"},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        invoiceNumber := mock.Anything
        _, errSetVoidOrder := amm.SetVoidOrder(invoiceNumber)
        if errSetVoidOrder != nil {
            t.Errorf("SetVoidOrder() returned an error: %s", errSetVoidOrder.Error())
        }
    })
}
func TestAmmanaSetSettlement(t *testing.T) {
    t.Run("Test Ammana Set Settlement", func(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            switch r.Method {
            case "POST":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"token_type": "Bearer","expires_in": 3600,"access_token":""},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            case "PATCH":
                w.Header().Set("Content-Type", "application-json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"data": {"transaction_id": "7474f079-5869-4e40-9989-8439b2bc34c5","status": "activated","invoice_code": "TEST","total": 0,"create_at": "2020-11-20 00:00:00","paid_at": "2020-11-20 00:00:00","settled_at": "2020-11-20 00:00:00"},"meta": {"hostname": "paylater-dev","client_ip": "180.244.234.91"}}`))
            }
        }))
        //close server
        defer ts.Close()

        //construct ammana
        amm := New(ts.URL, ClientID, ClientSecret, DefaultTimeout)
        invoiceNumber := mock.Anything
        _, errSetSettleOrder := amm.SetSettleOrder(invoiceNumber)
        if errSetSettleOrder != nil {
            t.Errorf("GetCardDetail() returned an error: %s", errSetSettleOrder.Error())
        }
    })
}
