package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jaronnie/mysrd/config"
	"github.com/jaronnie/mysrd/etcdsdk"
	"github.com/spf13/viper"
)

// Context is main application context
type Context struct {
	// ConfigReadMode is one of `local` or `etcd`
	ConfigReadMode string
	EtcdHost       string
	AppName        string

	// Magic context for hack
	Magic map[string]interface{}
}

func main() {
	ctx := Context{
		ConfigReadMode: "local",
		EtcdHost:       "127.0.0.1:2379",
		AppName:        "srd",
		Magic:          make(map[string]interface{}),
	}
	// new config manager
	manager := config.NewConfigManager(ctx.ConfigReadMode, ctx.EtcdHost, "srd")
	err := manager.NewConfig()
	if err != nil {
		panic(err)
	}

	preEtcdHost := viper.GetStringSlice("etcd.host")
	fmt.Printf("Currently etcd host is [%v]\n", preEtcdHost)

	dial, err := etcdsdk.NewEtcdDial(viper.GetStringSlice("etcd.host"), time.Duration(3)*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}

	// service register
	err = etcdsdk.RegisterService(context.Background(), dial, viper.GetString("service.prefix"), viper.GetString("service.name"), viper.GetString("service.grpc"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		tick := time.NewTicker(time.Second * 3)
		select {
		case <-tick.C:
			fmt.Println("Currently etcd host is", viper.GetStringSlice("etcd.host"))
			fmt.Println("Currently service host is", viper.GetStringSlice("service.grpc"))

		}
	}
}
