package EtcdResolver

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"time"
)

func init() {
	resolver.Register(NewBuilder())
}

const (
	defaultFreq = time.Second * 30
)

func NewBuilder() resolver.Builder {
	return &EtcdResolverBuilder{freq: defaultFreq}
}

type EtcdResolver struct {
	etcdServerUrl    string
	serverListPath   string
	etcdClient       client.Client
	etcdKVAPI        client.KeysAPI
	addrsRefreshFreq time.Duration
	ctx              context.Context
	cancel           context.CancelFunc
	cc               resolver.ClientConn
	timer            *time.Timer
	rn               chan struct{}
}

func (r *EtcdResolver) ResolveNow(opt resolver.ResolveNowOption) {
	select {
	case r.rn <- struct{}{}:
	default:
	}
}

func (r *EtcdResolver) Close() {
	r.cancel()
	r.timer.Stop()
}

func (r *EtcdResolver) watcher() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-r.timer.C:
		case <-r.rn:
		}
		addrs, _ := r.lookup()
		r.timer.Reset(r.addrsRefreshFreq)
		r.cc.NewAddress(addrs)
		fmt.Println("NewAddress:", addrs)
	}
}

func (r *EtcdResolver) lookup() ([]resolver.Address, string) {
	var newAddrs []resolver.Address
	serverList, err := r.etcdKVAPI.Get(r.ctx, r.serverListPath, nil)
	if err != nil {
		return newAddrs, ""
	}
	for _, node := range serverList.Node.Nodes {
		addr := strings.Replace(node.Key, "/"+r.serverListPath+"/", "", 1)
		fmt.Println(addr)
		newAddrs = append(newAddrs, resolver.Address{Addr: addr})
	}
	return newAddrs, ""
}

type EtcdResolverBuilder struct {
	freq time.Duration
}

func (erb *EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	ip, port, path, ok := parseTarget(target.Endpoint)
	if !ok {
		return nil, fmt.Errorf("The Endpoint of etcd error:%s", target.Endpoint)
	}
	cfg := client.Config{
		Endpoints:               []string{"http://" + ip + ":" + port},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatalf("New CLient err:%v\n", err)
	}
	keyApi := client.NewKeysAPI(c)

	ctx, cancel := context.WithCancel(context.Background())

	etcdResolver := &EtcdResolver{
		etcdServerUrl:    ip + ":" + port,
		serverListPath:   path,
		etcdClient:       c,
		etcdKVAPI:        keyApi,
		addrsRefreshFreq: erb.freq,
		ctx:              ctx,
		cancel:           cancel,
		cc:               cc,
		timer:            time.NewTimer(0),
		rn:               make(chan struct{}, 1),
	}
	go etcdResolver.watcher()
	return etcdResolver, nil
}

func parseTarget(endpoint string) (string, string, string, bool) {
	fmt.Println("In parseTarget", endpoint)
	st := strings.SplitN(endpoint, ":", 3)
	if len(st) < 3 {
		return "", "", "", false
	}
	return st[0], st[1], st[2], true
}

func (erb *EtcdResolverBuilder) Scheme() string {
	return "EtcdResolver"
}
