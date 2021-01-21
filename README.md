# proto

根据proto生成对应语言的代码：

~~~bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/demo.proto
~~~