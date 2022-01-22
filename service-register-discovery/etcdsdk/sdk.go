package etcdsdk

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func NewEtcdDial(endpoints []string, timeout time.Duration) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
}

func RegisterService(ctx context.Context, client *clientv3.Client, prefixName, serviceName string, value string) error {
	key := fmt.Sprintf("%s%s", prefixName, serviceName)
	fmt.Println("key", key)
	fmt.Println("Value", value)
	_, err := client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}
