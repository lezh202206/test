注意： 使用前，相关的环境配置已经完成，见samples/README.md 的环境配置章节

1. 本项目已编译好windows下使用的工具包，直接解压truss_win.zip中的内容到$GOPATH/bin目录下(非强制，建议第三方工具统一放到一个目录)

2. 获取可以通过src编译生成工具
   windows下直接运行wininstall.bat, 执行完成后会在$GOPATH/bin目录下看到truss工具;

3. 将当前目录下的github.com、google 目录拷贝到`$GOPATH`的src目录下(truss protoc生成时使用的一些import包)

4. 创建服务
例如，要创建merchant服务，我们可以在services目录下创建merchant目录，
然后创建merchant.proto文件，定位到该文件目录，执行：
```
truss merchant.proto
```
自动生成的代码在merchant-service目录下，当前只允许修改handler，svc/transport_http两个目录下的内容，其他的文件每次生成会被覆盖；

# 源码的修改
## 模板
 模板位于 src/gengokit/template 目录下, 修改模板后, 编译, `go-bindata`会自动生成模板的编译文件(go文件)
##
