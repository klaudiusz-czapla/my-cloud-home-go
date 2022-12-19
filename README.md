# WD My Cloud Home CLI

Aim of that project is to implement translation layer between WD My Cloud Home Rest API and fuse. Thanks to that it will be possible to mount WD My Cloud Home device in the same way like any other file(s)/disk(s)/NAS device(s).

Besides that majority of the app functionality will be also available to the user in the form of CLI commands (based on the cobra library). This also allows to reduce number of tests needed because each use case, each command can be tested from the terminal beforehand.

Learn more about <strong>WD My Cloud Home</strong> from here: https://www.mycloud.com

Strongly inspired by:
- https://github.com/mnencia/mchfuse
- https://github.com/uname-yang/mycloudhome

> Thoughest part was done by the authors of libraries mentioned above - as of now documentation for WD My Cloud REST Api does not exist anywhere in public. 

> Security concerns at least for now are totally not in scope of the development (credentials in configuration files). User responsibility is to keep them secret.

> Application relies of existence of configuration json file named <em>config.json</em> stored in default location (place from application were started). You can override it easily by passing its custom path and custom name through the command line. Configuration file is meant to store sensitive values like secrets, password so please keep them secret. 

## Prerequisites
- Delve debugger installed
- Linters and other tools are in place
```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
go install golang.org/x/lint/golint@latest
go install golang.org/x/tools/cmd/goimports@latest
```

## Other resources you may find useful
- https://home.mycloud.com
- https://golangci-lint.run
- https://github.com/spf13/cobra
- https://github.com/spf13/viper

## How to

```
./my-cloud-home-go token --as="./token"
token=`cat token`
echo $token
./my-cloud-home-go refresh-token -t=$(echo $token)
```