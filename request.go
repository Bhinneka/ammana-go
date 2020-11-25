package ammana

type GenerateAccessTokenRequest struct {
    ClientSecret string `json:"client_secret"`
}

type OrderRequest struct {
    CardID       string      `json:"card_id"`
    ResponseType string      `json:"response_type"`
    InvoiceCode  string      `json:"invoice_code"`
    CallbackUrl  string      `json:"callback_url"`
    CancelUrl    string      `json:"cancel_url"`
    Items        []OrderItem `json:"items"`
}

type OrderItem struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Qty   int    `json:"qty"`
    Price int    `json:"price"`
    Total int    `json:"total"`
}

type GetCardDetailRequest struct {
    PhoneNumber    string `json:"phone_number"`
    IdentityNumber string `json:"identity_number"`
}

type SharedRequest struct {
    ClientId    string `json:"client_id"`
    InvoiceCode string `json:"invoice_code"`
    PhoneNumber string `json:"phone_number"`
}
