package common

type StorageMappingType int8

const (
	VOLUME_STORAGE = iota
	BIND_STORAGE
)

type StorageType string

const (
	LOCAL = "local"
	NAS = "nas"
)