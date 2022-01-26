# etcd-example

## Env

* mac amd64
* docker
* docker-compose
* etcd v3.4

## 搭建 etcd 服务
* 单节点
* 多节点

[点此查看搭建etcd服务文档](docs/build-etcd-service.md)

## example

- [x] Use etcdctl
  - put
  - get
  - watch

- [ ] Use go client
  - put
  - get
  - watch

- [x] service registration and discovery
  - 实现对配置文件的监控 viper + fsnotify
  - 实现对 etcd 的监控 watch 模式
- [x]
