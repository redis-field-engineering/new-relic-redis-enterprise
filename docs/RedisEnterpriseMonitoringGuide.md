# Redis Enterprise Monitoring Guide 

## Key Performance Indicators

### Average Latency

#### Query

```
SELECT average(`bdb.AvgLatency`) FROM RedisEnterprise FACET entityName SINCE 30 MINUTES AGO TIMESERIES
```

#### Description

Average latency of operations on the database from the time the request hits the Enterprise proxy until the request is returned. Measured in seconds

#### Possible Causes

Cause | Check
--- | ---
Possible spike in activity | Check the operations per second
Slow running queries | Check the slow log in the Redis Enterprise UI for the database

#### Remediation Steps

Action | Method
--- | ---
Increase the number of database shards to handle the load | The database can be scaled up online by going to the Web UI and enabling clustering on the database
Work with application team to write more performant queries | Developers can fetch the slow log with the cli.  redis-cli -h HOST -p PORT -a PASSWORD SLOWLOG GET 100

#### Notes

Redis Enterprise generally returns results in the single digit milliseconds when tuned properly.
Read, write and other (administrative commands) latencies are also available


### Operations Per Second

#### Query

```
SELECT average(`bdb.TotalReq`) FROM RedisEnterprise FACET entityName SINCE 30 MINUTES AGO TIMESERIES
```

#### Description

Total number of Redis operations per second

#### Possible Causes

Cause | Check
--- | ---
Increased Activity | Check application code changes

#### Remediation Steps

Action | Method
--- | ---
Enable Pipelining | https://redis.io/docs/manual/pipelining/ while this does not decrease the operation count it will make them more efficient
Scale Up | Redis Enterprise allows you to scale up database size with no downtime

#### Notes

A conservative rule for Redis operations is about 25K Ops/sec with Redis (5K Ops/Sec with Redis On Flash) per master shard in the database.  This varies greatly depending on the use case, but this is a safe starting point.


### Database Capacity

#### Query

```
SELECT average(`bdb.UsedMemoryPercent`) FROM RedisEnterprise FACET entityName SINCE 30 MINUTES AGO TIMESERIES
```

#### Description

Amount of used memory divided by the memory limit as a percentage.

#### Possible Causes

Cause | Check
--- | ---
Possible spike in activity | Check the operations per second graphs (especially write)
Database sized incorrectly  | View usage over time


#### Remediation Steps

Action | Method 
--- | ---
Increase the database size | In the Redis Enterprise Web UI, click configuration, then edit, then modify Memory Limit setting and save

#### Notes

80% usage is the default alerting rule that is recommended.
If there is a small hot subset of data [Redis On Flash](https://redis.com/redis-enterprise/technology/redis-on-flash/) may be a more cost effective solution.


### Cache Hit Rate

Only useful when Redis is used as a cache

#### Query

```
SELECT average(`bdb.ReadHits`) FROM RedisEnterprise FACET entityName SINCE 30 MINUTES AGO TIMESERIES
SELECT average(`bdb.WriteHits`) FROM RedisEnterprise FACET entityName SINCE 30 MINUTES AGO TIMESERIES

```

#### Description

The number of reads or writes where they key already exists

#### Notes

This is *extremely* application specific and it is recommended to look carefully at changes over time.

## Secondary Indicators

Metric | Description | Notes
--- | --- | ---
bdb.IngressBytes|Incoming Network Traffic|Helpful for spotting changes in behavior
bdb.EgressBytes|Outgoing Network Traffic|Helpful for spotting changes in behavior
bdb.Conns|Number of connections to the database|Recommend monitoring both a maximum and a minimum
bdb.ExpiredObjects|Items that have TTL'd out of the cache|This is normal and indicates expiration is working properly
bdb.EvictedObjects|Items that have been forced out of the cache|This is when the database is memory starved and should be investigated. [Evicition Policies](https://docs.redis.com/latest/rs/databases/configure/eviction-policy/)