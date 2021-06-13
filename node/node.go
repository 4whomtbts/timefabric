package node

import "github.com/4whomtbts/timefabric/config"

type Node struct  {
	config config.TimeFabricConfig
}

// This contains every useful system info of node that helps user to decide which server the user gonna use
type NodeSystemInfo struct {

}

func NewNode(config config.TimeFabricConfig) *Node {
	return &Node{
		config,
	}
}

func mockRegisterApiCall() {}

func (n *Node) tryRegister() {
	// response = registerApi(config.MasterServer/register)
	/*
		if err {
			retry..
		}

		if alreadyRegistered {

		} else {
			log
		}
	 */
}


