package group

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	proto "gitlab.id.vin/gami/gami-proto/group"
	"gitlab.id.vin/gami/gami-proto/grpc_client"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

type Adapter interface {
	IsInGroups(ctx context.Context, req *IsInGroupsRequest) (*IsInGroupsResponse, error)
	IsUserInGroups(ctx context.Context, req *IsUserInGroupsRequest) (*IsUserInGroupsResponse, error)
	ValuesInGroups(ctx context.Context, req *ValuesInGroupsRequest) (*ValuesInGroupsResponse, error)
}

type groupAdapter struct {
	ctx                    context.Context
	IsUserInGroupsEndpoint endpoint.Endpoint
	IsInGroupsEndpoint     endpoint.Endpoint
	ValuesInGroupsEndpoint endpoint.Endpoint
}

func (h *groupAdapter) IsInGroups(ctx context.Context, req *IsInGroupsRequest) (*IsInGroupsResponse, error) {
	res, err := h.IsInGroupsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	r, ok := res.(*IsInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	return r, nil
}

func encodeIsInGroupsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*IsInGroupsRequest)
	if !ok {
		return nil, errors.New("invalid data type")
	}

	v, err := structpb.NewValue(r.Value)
	if err != nil {
		return nil, err
	}

	any, err := anypb.New(v)
	if err != nil {
		return nil, err
	}

	return &proto.IsInGroupsRequest{
		Element:   any,
		GroupIDs:  r.GroupIDs,
		GroupType: r.GroupType,
	}, nil
}

func decodeIsInGroupsResponse(_ context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(*proto.IsInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	if r.Meta.Status != http.StatusOK {
		return &IsInGroupsResponse{
			Meta: Meta{
				Status:  int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	return &IsInGroupsResponse{
		Meta: Meta{
			Status:  int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: IsInGroupData{
			IsInGroups: r.Data.IsInGroups,
		},
	}, nil
}

func (h *groupAdapter) IsUserInGroups(ctx context.Context, req *IsUserInGroupsRequest) (*IsUserInGroupsResponse, error) {
	res, err := h.IsUserInGroupsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	r, ok := res.(*IsUserInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	return r, nil
}

func encodeIsUserInGroupsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*IsUserInGroupsRequest)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	return &proto.UserInGroupsRequest{
		UserID:   r.UserID,
		GroupIDs: r.GroupIDs,
	}, nil
}

func decodeIsUserInGroupsResponse(_ context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(*proto.IsInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	if r.Meta.Status != http.StatusOK {
		return &IsUserInGroupsResponse{
			Meta: Meta{
				Status:  int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	return &IsUserInGroupsResponse{
		Meta: Meta{
			Status:  int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: IsUserInGroupsData{
			IsInGroups: r.Data.IsInGroups,
		},
	}, nil
}

func (h *groupAdapter) ValuesInGroups(ctx context.Context, req *ValuesInGroupsRequest) (*ValuesInGroupsResponse, error) {
	res, err := h.ValuesInGroupsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	r, ok := res.(*ValuesInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	return r, nil
}
func encodeValuesInGroupsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*ValuesInGroupsRequest)
	if !ok {
		return nil, errors.New("invalid data type")
	}

	data := make([]*proto.ValuesInGroupsRequestData, 0)

	for _, v := range r.Data {
		value, err := structpb.NewValue(v.Value)
		if err != nil {
			return nil, err
		}
		any, err := anypb.New(value)
		if err != nil {
			return nil, err
		}

		data = append(data, &proto.ValuesInGroupsRequestData{
			Value:     any,
			GroupIDs:  v.GroupIDs,
			GroupType: v.GroupType,
		})
	}

	return &proto.ValuesInGroupsRequest{
		Data: data,
	}, nil
}

func decodeValuesInGroupsResponse(_ context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(*proto.ValuesInGroupsResponse)
	if !ok {
		return nil, errors.New("invalid data type")
	}
	if r.Meta.Status != http.StatusOK {
		return &ValuesInGroupsResponse{
			Meta: Meta{
				Status:  int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	data := make(map[string]map[int64]bool, len(r.Data))
	for k, v := range r.Data {
		data[k] = v.Data
	}
	return &ValuesInGroupsResponse{
		Meta: Meta{
			Status:  int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: data,
	}, nil
}
func NewAdapter(ctx context.Context, connection *grpc.ClientConn) Adapter {
	return &groupAdapter{
		ctx: ctx,
		IsUserInGroupsEndpoint: grpc_client.NewgRPCClient(
			connection, "group.GroupService", "UserInGroups",
			encodeIsUserInGroupsRequest, decodeIsUserInGroupsResponse, proto.IsInGroupsResponse{},
		).Endpoint(),
		IsInGroupsEndpoint: grpc_client.NewgRPCClient(
			connection, "group.GroupService", "IsInGroups",
			encodeIsInGroupsRequest, decodeIsInGroupsResponse, proto.IsInGroupsResponse{},
		).Endpoint(),
		ValuesInGroupsEndpoint: grpc_client.NewgRPCClient(
			connection, "group.GroupService", "ValuesInGroups",
			encodeValuesInGroupsRequest, decodeValuesInGroupsResponse, proto.ValuesInGroupsResponse{},
		).Endpoint(),
	}
}
