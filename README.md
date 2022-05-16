说明文档
=======

## 环境搭建
* Golang v1.14+ 
* Beego v2.0.0+
* Mysql 5.6+ 主从数据库，可使用Docker快速搭建环境，参考文档 [agilejzl/master-slave-db](https://github.com/agilejzl/master-slave-db)

## 运行环境
* Step1: 配置数据库连接，复制example文件新建 conf/app.conf，修改数据库的连接参数为你的
* Step2: 安装依赖包，启动项目
```bash  
bee run                            # 启动项目
bee run -gendoc=true -downdoc=true # 启动并生成文档
```
打开浏览器查看文档 [http://localhost:8080/swagger](http://localhost:8080/swagger)

## Postman测试并发 (顺序请求)
* Postman可以选择接口集合，然后执行 "Run collection"，设置总执行回合即可测试。可下载 [接口集合文件](https://raw.githubusercontent.com/agilejzl/master-slave-db-demo/master/public/demo/DB2-Tester.postman_collection.json) ，然后导入到 Postman

## ab工具做压力测试 (并行请求)
* 以下是分别测试单个接口  
  获取我的商品列表
```bash
ab -c 100 -n 1000 -m GET -H 'Authorization: 1-1000' 'http://localhost:8080/api/products?scope=my'
```
发布一个我的商品
```bash
ab -c 100 -n 1000 -m POST -H 'Authorization: 1-1000' 'http://localhost:8080/api/products'
```
获取我的订单列表
```bash
ab -c 100 -n 1000 -m GET -H 'Authorization: 1-1000' 'http://localhost:8080/api/orders?scope=my'
```
创建待付款的订单
```bash
ab -c 50 -n 1000 -m POST -H 'Authorization: 1-1000' 'http://localhost:8080/api/orders'
```
随机给订单付款或关闭
```bash
ab -c 50 -n 500 -m PUT -H 'Authorization: 1-1000' 'http://localhost:8080/api/orders/random_id'
```

