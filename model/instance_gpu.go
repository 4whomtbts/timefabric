package model

import "time"

type InstanceGpu struct {
	Id int32 `json:"id" db:"id"`
	InstanceId int32 `json:"instance_id" db:"instance_id"`
	GpuId int32 `json:"gpu_id" db:"gpu_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
