package gami_service

import (
	"context"

	"gitlab.id.vin/gami/gami-proto/grpc_client"
	gamiProtobuf "gitlab.id.vin/gami/gami-proto/pb"
	"gitlab.id.vin/gami/ps2-gami-common/errors"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
)

// Adapter handles all APIs for calling to rewards service.
type Adapter interface {
	GetRewardByID(ctx context.Context, ID int64) (*GetRewardResponse, error)
	GetRewardByCampaignID(ctx context.Context, campaignID int64) (*GetRewardPoolResponse, error)
	GetListRewardByIDs(ctx context.Context, req GetListRewardByIDsRequest) (*ListRewardResponse, error)
	Redeem(ctx context.Context, request *RedeemRequest) (*RedeemResponse, error)
	GetTransaction(ctx context.Context, request *TransactionRequest) (*TransactionResponse, error)
	GetCampaignByID(ctx context.Context, id int64) (*GetCampaignResponse, error)
	GetCampaignByUser(ctx context.Context, id, userID int64) (*GetCampaignResponse, error)
	GetCampaignByCode(ctx context.Context, code string, userID int64) (*GetCampaignResponse, error)
	GetRuleByCampaign(ctx context.Context, campaignID int64) (*GetRulesResponse, error)
	ListCampaigns(ctx context.Context, req ListCampaignsRequest) (*ListCampaignsResponse, error)
	GetListCampaignByIDs(ctx context.Context, req GetListCampaignByIDsRequest) (*ListCampaignsResponse, error)
	AcceptChallenge(ctx context.Context, req AcceptChallengeRequest) (*AcceptChallengeResponse, error)
	MaintenanceStatusCampaign(ctx context.Context, req MaintenanceStatusCampaignRequest) (*MaintenanceStatusCampaignResponse, error)
	GetRewardAmountInCampaign(ctx context.Context, req RewardAmountStatisticRequest) (*RewardAmountStatisticResponse, error)
	GetBlackWhiteList(ctx context.Context, campaignID int64) (*GetBlackWhiteListResponse, error)
	GetRewardsByCampaignID(ctx context.Context, req GetRewardsByCampaignIDRequest) (*GetRewardsByCampaignIDResponse, error)
	GetCampaignByUserV2(ctx context.Context, id, userID int64) (*GetCampaignResponse, error)
}

type adapter struct {
	ctx                                   context.Context
	GetRewardEndpoint                     endpoint.Endpoint
	GetRewardByCampaignIDEndpoint         endpoint.Endpoint
	GetListRewardByIDsEndpoint            endpoint.Endpoint
	RedeemEndpoint                        endpoint.Endpoint
	GetTransactionEndpoint                endpoint.Endpoint
	GetCampaignByIDEndpoint               endpoint.Endpoint
	GetCampaignByUserEndpoint             endpoint.Endpoint
	GetCampaignByCodeEndpoint             endpoint.Endpoint
	GetRuleEndpoint                       endpoint.Endpoint
	ListCampaignsEndpoint                 endpoint.Endpoint
	GetListCampaignByIDsEndPoint          endpoint.Endpoint
	AcceptChallengeEndpoint               endpoint.Endpoint
	MaintenanceStatusCampaignEndpoint     endpoint.Endpoint
	CampaignRewardAmountStatisticEndpoint endpoint.Endpoint
	GetBlackWhiteListEndpoint             endpoint.Endpoint
	GetRewardsByCampaignIDEndpoint        endpoint.Endpoint
	GetCampaignByUserV2Endpoint           endpoint.Endpoint
}

func (a *adapter) GetBlackWhiteList(ctx context.Context, campaignID int64) (*GetBlackWhiteListResponse, error) {
	response, err := a.GetBlackWhiteListEndpoint(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	r, ok := response.(*GetBlackWhiteListResponse)
	if !ok {
		return nil, errors.New(errors.ErrorCode{})
	}
	return r, nil
}

func (a *adapter) GetRewardAmountInCampaign(ctx context.Context, req RewardAmountStatisticRequest) (*RewardAmountStatisticResponse, error) {
	response, err := a.CampaignRewardAmountStatisticEndpoint(ctx, &req)
	if err != nil {
		return nil, err
	}
	r, ok := response.(*RewardAmountStatisticResponse)
	if !ok {
		return nil, errors.New(errors.ErrorCode{})
	}
	return r, nil
}

func (a *adapter) AcceptChallenge(ctx context.Context, req AcceptChallengeRequest) (*AcceptChallengeResponse, error) {
	response, err := a.AcceptChallengeEndpoint(ctx, &req)
	if err != nil {
		return nil, err
	}
	r, ok := response.(*AcceptChallengeResponse)
	if !ok {
		return nil, errors.New(errors.ErrorCode{})
	}
	return r, nil
}

func (a *adapter) GetTransaction(ctx context.Context, request *TransactionRequest) (*TransactionResponse, error) {
	response, err := a.GetTransactionEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	r := response.(*TransactionResponse)
	return r, nil
}

func (a *adapter) GetRewardByID(ctx context.Context, ID int64) (*GetRewardResponse, error) {
	response, err := a.GetRewardEndpoint(ctx, ID)
	if err != nil {
		return nil, err
	}

	r := response.(*GetRewardResponse)
	return r, nil
}

func (a *adapter) GetRewardByCampaignID(ctx context.Context, campaignID int64) (*GetRewardPoolResponse, error) {
	response, err := a.GetRewardByCampaignIDEndpoint(ctx, GetRewardPoolRequest{
		CampaignID: campaignID,
	})
	if err != nil {
		return nil, err
	}

	r := response.(*GetRewardPoolResponse)
	return r, nil
}

func (a *adapter) GetListRewardByIDs(ctx context.Context, req GetListRewardByIDsRequest) (*ListRewardResponse, error) {
	response, err := a.GetListRewardByIDsEndpoint(ctx, &req)
	if err != nil {
		return nil, err
	}
	r := response.(*ListRewardResponse)
	return r, nil
}

func (a *adapter) Redeem(ctx context.Context, request *RedeemRequest) (*RedeemResponse, error) {
	response, err := a.RedeemEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	r := response.(*RedeemResponse)
	return r, nil
}

func (a *adapter) GetCampaignByID(ctx context.Context, id int64) (*GetCampaignResponse, error) {
	response, err := a.GetCampaignByIDEndpoint(ctx, GetCampaignRequest{ID: id})
	if err != nil {
		return nil, err
	}

	r := response.(*GetCampaignResponse)
	return r, nil
}

func (a *adapter) GetCampaignByUser(ctx context.Context, id, userID int64) (*GetCampaignResponse, error) {
	response, err := a.GetCampaignByUserEndpoint(ctx, GetCampaignUserRequest{ID: id, UserID: userID})
	if err != nil {
		return nil, err
	}

	r, _ := response.(*GetCampaignResponse)
	return r, nil
}

func (a *adapter) GetCampaignByCode(ctx context.Context, code string, userID int64) (*GetCampaignResponse, error) {
	response, err := a.GetCampaignByCodeEndpoint(ctx, GetCampaignUserRequest{Code: code, UserID: userID})
	if err != nil {
		return nil, err
	}

	r := response.(*GetCampaignResponse)
	return r, nil
}

func (a *adapter) GetRuleByCampaign(ctx context.Context, campaignID int64) (*GetRulesResponse, error) {
	response, err := a.GetRuleEndpoint(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	r := response.(*GetRulesResponse)
	return r, nil
}

func (a *adapter) ListCampaigns(ctx context.Context, req ListCampaignsRequest) (*ListCampaignsResponse, error) {
	response, err := a.ListCampaignsEndpoint(ctx, &req)
	if err != nil {
		return nil, err
	}

	r := response.(*ListCampaignsResponse)
	return r, nil
}

func (a *adapter) GetListCampaignByIDs(ctx context.Context, req GetListCampaignByIDsRequest) (*ListCampaignsResponse, error) {
	response, err := a.GetListCampaignByIDsEndPoint(ctx, &req)
	if err != nil {
		return nil, err
	}
	r := response.(*ListCampaignsResponse)
	return r, nil
}

func (a *adapter) MaintenanceStatusCampaign(ctx context.Context, req MaintenanceStatusCampaignRequest) (*MaintenanceStatusCampaignResponse, error) {
	response, err := a.MaintenanceStatusCampaignEndpoint(ctx, &req)
	if err != nil {
		return nil, err
	}

	r := response.(*MaintenanceStatusCampaignResponse)
	return r, nil
}

func (a *adapter) GetRewardsByCampaignID(ctx context.Context, req GetRewardsByCampaignIDRequest) (*GetRewardsByCampaignIDResponse, error) {
	response, err := a.GetRewardsByCampaignIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	r := response.(*GetRewardsByCampaignIDResponse)
	return r, nil
}

func (a *adapter) GetCampaignByUserV2(ctx context.Context, id, userID int64) (*GetCampaignResponse, error) {
	response, err := a.GetCampaignByUserV2Endpoint(ctx, GetCampaignUserRequest{ID: id, UserID: userID})
	if err != nil {
		return nil, err
	}

	r, _ := response.(*GetCampaignResponse)
	return r, nil
}

// NewAdapter returns a new instance of NewAdapter.
func NewAdapter(ctx context.Context, connection *grpc.ClientConn) Adapter {
	return &adapter{
		ctx: ctx,
		GetCampaignByIDEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.CampaignService",
			"GetByID",
			encodeGetCampaignRequest,
			decodeGetCampaignResponse,
			gamiProtobuf.GetCampaignResponse{},
		).Endpoint(),
		GetCampaignByUserEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.CampaignService",
			"GetByUser",
			encodeGetCampaignUserRequest,
			decodeGetCampaignResponse,
			gamiProtobuf.GetCampaignResponse{},
		).Endpoint(),
		GetCampaignByCodeEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.CampaignService",
			"GetByCode",
			encodeGetCampaignUserRequest,
			decodeGetCampaignResponse,
			gamiProtobuf.GetCampaignResponse{},
		).Endpoint(),
		GetRewardEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RewardService", "GetByID",
			encodeGetRewardRequest,
			decodeGetRewardResponse,
			gamiProtobuf.GetRewardResponse{},
		).Endpoint(),
		GetRewardByCampaignIDEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RewardService", "GetRewardByCampaignID",
			encodeGetRewardPoolRequest,
			decodeGetRewardPoolResponse,
			gamiProtobuf.GetRewardPoolResponse{},
		).Endpoint(),
		GetListRewardByIDsEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RewardService",
			"GetList",
			encodeGetListRewardByIDsRequest,
			decodeGetListRewardResponse,
			gamiProtobuf.GetListRewardResponse{}).Endpoint(),
		RedeemEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RewardService", "Redeem",
			encodeRedeemRequest,
			decodeRedeemResponse,
			gamiProtobuf.RedeemResponse{},
		).Endpoint(),
		GetTransactionEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RewardService", "GetTransaction",
			encodeTransactionRequest,
			decodeTransactionResponse,
			gamiProtobuf.GetTransactionResponse{},
		).Endpoint(),
		GetRuleEndpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.RuleService", "GetByCampaignID",
			encodeGetRuleRequest,
			decodeGetRuleResponse,
			gamiProtobuf.GetRuleResponse{},
		).Endpoint(),
		ListCampaignsEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.CampaignService",
			"GetList",
			encodeListCampaignsRequest,
			decodeListCampaignsResponse,
			gamiProtobuf.GetListCampaignResponse{},
		).Endpoint(),
		GetListCampaignByIDsEndPoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.CampaignService",
			"GetListByIDs",
			encodeGetListCampaignByIDsRequest,
			decodeListCampaignsResponse,
			gamiProtobuf.GetListCampaignResponse{}).Endpoint(),
		AcceptChallengeEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.CampaignService",
			"AcceptChallenge",
			encodeAcceptChallengeRequest,
			decodeAcceptChallengeResponse,
			gamiProtobuf.AcceptChallengeResponse{}).Endpoint(),
		MaintenanceStatusCampaignEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.MaintenanceService",
			"MaintenanceStatusCampaign",
			encodeMaintenanceStatusCampaignRequest,
			decodeMaintenanceStatusCampaignResponse,
			gamiProtobuf.MaintenanceStatusCampaignResponse{}).Endpoint(),
		CampaignRewardAmountStatisticEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.CampaignService",
			"GetRewardAmountInCampaign",
			encodeRewardAmountStatisticRequest,
			decodeRewardAmountStatisticResponse,
			gamiProtobuf.RewardAmountStatisticResponse{}).Endpoint(),
		GetBlackWhiteListEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.BlackWhiteListService",
			"GetBlackWhiteList",
			encodeGetBlackWhiteListRequest,
			decodeGetBlackWhiteListResponse,
			gamiProtobuf.GetBlackWhiteListResponse{}).Endpoint(),
		GetRewardsByCampaignIDEndpoint: grpc_client.NewgRPCClient(
			connection,
			"gami_protobuf.RewardService",
			"GetRewardsByCampaignID",
			encodeGetRewardsByCampaignIDRequest,
			decodeGetRewardsByCampaignIDResponse,
			gamiProtobuf.GetRewardsByCampaignIDResponse{}).Endpoint(),
		GetCampaignByUserV2Endpoint: grpc_client.NewgRPCClient(
			connection, "gami_protobuf.CampaignService",
			"GetByUserV2",
			encodeGetCampaignUserRequest,
			decodeGetCampaignResponse,
			gamiProtobuf.GetCampaignResponse{},
		).Endpoint(),
	}
}
