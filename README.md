# GOIPFS
你还毕不毕业了( ᖛ ̫ ᖛ )
## 下载
`git pull origin master  `
## 上载
```
    cd $GOPATH/src/github.com/\!plague\!cat-\!miao/GOIPFS;git add . ;git commit -m "快速上传"; git push origin master 

```
## install
### for Linux

1. 安装IPFS终端
  - [go-ipfs编译](https://github.com/ipfs/go-ipfs)
  - [IPFS终端下载](https://github.com/ipfs/go-ipfs/releases)
2. 安装 API
  1. [github搜索go-ipfs-api](https://github.com/ipfs?q=api&type=&language=)
  2. 于本地$GOPATH下安装
3. 更新一下依赖
  - 简单: 下载个[main.go]((https://blog.csdn.net/Aaron_Kings/article/details/105256309?utm_medium=distribute.pc_relevant.none-task-blog-OPENSEARCH-2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-OPENSEARCH-2)) 开始`go get -u` 
  - 复杂一个一个`git clone`
4. test
  - go run main.go
  - go test -v ipfs_api_test.go

### 参照贴
1. [go-ipfs-api在项目中的使用](https://blog.csdn.net/qq_37133717/article/details/83375644)
2. [ipfs在linux的安装、同步、测试](https://blog.csdn.net/wangkaizheng123/article/details/83089897)
3. [go-ipfs-api初步使用](https://blog.csdn.net/Aaron_Kings/article/details/105256309?utm_medium=distribute.pc_relevant.none-task-blog-OPENSEARCH-2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-OPENSEARCH-2)
4. [go test命令（Go语言测试命令）完全攻略](http://c.biancheng.net/view/124.html)
## GetStart

## 墓碑
### Makefile
 - [Makefile介绍](https://wiki.ubuntu.org.cn/%E8%B7%9F%E6%88%91%E4%B8%80%E8%B5%B7%E5%86%99Makefile:MakeFile%E4%BB%8B%E7%BB%8D)
 - [Makefile 使用总结](https://www.cnblogs.com/wang_yb/p/3990952.html)
   - 我们的GO项目中基本都是伪目标 
### window 配置GO
[win配置GO教程](https://blog.csdn.net/youshijian99/article/details/81782875)
### MySQL
#### MySQL && GO
- [Go操作MySQL](https://www.cnblogs.com/nickchen121/p/11517430.html)
  - &quot显示错误 应为“"(引号)
- [go中mysql使用](https://blog.csdn.net/chinabestchina/article/details/90745511)
- [go操作数据库 Go-SQL-Driver/MySQL 使用详解](https://www.cnblogs.com/zhja/p/5604553.html)
  - db, err := sql.Open("mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")
#### MySql 同步方案
- [mysql数据备份与同步](https://www.cnblogs.com/yhlboke-1992/p/11685436.html)
  - 因为贫穷而选择 主从
- [MySQL数据同步方案实战-Canal ](https://www.jianshu.com/p/926956e7c058)
- [Mysql主主同步方案](https://blog.csdn.net/qq_35298894/article/details/79023093)
- [详解MySQL双活同步复制四种解决方案](https://www.jb51.net/article/146131.htm)
#### MySQL 表格关系
 - 外键 

    为表格设置外键，可以限制插入，只有当全部外键的值都在关联表格的主键中存在，方可插入。
 - 联合主键
 
    主键主要防止重复、多主键指当多个主键同时重复时，进行拦截。
 - 我的设计
 
 | video | - | Ownership| - | E/N/C |
 - 参考链接
   - [MySQL - 数据库设计（表之间的 3 种关联关系）](https://blog.csdn.net/Dream_Weave/article/details/85172796)
   - [用MySQL Workbench为一对多，多对多的关系创建外键关联](https://blog.csdn.net/qq_32808253/article/details/78941625)

### GO
#### panic
- 介绍 

  在函数内部调用 panic 会立即终止当前函数的执行，由当前调用栈逐层返回，一直到最顶层的 main 函数或是被某一层的 recover 捕捉到。
- 打印内容样例:
  > panic: assignment to entry in nil map [recovered] \
  >  panic: assignment to entry in nil map
- [Go异常处理 ](https://www.cnblogs.com/baiyuxiong/p/4770980.html) \
#### map
[go map 基本操作](https://www.cnblogs.com/bingzhen/p/10503967.html)
## 雷区
### 1. 报错 : "Println call has possible formatting directive %v"
请将 Println 改为 fmt.Printf() \
[参考链接1] (https://stackoverflow.com/questions/53961617/call-has-possible-formatting-directive)

### 2. 解决 go get 被墙问题
- 一般报错：
  - connectex: A connection attempt failed because the connected party did not properly 。。。 
  - i/o timeout
- 解决方案  _GOPROXY = "https://goproxy.io"_
  - liunx 参照[一键解决 go get golang.org/x 包失败](https://www.jianshu.com/p/a537ee63d606)
  - window 有些特殊 [windows下解决 go get golang.org/x 包失败](https://www.jianshu.com/p/663b2dbf55ce)
### 3. go get -u 报错  go get .: /xxxx is not a package in module rooted at /XXXX
一般是你的这个项目没写main 或者没有GO文件。go get 无法建立依赖树 （子目录有GO文件不算）
通过在文件下go build 可以达成同样的效果
#### 4. MySQL导入出现问题：Unknowncollation:utf8mb4_0900_ai_ci 
- 报错原因：

   生成转储文件的数据库版本为8.0,要导入sql文件的数据库版本为5.6,因为是高版本导入到低版本，引起1273错误
- 解决方法：

   打开sql文件，将文件中的所有
`utf8mb4_0900_ai_ci`替换为`utf8_general_ci` \
`utf8mb4`替换为`utf8` \
保存后再次运行sql文件，运行成功
- 参考链接:\
 <http://blog.sina.com.cn/s/blog_591c97810102yszr.html>
## 扩展阅读
### 负载均衡
[Keepalived之——Keepalived + Nginx 实现高可用 Web 负载均衡](https://blog.csdn.net/l1028386804/article/details/72801492)
### 高可用
[什么是高可用性](https://baike.baidu.com/item/%E9%AB%98%E5%8F%AF%E7%94%A8%E6%80%A7/909038?fr=aladdin)
：用来描述一个系统经过专门的设计，从而减少停工时间，而保持其服务的高度可用性。 \
工作原理：多台主机一起工作，各自运行一个或几个服务，各为服务定义一个或多个备用主机，当某个主机故障时，运行在其上的服务就可以被其它主机接管。
### 集群
[科普：并行计算、分布式计算、集群计算和云计算](https://blog.csdn.net/datamining2005/article/details/85229541)
### zoomkeeper
[Zookeeper入门看这篇就够了](https://blog.csdn.net/java_66666/article/details/81015302)
[Go语言连接 zookeeper](https://blog.csdn.net/bingfeilongxin/article/details/88578887)
### MQ消息队列-kafaka
[Kafka 集群](https://www.jianshu.com/p/430454d9e30b)
### MQ消息队列-Rocketmq
[kafaka和Rocketmq的区别](https://www.cnblogs.com/ynyhl/p/11320797.html) 
### canal
- [使用Canal实现MySQL的数据实时同步](https://blog.csdn.net/qq_23329167/article/details/89873870)
  canal的工作原理：
  1. canal模拟mysql slave的交互协议，伪装自己为mysql slave，向mysql master发送dump协议
  2. mysql master收到dump请求，开始推送binary log给slave(也就是canal)
  3. canal解析binary log对象(原始为byte流)
- [canal实现mysql数据实时同步es](https://blog.csdn.net/zhangcongyi420/article/details/105176534?utm_medium=distribute.pc_relevant.none-task-blog-OPENSEARCH-1&depth_1-utm_source=distribute.pc_relevant.none-task-blog-OPENSEARCH-1)