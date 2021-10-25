package main

import (
	"fmt"
	"time"

	"github.com/Redis-Field-Engineering/newrelic-redis-enterprise/plugin/utils"
	sdkArgs "github.com/newrelic/infra-integrations-sdk/v4/args"
	"github.com/newrelic/infra-integrations-sdk/v4/integration"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname string `default:"localhost" help:"Hostname or IP where Redis server is running."`
	Port     int    `default:"9443" help:"Port on which Redis server is listening."`
	Username string `default:"admin@example.com" help:"Username to login as."`
	Password string `default:"myPass" help:"Password for login."`
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

	// Fetch the list of Redis databases
	conf := &utils.RLConf{
		Hostname: args.Hostname,
		Port:     args.Port,
		User:     args.Username,
		Pass:     args.Password,
	}

	// Get the license information
	license, err := utils.GetLicense(conf)
	panicOnErr(err)

	// Get node information
	nodes, err := utils.GetNodes(conf)
	panicOnErr(err)
	fmt.Printf("+%v\n", nodes)

	// Create Entity, entities name must be unique
	e1, err := i.NewEntity(args.Hostname, "RedisEnterpriseCluster", args.Hostname)
	panicOnErr(err)
	// Add event when redis starts
	if args.All() || args.Events {
		fmt.Println("Events go here")
		//ev, _ := event.NewNotification("Redis Server recently started")
		//e1.AddEvent(ev)
	}

	// Add Inventory item
	if args.All() || args.Inventory {
		fmt.Print("inventory goes here")
		//key, value := queryAttrRedisInfo("redis_version", instanceOnePort)
		//if key != "" {
		//	err = e1.AddInventoryItem(key, "value", value)
		//}
		//panicOnErr(err)
	}

	// Add Metric
	if args.All() || args.Metrics {
		fmt.Println("Metrics go here")
		g, _ := integration.Gauge(time.Now(), "license.DaysUntilExpiration", float64(license.DaysUntilExpiration))
		e1.AddMetric(g)
		f, _ := integration.Gauge(time.Now(), "license.ShardsLicense", float64(license.ShardsLimit))
		e1.AddMetric(f)
		h, _ := integration.Gauge(time.Now(), "license.ClusterTotalMemory", float64(nodes.NodeMemory))
		e1.AddMetric(h)
		i, _ := integration.Gauge(time.Now(), "license.ClusterTotalCores", float64(nodes.NodeCores))
		e1.AddMetric(i)
		j, _ := integration.Gauge(time.Now(), "license.ClusterActiveNodes", float64(nodes.ActiveNodes))
		e1.AddMetric(j)
		k, _ := integration.Gauge(time.Now(), "license.ClusterNodes", float64(nodes.NodeCount))
		e1.AddMetric(k)
	}

	// Add the first entity to the integration
	i.AddEntity(e1)

	// Print the JSON document to stdout
	panicOnErr(i.Publish())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
