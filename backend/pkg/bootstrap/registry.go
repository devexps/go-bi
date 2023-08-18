package bootstrap

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"

	// consul
	consulMicro "github.com/devexps/go-micro/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// etcd
	etcdMicro "github.com/devexps/go-micro/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"
)

type RegistryType string

const (
	RegistryTypeConsul RegistryType = "consul"
	RegistryTypeEtcd   RegistryType = "etcd"
)

// NewRegistry creates new registry
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}
	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	}
	return nil
}

// NewDiscovery creates a discovery client
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	}
	return nil
}

// NewConsulRegistry creates a registry discovery client - Consul
func NewConsulRegistry(c *conf.Registry) *consulMicro.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}
	return consulMicro.New(cli, consulMicro.WithHealthCheck(c.Consul.HealthCheck))
}

// NewEtcdRegistry creates a registry discovery client - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdMicro.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		log.Fatal(err)
	}
	return etcdMicro.New(cli)
}
