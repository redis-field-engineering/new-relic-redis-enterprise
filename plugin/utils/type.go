package utils

import (
	"time"
)

//RLConf the configuration necesseary to get information fromt the RL API
type RLConf struct {
	Hostname string
	Port     int
	User     string
	Pass     string
}

type RLlicenseConfig struct {
	ActivationDate      time.Time `json:"activation_date"`
	ExpirationDate      time.Time `json:"expiration_date"`
	Expired             bool      `json:"expired"`
	Features            []string  `json:"features"`
	ShardsLimit         int       `json:"shards_limit"`
	DaysUntilExpiration int
}

type RLNodes []struct {
	AcceptServers             bool     `json:"accept_servers"`
	Addr                      string   `json:"addr"`
	Architecture              string   `json:"architecture"`
	BigredisStoragePath       string   `json:"bigredis_storage_path"`
	BigstoreDriver            string   `json:"bigstore_driver,omitempty"`
	Cores                     int      `json:"cores"`
	EphemeralStoragePath      string   `json:"ephemeral_storage_path"`
	EphemeralStorageSize      float64  `json:"ephemeral_storage_size"`
	ExternalAddr              []string `json:"external_addr"`
	MaxListeners              int      `json:"max_listeners"`
	MaxRedisServers           int      `json:"max_redis_servers"`
	OsName                    string   `json:"os_name"`
	OsSemanticVersion         string   `json:"os_semantic_version"`
	OsVersion                 string   `json:"os_version"`
	PersistentStoragePath     string   `json:"persistent_storage_path"`
	PersistentStorageSize     float64  `json:"persistent_storage_size"`
	RackID                    string   `json:"rack_id"`
	ShardCount                int      `json:"shard_count"`
	ShardList                 []int    `json:"shard_list"`
	SoftwareVersion           string   `json:"software_version"`
	Status                    string   `json:"status"`
	SupportedDatabaseVersions []struct {
		DbType  string `json:"db_type"`
		Version string `json:"version"`
	} `json:"supported_database_versions"`
	TotalMemory int64 `json:"total_memory"`
	UID         int   `json:"uid"`
	Uptime      int   `json:"uptime"`
}

type RLNodeCapacity struct {
	NodeCores   int
	NodeMemory  int64
	NodeCount   int
	ActiveNodes int
}
