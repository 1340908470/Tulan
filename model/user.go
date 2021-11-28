package model

type UserKey struct {
	UnionId string `json:"union_id"`
	UserId  string `json:"user_id"`
	OpenId  string `json:"open_id"`
}
