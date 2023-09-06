# redis集群
## 集群搭建步骤
1. 运行命令`docker-compose up -d`创建容器
2. 运行`docker exec -it redis1 bash`进入容器redis1
3. 在redis1容器环境中运行`redis-cli --cluster create 172.16.10.1:7001 172.16.10.2:7002 172.16.10.3:7003 172.16.10.4:7004 172.16.10.5:7005 172.16.10.6:7006 --cluster-replicas 1`

以上过程就完成了redis集群环境的搭建，创建了一个3主3从的redis集群。

## 数据读写
1. 写入数据：`redis-cli -c -p 7001 set name leebai`
2. 读取数据：`redis-cli -c -p 7001 get name`
> 以上2条命令的执行都需要在redis1的容器环境下执行
