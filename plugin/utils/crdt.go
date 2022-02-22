package utils

import (
	"encoding/json"
	"fmt"
)

type CrdtPeerStats struct {
	PeerStats []PeerStats `json:"peer_stats"`
}
type Intervals struct {
	CrdtEgressBytes              float64 `json:"egress_bytes"`
	CrdtEgressBytesDecompressed  float64 `json:"egress_bytes_decompressed"`
	CrdtPendingLocalWritesMax    float64 `json:"pending_local_writes_max"`
	CrdtPendingLocalWritesMin    float64 `json:"pending_local_writes_min"`
	CrdtIngressBytes             float64 `json:"ingress_bytes,omitempty"`
	CrdtIngressBytesDecompressed float64 `json:"ingress_bytes_decompressed,omitempty"`
	CrdtLocalIngressLagTime      float64 `json:"local_ingress_lag_time,omitempty"`
}
type PeerStats struct {
	Intervals []Intervals `json:"intervals"`
	UID       string      `json:"uid"`
}

func GetCrdt(conf *RLConf, uid int, params map[string]string) (Intervals, error) {
	dat := CrdtPeerStats{}
	results := Intervals{}
	u, httpCode, err := APIget(
		conf,
		fmt.Sprintf("/v1/bdbs/%d/peer_stats", uid),
		params,
	)
	if err != nil {
		return results, fmt.Errorf("unable to connect: %s", err)
	}
	if httpCode != 200 {
		return results, fmt.Errorf("HTTP Status code is wrong:%d - should be 200", httpCode)
	}

	if err := json.Unmarshal([]byte(u), &dat); err != nil {
		return results, err
	}

	if len(dat.PeerStats) > 0 {
		results = dat.PeerStats[0].Intervals[len(dat.PeerStats[0].Intervals)-1]
	}

	return results, nil

}
