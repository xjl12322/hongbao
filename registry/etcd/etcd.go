package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"hongbao/registry"
	"sync"
	"sync/atomic"
)

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

type RegisterService struct {
	id clientv3.LeaseID
	service *registry.Service
	registered bool
	keepALiveCh <-chan *clientv3.LeaseKeepAliveResponse



}

type EtcdRegistry struct {
	options *registry.Options
	client *clientv3.Client
	serviceCh chan *registry.Service
	value              atomic.Value
	lock               sync.Mutex
	registryServiceMap map[string]*RegisterService


}













