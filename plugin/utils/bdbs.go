package utils

import (
	"encoding/json"
	"fmt"
)

type bdbConf struct {
	Uid        int
	DBName     string
	Limit      int64
	ShardsUsed int
	Endpoints  int
}

var bdbsDat []RLdbd

func GetBDBs(conf *RLConf) (map[int]bdbConf, error) {
	d := make(map[int]bdbConf)
	u, httpCode, err := APIget(conf, "/v1/bdbs")
	if err != nil {
		return d, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return d, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &bdbsDat); err != nil {
		return d, err
	}

	for _, bdb := range bdbsDat {
		j := bdbConf{
			DBName:     bdb.Name,
			Uid:        bdb.Uid,
			Limit:      bdb.MemorySize,
			ShardsUsed: bdb.ShardsCount,
		}

		if bdb.Replication {
			j.ShardsUsed = j.ShardsUsed * 2 // replication
		}

		j.Endpoints = len(bdb.Endpoints[0].Addr)

		d[bdb.Uid] = j

	}

	return d, nil

}