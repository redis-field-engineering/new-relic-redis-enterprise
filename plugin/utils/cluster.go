package utils

import (
	"encoding/json"
	"fmt"
)

type rlClusterConfig struct {
	Name string `json:"name"`
}

var cluster rlClusterConfig

func GetCluster(conf *RLConf) (rlClusterConfig, error) {
	u, httpCode, err := APIget(conf, "/v1/cluster", nil)
	if err != nil {
		return cluster, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return cluster, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &cluster); err != nil {
		return cluster, err
	}

	return cluster, nil

}
