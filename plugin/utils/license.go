package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

var dat RLlicenseConfig

func GetLicense(conf *RLConf) (RLlicenseConfig, error) {
	u, httpCode, err := APIget(conf, "/v1/license")
	if err != nil {
		return dat, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return dat, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &dat); err != nil {
		return dat, err
	}

	t1 := time.Now()
	diff := dat.ExpirationDate.Sub(t1)
	dat.DaysUntilExpiration = int(diff.Hours() / 24)

	return dat, nil

}
