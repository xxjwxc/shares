NAME := shares
all: # 构建
	make clear
build:
	make clear
	./tools/gormt -o internal/model
	go build -o $(NAME) main.go 
	./$(NAME) debug
run:
	# make clear
	# ./tools/gormt -o internal/model
	go build -o $(NAME) *.go 
	./$(NAME) debug
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(NAME).exe main.go 
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(NAME) main.go 
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(NAME) main.go 
clear: 
	test ! -d internal/model/ || rm -rf  internal/model/*
	test ! -d err/ || rm -rf err/
	test ! -f $(NAME) || rm -rf $(NAME)
	test ! -f $(NAME).exe || rm -rf $(NAME).exe
	rm -rf ./rpc/$(NAME)/*.go # 删除proto文件夹下所有go文件（谨慎操作）
# gen:
# 	- mkdir ../rpc
# 	- mkdir ./rpc/$(+)
# 	test -L rpc || ln -s ../rpc ./
# 	protoc --proto_path="../apidoc/proto/$(SaverName)/" --gmsec_out=plugins=gmsec:./rpc/$(SaverName) hello.proto
gen:
ifeq ($(LANG),) # windows
	# test ! -d rpc/ || rm -rf  rpc/
	# test -h rpc || ln -s ../rpc ./
else
	# test -h rpc || ln -s ../rpc ./ # with linux/mac
	test -d rpc || ln -s ../rpc ./
endif
	go generate
orm: # gormt 生成 orm代码
	./tools/gormt -o internal/model
install:
	../proto_install.sh
source_install:
	../proto_install.sh
master:
	go get -u github.com/xxjwxc/public@master
	go get -u github.com/xxjwxc/ginrpc@master
	go get -u github.com/gmsec/micro@master
	go get -u github.com/gmsec/goplugins@master
crontab:
	crontab testCrontab 
# crontab -l 查看定时任务
# crontab -r 删除定时任务