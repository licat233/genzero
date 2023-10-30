# genzero

根据mysql，一键生成gozero框架的增删改查等相关代码，涉及*.api与*.proto文件的生成，以及model代码的生成，api logic代码的生成，rpc logic代码的生成。

This is a tool to generate gozero service based on mysql.

## Install

```shell
go install github.com/licat233/genzero@latest
```

## Usage steps - 使用步骤

1. 【初始化】
   1. 使用`genzero init config`命令生成配置文件: genzero.yaml，然后根据自己的需求编辑genzero.yaml;
2. 【生成.api或者.proto文件】
   1. 使用`genzero start`命令生成 xxx.api 或者 xxx.proto 文件，然后根据自己的需求编辑 xxx.api 或者 xxx.proto文件;
3. 【生成model文件、api服务文件或者rpc服务文件】
   1. 使用`goctl model mysql ddl --src="xxx.sql" --dir="model" --style="goZero"`命令生成model包文件
   2. 使用`goctl api go --api xxx.api --dir ./api --style goZero`命令生成api服务文件;
   3. 使用`goctl rpc protoc xxx.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style goZero`命令生成rpc服务文件；
4. 【生成model代码、api logic代码或者rpc logic代码】
   1. 使用`genzero start`命令生成相关的gozero服务代码，可选择的模块有"api服务中的logic代码"或者"rpc服务中的logic代码"，以及"model包的extend方法代码"

### Use from the command line

```text
$genzero -h
This is a tool to generate gozero service based on mysql.
The goctl tool must be installed before use.
current version: v1.1.5-alpha.9
Github: https://github.com/licat233/genzero

Usage:
  genzero [flags]
  genzero [command]

modules:
  api         Generate .api files
  logic       Modify logic files, this feature has not been developed yet
  model       Generate model code
  pb          Generate .proto files

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Create the default genzero configuration file in the current directory
  start       Use yaml file configuration to start genzero
  upgrade     Upgrade genzero to latest version
  version     Print the version number of genzero

Flags:
      --conf string              file location for yaml configuration (default "genzero.yaml")
      --dev                      dev mode, print error message
      --dsn string               data source name (DSN) to use when connecting to the database
  -h, --help                     help for genzero
      --ignore_columns strings   ignore column string, default is none，split multiple value by ','
      --ignore_tables strings    ignore table string, default is none，split multiple value by ','
      --src string               sql file to use when connecting to the database
      --tables strings           need to generate tables, default is all tables，split multiple value by ','
  -v, --version                  version for genzero

Use "genzero [command] --help" for more information about a command.
```

### By cmd, generate the xxx.proto file and update it

```shell
genzero pb --src="../sql/admin.sql" --service_name="admin" --multiple=false
```

### By cmd, generate the xxx.api file and update it

```shell
genzero api --src="../sql/admin.sql" --service_name="admin-api" --jwt="Auth" --middleware="AuthMiddleware" --prefix="/v1/api/admin" --multiple=false --ignore_tables="jwt_blacklist"
```

### By cmd, generate the xxxModel_extend.go file and update it

```shell
genzero model --src="../sql/admin.sql" --service_name="admin" --dir="model"
```

### Generate a configuration file

genzero.yaml will be created in the current directory

```shell
genzero init config
```

A genzero.yaml configuration file will be created. [Sample file](./examples/genzero.yaml)

### By configuration file, generate the gozero and update it

Please ensure that genzeroConfig.yaml already exists in the current directory

```shell
genzero start --src=./genzero.yaml
```

### Use [goZero](https://github.com/zeromicro/go-zero)'s goctl tool to generate service code

Generate api service code

```shell
goctl api go -api admin.api -dir ./api -style goZero
```

Generate rpc service code

```shell
goctl rpc protoc "admin.proto" --go_out="./rpc" --go-grpc_out="./rpc" --zrpc_out="./rpc"
```

Generate model service code

```shell
goctl model mysql ddl --src "admin.sql" -dir . --style goZero
```

### Upgrade genzero to latest version

```shell
genzero upgrade
```

### Configuration description

If there is a yaml configuration file, the configuration data in the file will be used first

### Update file content description

The content in this block will not be updated

```proto
// The content in this block will not be updated
// 此区块内的内容不会被更新
//[custom messages start]

message Person {
  int64 id = 1;
  string name = 2;
  bool is_student = 3;
}

//[custom messages end]
```

### examples

[see](./examples/)

### Thanks

+ [goZero](https://github.com/zeromicro/go-zero)
