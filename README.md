# etcd-example

## Env

* mac amd64
* docker
* etcd v3.4

docker 启动 etcd

```shell
docker run -itd --name myetcd -p 23790:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.13-debian-10-r0
```

## example

- [ ] Use etcdctl
  - put
  - get
  - watch

- [ ] service registration and discovery
  - 实现对配置文件的监控 viper + fsnotify
  - 实现对 etcd 的监控 watch 模式
