FROM golang:1.20
ENV ZBE_PATH /biz-code/kobe
ENV GOPROXY https://goproxy.cn,direct
ENV GOPRIVATE gitee.com/zhubaoe-go/jordan

COPY ./truss /biz_code/jordan
COPY ./protoc-25.1-linux-x86_64/bin/protoc /bin/protoc

WORKDIR /biz_code/jordan

RUN ls && \
    cd src && \
    make dependencies &&  make && \
    go get -u github.com/golang/protobuf/protoc-enum_gen-go  && \
    cp -rf  ../github.com  $GOPATH/src && \
    cp -rf  ../google  $GOPATH/src