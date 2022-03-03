# NewRelic Redis Enterprise Plugin Development Environment

## Installing the plugin

#### Ensure New Relic agent is installed and running

[Agent installation instructions](https://docs.newrelic.com/docs/infrastructure/install-infrastructure-agent/get-started/install-infrastructure-agent/)


Check the service status

```
sudo service newrelic-infra status -l

# OR depending on your Linux distro

sudo /bin/systemctl status  -l newrelic-infra.service
```


#### Pull the rlease from Github Releases

https://github.com/redis-field-engineering/newrelic-redis-enterprise/releases


#### Unarchive

```
sudo su -
mkdir -p /tmp/nr_install
cd /tmp/nr_install
wget $RELEASE_DOWNLOAD
tar zxvf *.tar.gz 
mkdir -p /var/db/newrelic-infra/custom-integrations/bin
cp newrelic-redis-enterprise /var/db/newrelic-infra/custom-integrations/bin
cp conf/redis-redisenterprise-definition.yml /var/db/newrelic-infra/custom-integrations/
cp conf/redis-redisenterprise-multi-config.yml.example conf/redis-redisenterprise-multi-config.yml
vi conf/redis-redisenterprise-multi-config.yml
mv conf/redis-redisenterprise-multi-config.yml /etc/newrelic-infra/integrations.d/redis-redisenterprise-config.yaml
```

#### Restart New relic

```
sudo service  newrelic-infra  restart
```
