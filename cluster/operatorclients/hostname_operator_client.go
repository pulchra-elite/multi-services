package operatorclients

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/tendermint/tendermint/libs/log"
	"k8s.io/client-go/rest"

	clusterutil "github.com/ovrclk/provider-services/cluster/util"
)

const (
	hostnameOperatorHealthPath = "health"
)

type HostnameOperatorClient interface {
	Check(ctx context.Context) error
	String() string

	Stop()
}

type hostnameOperatorClient struct {
	sda    clusterutil.ServiceDiscoveryAgent
	client clusterutil.ServiceClient
	log    log.Logger
	l      sync.Locker
}

func NewHostnameOperatorClient(logger log.Logger, kubeConfig *rest.Config, endpoint *net.SRV) (HostnameOperatorClient, error) {
	sda, err := clusterutil.NewServiceDiscoveryAgent(logger, kubeConfig, "status", "akash-hostname-operator", "akash-services", endpoint)
	if err != nil {
		return nil, err
	}

	return &hostnameOperatorClient{
		log: logger.With("operator", "hostname"),
		sda: sda,
		l:   &sync.Mutex{},
	}, nil

}

func (hopc *hostnameOperatorClient) newRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	hopc.l.Lock()
	defer hopc.l.Unlock()

	if nil == hopc.client {
		var err error
		hopc.client, err = hopc.sda.GetClient(ctx, false, false)
		if err != nil {
			return nil, err
		}
	}

	return hopc.client.CreateRequest(ctx, method, path, body)
}

func (hopc *hostnameOperatorClient) Check(ctx context.Context) error {
	req, err := hopc.newRequest(ctx, http.MethodGet, hostnameOperatorHealthPath, nil)
	if err != nil {
		return err
	}

	response, err := hopc.client.DoRequest(req)
	if err != nil {
		return err
	}
	hopc.log.Info("check result", "status", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		return errNotAlive
	}

	return nil
}

func (hopc *hostnameOperatorClient) String() string {
	return fmt.Sprintf("<%T %p>", hopc, hopc)
}
func (hopc *hostnameOperatorClient) Stop() {
	hopc.sda.Stop()
}
