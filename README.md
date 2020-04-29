# GOIPFS
你还毕不毕业了( ᖛ ̫ ᖛ )
## 上载
```
    cd  \!plague\!cat-\!miao/GOIPFS;git add . ;git commit -m "快速上传"; git push origin master 

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


## 雷区
### 1. 报错 : "Println call has possible formatting directive %v"
请将 Println 改为 fmt.Printf() \
[参考链接1] (https://stackoverflow.com/questions/53961617/call-has-possible-formatting-directive)

### 2. go get 报错 : i/o timeout
一般是代理可以解决 参照[一键解决 go get golang.org/x 包失败](https://www.jianshu.com/p/a537ee63d606)




