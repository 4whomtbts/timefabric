package model

import "time"

type Instance struct {
	Id int32 `json:"id" db:"id"`
	NodeId int32  `json:"node_id" db:"node_id"`
	UserId int32 `json:"user_id" db:"user_id"`
	Name string `json:"name" db:"name"`
	InstanceHash string `json:"instance_hash" db:"instance_hash"`
	Image string `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Expiredat time.Time `json:"expired_at" db:"expired_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
