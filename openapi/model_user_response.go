package openapi

import Null "gopkg.in/guregu/null.v3"

// UserResponse - ユーザモデル
// ポインタは全部NULL.XXXXにしないとjsonでNullが返らない
type UserResponse struct {

	Id int32 `json:"id,omitempty"`

	Name Null.String `json:"name,omitempty"`

	Age int32 `json:"age,omitempty"`

	Weight Null.Int `json:"weight,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`
}
