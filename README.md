# 喵贴后端服务器

前端为小程序喵贴，后端使用golang的gin框架；依赖使用go-modules。

用于学习研究，请勿滥用和恶意泄露。

# 代码使用方式
a)	下载安装go语言环境至任意文件夹。
```
 cd ~
 wget https://studygolang.com/dl/golang/go1.14.1.linux-amd64.tar.gz
 tar -zxf go1.14.1.linux-amd64.tar.gz
```

b)	创建一个文件夹，作为go的开发目录，比如workspace；同时在workspace中创建bin、pkg、src三个文件夹
```
 cd ~
 mkdir go_workspace && cd go_workspace
 mkdir bin
 mkdir pkg
 mkdir src
```

c)	设置 $GOROOT=go所安装的文件夹（.../go）；$GOPATH= workspace所在文件夹（.../workspace）；$PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```
 export GOROOT=~/go
 export GOPATH=~/go_workspace
 export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

d)	进入.../workspace/src，使用git获取代码：
```
 cd ~/go_workspace/src
 git clone https://github.com/Chaunhewie/miao_sticker_server.git
```

e)	利用govendor拉取依赖：
```
 go get -u github.com/kardianos/govendor
 govendor sync
```

f)	进入文件夹启动项目
```
cd ~/go_workspace/src/miao_sticker_server
go run main.go
```


# 注：
代码使用的端口为14488；


