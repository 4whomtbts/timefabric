package model

import (
	"database/sql"
)

type TimeFabricNode struct {
	NodeId int32 `json:"node_id" db:"node_id"`
	StorageGroupId int32 `json:"storage_group_id" db:"storage_group_id"`
	HostName string `json:"host_name" db:"host_name"` // Literally, Linux hostname
	NodeName string `json:"node_name" db:"node_name"` // This is nickname for timefabric server which is named by admin
	HostIP string `json:"host_ip" db:"host_ip"`
	Port string `json:"port" db:"port"`
	Excluded bool `json:"excluded" db:"excluded"`
	Exclusive bool `json:"exclusive" db:"exclusive"`
	RegisteredAt sql.NullTime `json:"registered_at" db:"registered_at"`
	ExcludedAt sql.NullTime `json:"excluded_at" db:"excluded_at"`
	LastAllocatedAt sql.NullTime `json:"last_allocated_at" db:"last_allocated_at"`
	LastHeartbeatAt sql.NullTime `json:"last_heartbeat_at" db:"last_heartbeat_at"`
}