package makerchecker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"gitlab.id.vin/gami/gami-common/logger"
)

// MakerRequest data
type MakerRequest struct {
	ActionType string      `json:"action_type"`
	Module     string      `json:"module"`
	ObjectID   string      `json:"object_id"`
	Params     interface{} `json:"params"`
	Payload    interface{} `json:"payload"`
	UserID     int64       `json:"user_id,omitempty"`
	UserName   string      `json:"user_name,omitempty"`
}

// ParamsCampaign data
type ParamsCampaign struct {
	ID string `json:"id"`
}

// ParamsReward data
type ParamsReward struct {
	RewardID string `json:"reward_id"`
	PoolID   string `json:"pool_id"`
}

// ParamsRewardPool data
type ParamsRewardPool struct {
	ID string `json:"id"`
}

// MakerRewardPoolPayload data
type MakerRewardPoolPayload struct {
	Amount       int    `json:"amount"`
	PoolID       int64  `json:"pool_id"`
	PoolCode     string `json:"pool_code"`
	RewardCode   string `json:"reward_code"`
	ChangeAmount int    `json:"change_amount"`
}

// MakerResponse data
type MakerResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data struct {
		ID           int64       `json:"id"`
		MakerID      int         `json:"maker_id"`
		MakerName    string      `json:"maker_name"`
		Module       string      `json:"module"`
		ActionType   string      `json:"action_type"`
		ObjectID     string      `json:"object_id"`
		Payload      interface{} `json:"payload"`
		CheckerID    int         `json:"checker_id"`
		CheckerName  string      `json:"checker_name"`
		CheckedDate  int         `json:"checked_date"`
		Status       string      `json:"status"`
		Result       string      `json:"result"`
		Reason       string      `json:"reason"`
		ResponseBody interface{} `json:"response_body"`
		FailureCause string      `json:"failure_cause"`
		CreatedAt    int         `json:"created_at"`
		UpdatedAt    int         `json:"updated_at"`
		CreatedBy    int         `json:"created_by"`
		UpdatedBy    int         `json:"updated_by"`
	} `json:"data"`
}

func encodeMakerRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*MakerRequest)
	r.Header.Add("Auth-Maker-ID", strconv.FormatInt(request.UserID, 10))
	r.Header.Add("Auth-Maker-Name", request.UserName)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeMakerResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp MakerResponse

	data, _ := ioutil.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()

	err := json.Unmarshal(data, &resp)
	if err != nil {
		logger.Context(ctx).Errorf("maker checker decoding json error %v", err)
		return nil, err
	}

	if resp.Meta.Code != http.StatusOK {
		logger.Context(ctx).Errorf("maker checker raw response error %v", string(data))
		return nil, errors.New(resp.Meta.Message)
	}

	return resp, nil
}
