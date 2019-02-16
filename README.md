# docker-demo
演示一个小的微服务架构如何用docker来编排

## 目录结构

* api1: 模拟后端api调用，go语言编写，端口 `8001`
* api1: 模拟另一个后端api调用，go语言编写，端口 `8002`
* main-site: 主网站，只有前端代码。在本例中，直接作为gateway

## 说明

1. 在 `main-site` 中，使用`nginx`反向代理了两个后端api。配置在 `nginx-default.conf` 中。这里要注意的是，反向代理时直接使用了域名`api1`与`api2`，这两个域名由 `docker-compose` 负责管理。因此在`docker-compose.yml`中，service的名字 `api1`与`api2`不可随便变更。另外注意`docker-compose.yml`中对`main-site`还配置了`links`。

2. 注意两个后端服务的端口都没有开放出来，在`docker-compose.yml`中只开放了网关(`main-site`)的内部端口。因此从外部只能看到一个开放的端口，而不是3个。这样虽然是多个服务，但并不会产生端口污染问题。实际上，api1、api2完全可以使用相同的80端口，只要修改相关配置信息即可。

## 使用

在安装了docker和docker-compose的机器上执行：

```
$ cd docker-demo
docker-demo$ docker-componse build     #构建docker image
docker-demo$ docker-compose up -d      #启动所有服务
#这时就可以访问 http://ip:8000 看效果
docker-demo$ docker-compose down       #停止所有服务
```
