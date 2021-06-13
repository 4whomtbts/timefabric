package model

import "time"

type Gpu struct {
	Id int32 `json:"id" db:"id"`
	NodeId uint32  `json:"node_id" db:"node_id"`
	ModelName string `json:"model_name" db:"model_name"`
	BusId string `json:"bus_id" db:"bus_id"`
	SlotIndex uint32 `json:"slot_index" db:"slot_index"`
	CurrAllocCnt uint32 `json:"curr_alloc_cnt" db:"curr_alloc_cnt"`
	Exclusive bool `json:"exclusive" db:"exclusive"`
	ExcludedAt time.Time `json:"excluded_at" db:"excluded_at"`
	LastAllocatedAt time.Time `json:"last_allocated_at" db:"last_allocated_at"`
	//MaxMemory uint32 `json:"max_memory" db:"max_memory"`
}
