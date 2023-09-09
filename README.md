# redis集群
## 平台 
Windows + Docker
## 集群搭建步骤
1. 运行命令`docker-compose up -d`创建容器
2. 运行`docker exec -it redis1 bash`进入容器redis1
3. 在redis1容器环境中运行`redis-cli --cluster create 172.16.10.1:7001 172.16.10.2:7002 172.16.10.3:7003 172.16.10.4:7004 172.16.10.5:7005 172.16.10.6:7006 --cluster-replicas 1`

以上过程就完成了redis集群环境的搭建，创建了一个3主3从的redis集群。

## 数据读写
1. 写入数据：`redis-cli -c -p 7001 set name leebai`
2. 读取数据：`redis-cli -c -p 7001 get name`
> 以上2条命令的执行都需要在redis1的容器环境下执行


## 测试
`redis_test.go`中两个函数。
* TestRedisNodeGet：以节点方式连接
* TestRedisClusterGet：以集群方式连接

### 出现的问题
* 方式一能正常访问docker环境内的容器，并且能够拿到正常的数据
* 方式二不能， 好像是连访问都访问不了。显示错误如下
    ```shell
    panic: dial tcp 172.16.10.2:7002: i/o timeout [recovered]
        panic: dial tcp 172.16.10.2:7002: i/o timeout
    
    ```
  不知道什么原因。照理说，第一种方式能够访问说明主机到容器的网络方式应该没什么问题，但是第二种就有问题想不通为什么。