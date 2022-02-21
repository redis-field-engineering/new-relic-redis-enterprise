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

// RLdbs is the structure of the databases
type RLdbd struct {
	Uid       int `json:"uid"`
	Endpoints []struct {
		Addr []string `json:"addr"`
	} `json:"endpoints"`
	MemorySize  int64  `json:"memory_size"`
	Name        string `json:"name"`
	Replication bool   `json:"replication"`
	ShardsCount int    `json:"shards_count"`
	Bigstore    bool   `json:"bigstore"`
	Crdt        bool   `json:"crdt"`
	SyncSources []struct {
		Status string `json:"status"`
	} `json:"sync_sources"`
}

type RLdbStats map[int]RLdbStat

type RLdbStat struct {
	AvgLatency         float64 `json:"avg_latency"`
	AvgReadLatency     float64 `json:"avg_read_latency"`
	AvgWriteLatency    float64 `json:"avg_write_latency"`
	Conns              float64 `json:"conns"`
	EgressBytes        float64 `json:"egress_bytes"`
	EvictedObjects     float64 `json:"evicted_objects"`
	ExpiredObjects     float64 `json:"expired_objects"`
	IngressBytes       float64 `json:"ingress_bytes"`
	OtherReq           float64 `json:"other_req"`
	ReadHits           float64 `json:"read_hits"`
	ReadMisses         float64 `json:"read_misses"`
	ReadReq            float64 `json:"read_req"`
	ShardCPUSystem     float64 `json:"shard_cpu_system"`
	ShardCPUUser       float64 `json:"shard_cpu_user"`
	TotalReq           float64 `json:"total_req"`
	UsedMemory         float64 `json:"used_memory"`
	WriteHits          float64 `json:"write_hits"`
	WriteMisses        float64 `json:"write_misses"`
	WriteReq           float64 `json:"write_req"`
	BigstoreObjsRam    float64 `json:"bigstore_objs_ram"`
	BigstoreObjsFlash  float64 `json:"bigstore_objs_flash"`
	BigstoreIoReads    float64 `json:"bigstore_io_reads"`
	BigstoreIoWrites   float64 `json:"bigstore_io_writes"`
	BigstoreThroughput float64 `json:"bigstore_throughput"`
	BigWriteRam        float64 `json:"big_write_ram"`
	BigWriteFlash      float64 `json:"big_write_flash"`
	BigDelRam          float64 `json:"big_del_ram"`
	BigDelFlash        float64 `json:"big_del_flash"`
}

type RLEvents []RLEvent
type RLEvent struct {
	ModuleName         string    `json:"module_name,omitempty"`
	ModuleUID          string    `json:"module_uid,omitempty"`
	OriginatorEmail    string    `json:"originator_email,omitempty"`
	OriginatorUID      string    `json:"originator_uid"`
	OriginatorUsername string    `json:"originator_username"`
	Severity           string    `json:"severity"`
	Time               time.Time `json:"time"`
	Type               string    `json:"type"`
	UserUID            string    `json:"user_uid,omitempty"`
	NodeCount          int       `json:"node_count,omitempty"`
	State              bool      `json:"state,omitempty"`
	Username           string    `json:"username,omitempty"`
	NodeUID            string    `json:"node_uid,omitempty"`
	BdbName            string    `json:"bdb_name,omitempty"`
	BdbUID             int       `json:"bdb_uid,omitempty"`
	CheckName          string    `json:"check_name,omitempty"`
	ChecksError        string    `json:"checks_error,omitempty"`
	ChecksTime         int       `json:"checks_time,omitempty"`
	Description        string    `json:"description,omitempty"`
	UserName           string    `json:"user_name,omitempty"`
}

type CrdtStats struct {
	PeerStats []struct {
		Intervals []struct {
			CrdtEgressBytes             float64 `json:"egress_bytes"`
			CrdtEgressBytesDecompressed float64 `json:"egress_bytes_decompressed"`
			CrdtPendingLocalWritesMax   float64 `json:"pending_local_writes_max"`
			CrdtPendingLocalWritesMin   float64 `json:"pending_local_writes_min"`
		} `json:"intervals"`
	} `json:"peer_stats"`
}
