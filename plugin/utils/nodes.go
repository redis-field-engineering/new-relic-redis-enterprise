package utils

import (
	"encoding/json"
	"fmt"
)

var nodeDat RLNodes
var capa RLNodeCapacity

func GetNodes(conf *RLConf) (RLNodeCapacity, error) {
	u, httpCode, err := APIget(conf, "/v1/nodes")
	if err != nil {
		return capa, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return capa, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &nodeDat); err != nil {
		return capa, err
	}

	for _, node := range nodeDat {
		capa.NodeMemory += node.TotalMemory
		capa.NodeCores += node.Cores
		if node.Status == "active" {
			capa.ActiveNodes++
		}
	}

	capa.NodeCount = len(nodeDat)

	return capa, nil

}
