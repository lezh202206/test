#bin/bash
biz_code=/Users/lezh/Desktop/GO/src/

cd $biz_code/jordan/tools/truss/src &&
  make dependencies &&
  make &&
  cd /biz-code/jordan/tools/truss &&
  go get -u github.com/golang/protobuf/protoc-gen-go &&
  cp -rf github.com $GOPATH/src &&
  cp -rf google $GOPATH/src



