# genzero

This is a tool to generate gozero service based on mysql.

## Install

```shell
go install github.com/licat233/genzero@latest
```

### Use from the command line

```text
$genzero -h
This is a tool to generate gozero service based on mysql.
The goctl tool must be installed before use.

Github: https://github.com/licat233/genzero

Usage:
  genzero [flags]
  genzero [command]

Available Commands:
  api         Generate .api files
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Create the default genzero configuration file in the current directory
  model       Generate model code
  pb          Generate .proto files
  upgrade     Upgrade genzero to latest version
  version     Print the version number of genzero
  yaml        Use yaml file configuration

Flags:
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

genzeroConfig.yaml will be created in the current directory

```shell
genzero init
```

A genzeroConfig.yaml configuration file will be created. [Sample file](./examples/genzeroConfig.yaml)

### By configuration file, generate the gozero and update it

Please ensure that genzeroConfig.yaml already exists in the current directory

```shell
genzero yaml
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
genzero -upgrade
```

### Configuration description

If there is a yaml configuration file, the configuration data in the file will be used first

### Update file content description

The content in this block will not be updated

```protobuf
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
