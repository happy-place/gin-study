## html文件二进制打包

```aidl
检查GOPATH目录
echo $GOPATH

稍后插件会安装在$GOPATH/bin目录下
vi ~/.bash_profile
export PATH=$PATH:$GOPATH/bin
source ~/.bash_profile

下载打包插件
go get github.com/jessevdk/go-assets-builder

mkdir html
编写 bar.tmpl、index.tmpl

将.html所在文件打包为go脚本，稍后可以编译为二进制执行文件
go-assets-builder ./html -o assets.go

```

