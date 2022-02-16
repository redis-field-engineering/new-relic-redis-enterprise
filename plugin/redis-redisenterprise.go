package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Redis-Field-Engineering/newrelic-redis-enterprise/plugin/utils"
	"github.com/leekchan/timeutil"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/data/event"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname  string `default:"localhost" help:"Hostname or IP where Redis server is running."`
	Port      int    `default:"9443" help:"Port on which Redis server is listening."`
	Username  string `default:"admin@example.com" help:"Username to login as."`
	Password  string `default:"myPass" help:"Password for login."`
	Eventtime int    `default:"60" help:"Interval for scraping events."`
}

const (
	integrationName    = "com.redis.redisenterprise"
	integrationVersion = "0.1.0"
)

var (
	args argumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	panicOnErr(err)

	bdbEnts := make(map[int]*integration.Entity)

	// Fetch the list of Redis databases
	conf := &utils.RLConf{
		Hostname: args.Hostname,
		Port:     args.Port,
		User:     args.Username,
		Pass:     args.Password,
	}

	// Check to see if we are connecting to the cluster leader if not just exit
	// This is useful when running the integration all cluster nodes as only the master will submit metrics/events

	redir, err := utils.APIisRedirect(conf, "/v1/cluster")
	if err != nil || redir {
		return
	}

	// Get the license information
	license, err := utils.GetLicense(conf)
	panicOnErr(err)

	// Get node information
	nodes, err := utils.GetNodes(conf)
	panicOnErr(err)

	// Get the list of Redis databases
	bdbs, err := utils.GetBDBs(conf)
	panicOnErr(err)

	// Grab the Redis DB stats
	bdbStats, err := utils.GetBDBStats(conf)
	panicOnErr(err)

	// Create Entity, entities name must be unique
	e1, err := i.Entity(args.Hostname, "custom")
	panicOnErr(err)

	for _, val := range bdbs {
		s := fmt.Sprintf("%s:%s", args.Hostname, val.DBName)
		bdbEnts[val.Uid], err = i.Entity(s, "custom")
		panicOnErr(err)
	}

	// Add event when redis starts
	if args.All() || args.Events {
		st := time.Now().Add(time.Duration(-args.Eventtime) * time.Second)
		params := map[string]string{"stime": timeutil.Strftime(&st, "%Y-%m-%dT%H:%M:%SZ")}
		events, err := utils.GetEvents(conf, params)
		panicOnErr(err)
		for _, evnt := range events {
			tags := make(map[string]interface{})
			outstring := ""
			if evnt.Description != "" {
				outstring = evnt.Description
			} else {
				outstring = evnt.Type
			}
			p := []string{"OriginatorUsername", "OriginatorEmail", "NodeUID", "ModuleName", "BdbName", "UserUID", "Type", "Description"}
			for _, x := range p {
				s := reflect.ValueOf(evnt).FieldByName(x).Interface().(string)
				if s != "" {
					tags[x] = s
				}
			}
			ev := event.NewNotification(outstring)
			ev.Attributes = tags
			/*
			 TODO: Add event time somehow
			 ev.Timestamp = evnt.Time.Unix() doesn't work in v3
			 maybe persist and just get the current run time so we don't duplicate events
			 https://github.com/newrelic/infra-integrations-sdk/blob/master/docs/toolset/persist.md
			*/
			err = e1.AddEvent(ev)
			panicOnErr(err)

		}
	}

	// Add Inventory
	if args.All() || args.Inventory {
		err := e1.SetInventoryItem("RedisEnterpriseType", "value", "cluster")
		panicOnErr(err)
		for _, x := range bdbEnts {
			err := x.SetInventoryItem("RedisEnterpriseType", "value", "database")
			panicOnErr(err)
		}
	}
	// Add Metric
	if args.All() || args.Metrics {
		ms := e1.NewMetricSet("RedisEnterprise")
		err = ms.SetMetric("cluster.DaysUntilExpiration", float64(license.DaysUntilExpiration), metric.GAUGE)
		panicOnErr(err)
		err = ms.SetMetric("cluster.ShardsLicense", float64(license.ShardsLimit), metric.GAUGE)
		panicOnErr(err)
		err = ms.SetMetric("cluster.ClusterTotalMemory", float64(nodes.NodeMemory), metric.GAUGE)
		panicOnErr(err)
		err = ms.SetMetric("cluster.ClusterTotalCores", float64(nodes.NodeCores), metric.GAUGE)
		panicOnErr(err)
		err = ms.SetMetric("cluster.ClusterActiveNodes", float64(nodes.ActiveNodes), metric.GAUGE)
		panicOnErr(err)
		err = ms.SetMetric("cluster.ClusterNodes", float64(nodes.NodeCount), metric.GAUGE)
		panicOnErr(err)
		for _, val := range bdbs {
			// Setup the list of metrics we want to submit
			bdb_stats_list := []string{
				"AvgLatency", "AvgReadLatency", "AvgWriteLatency", "Conns", "EgressBytes", "EvictedObjects",
				"ExpiredObjects", "IngressBytes", "OtherReq", "ReadHits", "ReadMisses", "ReadReq",
				"ShardCPUSystem", "ShardCPUUser", "TotalReq", "UsedMemory", "WriteHits", "WriteMisses", "WriteReq"}

			// If the DB has RoF enabled, add the RoF metrics otherwise skip
			if val.Bigstore {
				bdb_stats_list = append(bdb_stats_list, []string{
					"BigstoreObjsRam", "BigstoreObjsFlash", "BigstoreIoReads",
					"BigstoreIoWrites", "BigstoreThroughput", "BigWriteRam",
					"BigWriteFlash", "BigDelRam", "BigDelFlash"}...)
			}

			bdbMs := bdbEnts[val.Uid].NewMetricSet("RedisEnterprise")
			err = bdbMs.SetMetric("bdb.ShardCount", float64(val.ShardsUsed), metric.GAUGE)
			panicOnErr(err)
			err = bdbMs.SetMetric("bdb.Endpoints", float64(val.Endpoints), metric.GAUGE)
			panicOnErr(err)
			err = bdbMs.SetMetric("bdb.MemoryLimit", float64(val.Limit), metric.GAUGE)
			panicOnErr(err)
			//// Grab all bdb Gauges
			for _, x := range bdb_stats_list {
				s := reflect.ValueOf(bdbStats[val.Uid]).FieldByName(x).Interface().(float64)
				err = bdbMs.SetMetric(fmt.Sprintf("bdb.%s", x), float64(s), metric.GAUGE)
				panicOnErr(err)
			}
		}
	}

	panicOnErr(i.Publish())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
