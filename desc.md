## 环境设置

* mkdir -P /opt/go-user/Golang, 创建代码保存路径

* export GoHome=/opt/go-user/Golang, 设置Go的代码主目录

* cd $GoHome && git clone repository, clone代码

* export GOPATH=$GoHome/work, 将代码的路径设置为GOPATH

* export PATH=$PATH:$GOPATH/bin, GOPATH下的src保存源代码，bin保存编译后的二进制，同时将bin目录也加入到PATH环境变量中，方便后面直接执行编译后的程序

## 代码编译和包的安装

* go build/install hello, src下有包hello，可以直接install和build
