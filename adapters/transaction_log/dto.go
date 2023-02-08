package transaction_log

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ActivityPayment = "PAYMENT"
const PaymentStatusSuccess = "SUCCESS"
const ActivityLogLimit = 2

type GetTransactionRequest struct {
	Filter struct {
		ActivityType struct {
			Eq string `json:"eq"`
		} `json:"activity_type"`
		TransactionStatus struct {
			Eq string `json:"eq"`
		} `json:"transaction_status"`
		CustomerID struct {
			Eq string `json:"eq"`
		} `json:"customer_id"`
	} `json:"filter"`
	Pageable struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"pageable"`
}

type GetTransactionResponse struct {
	Data []TransactionData `json:"data"`
	Meta Meta              `json:"meta"`
}

type Meta struct {
	Code   int `json:"code"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type TransactionData struct {
	TransactionID            string  `json:"transaction_id"`
	TransactionDisplayID     string  `json:"transaction_display_id,omitempty"`
	TransactionRef           string  `json:"transaction_ref"`
	TransactionAmount        float64 `json:"transaction_amount"`
	TransactionAmountCharged float64 `json:"transaction_amount_charged"`
	TotalFee                 float64 `json:"total_fee"`
	TransactionStatus        string  `json:"transaction_status"`
	TransactionCreateTime    int64   `json:"transaction_create_time"`
	TransactionUpdateTime    int64   `json:"transaction_update_time"`
	CustomerID               string  `json:"customer_id"`
	ActivityType             string  `json:"activity_type"`
	SourceType               string  `json:"source_type"`
	DestinationType          string  `json:"destination_type"`
	ServiceID                string  `json:"service_id"`
	OrderID                  string  `json:"order_id"`
	MerchantCode             string  `json:"merchant_code"`
	ChannelCode              string  `json:"channel_code,omitempty"`
}

func encodeGetTransactionLogRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*GetTransactionRequest)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeGetTransactionLogRequest(_ context.Context, r *http.Response) (interface{}, error) {
	var resp GetTransactionResponse

	data, _ := ioutil.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
