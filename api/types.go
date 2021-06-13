package api

import "github.com/4whomtbts/timefabric/model"

type RequiredResource struct {
	NumOfGpu int32


}

type NewNode struct {
	node model.TimeFabricNode
	gpu []model.Gpu
}