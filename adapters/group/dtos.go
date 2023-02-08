package group

const (
	UserGroupType = "USER"
	DataGroupType = "DATA"
)

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type IsUserInGroupsRequest struct {
	UserID   int64   `json:"user_id"`
	GroupIDs []int64 `json:"group_ids"`
}

type IsUserInGroupsResponse struct {
	Meta Meta               `json:"meta"`
	Data IsUserInGroupsData `json:"data"`
}

type IsUserInGroupsData struct {
	IsInGroups bool `json:"is_in_groups"`
}

type IsInGroupsRequest struct {
	Value     interface{} `json:"value"`
	GroupIDs  []int64     `json:"group_ids"`
	GroupType string      `json:"group_type"` // the type of the group, could use UserGroupType or DataGroupType
}

type IsInGroupsResponse struct {
	Meta Meta          `json:"meta"`
	Data IsInGroupData `json:"data"`
}

type IsInGroupData struct {
	IsInGroups bool `json:"is_in_groups"`
}

type ValuesInGroupsRequest struct {
	Data []*ValueInGroupsRequestData `json:"data"`
}

type ValueInGroupsRequestData struct {
	Value     interface{} `json:"value"`
	GroupIDs  []int64     `json:"group_ids"`
	GroupType string      `json:"group_type"`
}

type ValuesInGroupsResponse struct {
	Meta Meta                      `json:"meta"`
	Data map[string]map[int64]bool `json:"data"`
}

type ValuesInGroupsData struct {
	Data map[int64]bool `json:"data"`
}
