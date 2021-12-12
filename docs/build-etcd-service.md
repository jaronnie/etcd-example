# 搭建 etcd 服务

image: bitnami/etcd:3.4.13-debian-10-r0

## 单节点

```shell
docker run -itd --name myetcd0 -p 23790:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.13-debian-10-r0
```

```shell
# 进入 etcd 容器
docker exec -it myetcd0 bash

# put jaronnie with gocloudcoder
etcdctl put jaronnie gocloudcoder

# get jaronnie
etcdctl get jaronnie

# watch jaronnie
etcdctl watch jaronnie

# put jaronnie new value, watch the value of jaronnie
etcdctl put jaronnie gocloudcoder1
```

## 多节点

```shell
cd etcd-example

# the path have docker-compose.yml
docker-compose up -d
```

### 检查健康情况

```shell
etcdctl member list
```

![image-20211212195617114](https://resource.gocloudcoder.com/image-20211212195617114.png)

### put/get

![image-20211212195500002](https://resource.gocloudcoder.com/image-20211212195500002.png)
