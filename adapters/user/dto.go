package user

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	"gitlab.id.vin/gami/ps2-gami-common/logger"
)

const (
	applicationTypeGzip = "gzip"
)

// GetTokenResponse struct define
type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// ProfileRequest struct define
type ProfileRequest struct {
	UserID string
	Type   string
	Token  string
}

// TokenRequest struct define
type TokenRequest struct {
	Scope     string `json:"scope"`
	GrantType string `json:"grant_type"`
}

// ProfileResponse struct define
type ProfileResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data struct {
		UserID          int64  `json:"id"`
		PhoneNumber     string `json:"phone_number"`
		FullName        string `json:"full_name"`
		Email           string `json:"email"`
		ReferralCode    string `json:"referral_code"`
		ReferredBy      string `json:"referred_by"`
		CSN             string `json:"csn"`
		Gender          int    `json:"gender"`
		AddressProvince string `json:"address_province"`
		AddressDistrict string `json:"address_district"`
		AddressWard     string `json:"address_ward"`
		DOB             int64  `json:"dob"`
		FullAddress     string `json:"full_address"`
	} `json:"data"`
}

type ListUserRequest struct {
	UserIDs []string
	Token   string
}

type ListUserResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data []*Profile `json:"data"`
}

type Profile struct {
	UserID          string `json:"user_id"`
	PhoneNumber     string `json:"phone_number"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	ReferralCode    string `json:"referral_code"`
	ReferredBy      string `json:"referred_by"`
	CSN             string `json:"csn"`
	Gender          int    `json:"gender"`
	AddressProvince string `json:"address_province"`
	AddressDistrict string `json:"address_district"`
	AddressWard     string `json:"address_ward"`
	FullAddress     string `json:"full_address"`
}

func encodeGetUserProfileRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, ok := req.(*ProfileRequest)

	if ok {
		r.Header.Set("Authorization", request.Token)
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("X-User-ID", request.UserID)
		r.Header.Set("X-User-ID-Type", request.Type)
		r.Header.Set("Accept-Encoding", "gzip, deflate")
	}

	return nil
}

func decodeGetUserProfileResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp ProfileResponse

	if r.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			logger.Context(ctx).Errorf("gzip encode error: %v", err)
		}
		defer func() {
			_ = reader.Close()
		}()

		data := StreamToByte(reader)
		if err = json.Unmarshal(data, &resp); err != nil {
			logger.Context(ctx).Errorf("cannot decode token response %v, raw response %v", err, string(data))
			return nil, err
		}
	} else {

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &resp); err != nil {
			logger.Context(ctx).Errorf("cannot decode token response %v, raw response %v", err, string(data))
			return nil, err
		}
	}

	if resp.Meta.Code == http.StatusOK {
		return resp, nil
	}

	if resp.Meta.Code == 404001 {
		return nil, dtos.NewAppError(dtos.UserNotFoundError)
	}

	msg := fmt.Sprintf("call user profile has problem Code: %v, Message: %v", resp.Meta.Code, resp.Meta.Message)
	return nil, errors.New(msg)
}

func encodeTokenRequest(_ context.Context, r *http.Request, req interface{}) error {
	var (
		form = url.Values{}
	)

	request, ok := req.(*TokenRequest)
	if ok {
		form.Add("grant_type", request.GrantType)
		form.Add("scope", request.Scope)
	}

	r.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	return nil
}

func decodeTokenResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp GetTokenResponse

	data, _ := ioutil.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()

	if err := json.Unmarshal(data, &resp); err != nil {
		logger.Context(ctx).Errorf("cannot decode token response %v, raw response %v", err, string(data))
		return nil, err
	}

	return resp, nil
}

// StreamToByte converter
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stream)
	return buf.Bytes()
}

func encodeListProfileRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, ok := req.(*ListUserRequest)
	if !ok {
		return errors.New("wrong request format")
	}

	r.Header.Set("Authorization", request.Token)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept-Encoding", "gzip, deflate")

	q := r.URL.Query()
	q.Add("list_user_id", strings.Join(request.UserIDs, ","))
	r.URL.RawQuery = q.Encode()

	return nil
}

func decodeListProfileResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var (
		resp   ListUserResponse
		err    error
		reader io.ReadCloser
	)

	if r.Header.Get("Content-Encoding") == applicationTypeGzip {
		reader, err = gzip.NewReader(r.Body)
		if err != nil {
			logger.Context(ctx).Errorf("gzip encode error: %v", err)
		}
		defer func() {
			_ = reader.Close()
		}()
	} else {
		reader = r.Body
	}

	data, _ := ioutil.ReadAll(reader)
	defer func() {
		_ = r.Body.Close()
	}()

	if err := json.Unmarshal(data, &resp); err != nil {
		logger.Context(ctx).Errorf("cannot decode response, raw response %v", string(data))
		return nil, err
	}

	return resp, nil
}
