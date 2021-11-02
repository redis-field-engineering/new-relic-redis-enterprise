package utils

import (
	"encoding/json"
	"fmt"
)

func GetEvents(conf *RLConf) (RLEvents, error) {
	dat := RLEvents{}
	u, httpCode, err := APIget(conf, "/v1/logs", nil)
	if err != nil {
		return dat, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return dat, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &dat); err != nil {
		return dat, err
	}

	return dat, nil

}
