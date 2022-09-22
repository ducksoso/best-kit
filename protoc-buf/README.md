

protoc --proto_path=$GOPATH/src --proto_path=. --go_out=. ./*.proto

a. 上面的句编译语句中，--proto_path用于表示要编译的proto文件所依赖的其他proto文件的查找位置，可以使用-I来替代。如果没有指定则从当前目录中查找。
b. --go_out有两层含义，一层是输出的是go语言对应的文件；一层是指定生成的go文件的存放位置。
c. --go_out=plugins=grpc:helloworld，这里使用了grpc插件。如果proto文件想在rpc中使用，可以在proto中定义接口如下：
service SearchService {
rpc Search(SearchRequest) returns (SearchResponse);
}
helloworld表示生成的文件存放地址。


protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./update.proto

a. --go_opt表示生成go文件时候的目录选项，如上面写时表示生成的文件与proto在同一目录。


import、go_package、package
a. package主要是用于避免命名冲突的，不同的项目（project）需要指定不同的package。
b. import，如果proto文件需要使用在其他proto文件中已经定义的结构，可以使用import引入。
c. option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb"; go_packge有两层意思，一层是表明如果要引用这个proto生成的文件的时候import后面的路径；一层是如果不指定--go_opt（默认值），生成的go文件存放的路径。
d. 需要注意的是package和go_package的含义。在官方给的文档中，package和go_package的最后一个单词不一样：


他们的含义分别是：package用于防止不同project之间定义了同名message结构的冲突，因为package名的一个作用是用于init方法中的注册：


而当go_package存在时，其最后一个单词是生成的go文件的package名字：

而当go_package不存在时，go文件件的package名字就变成了proto中package指定的名字了。

