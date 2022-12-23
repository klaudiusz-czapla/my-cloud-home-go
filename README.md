# WD My Cloud Home CLI

Aim of that project is to implement translation layer between WD My Cloud Home Rest API and fuse. Thanks to that it will be possible to mount WD My Cloud Home device in the same way like any other network storage devices.

Besides that majority of the app functionality will be also available to the user in the form of CLI commands (based on the cobra library). This also allows to reduce number of tests needed because each use case, each command can be tested from the terminal beforehand.

Learn more about <strong>WD My Cloud Home</strong> from here: https://www.mycloud.com

Strongly inspired by:
- https://github.com/mnencia/mchfuse
- https://github.com/uname-yang/mycloudhome

> Thoughest part was done by the authors of libraries mentioned above - as of now documentation for WD My Cloud REST Api does not exist anywhere in public.

> Security concerns at least for now are totally not in scope of the development (credentials in configuration files). User responsibility is to keep them secret.

> Application relies of existence of configuration json file named <em>config.json</em> stored in default location (place from application were started). You can override it easily by passing its custom path and custom name through the command line. Configuration file is meant to store sensitive values like secrets, password so please keep them secret. 

```
{
    "username": "",
    "password": "",
    "clientId": "",
    "clientSecret": "",
    "deviceName": ""
}
```

## Prerequisites

From dev perspective
- go

    After installing go according to docs you may add those lines to your .profile file:
    ```
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
    and then do sth like this:
    ```
    source ~/.profile
    ```

- delve debugger installed
- linters
    ```
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
    #golangci-lint help linters
    go install golang.org/x/lint/golint@latest
    go install golang.org/x/tools/cmd/goimports@latest
    ```

From runtime perspective
- config.json

## Other resources you may find useful
- https://home.mycloud.com
- https://golangci-lint.run
- https://github.com/spf13/cobra
- https://github.com/spf13/viper

## How to
- Get token and save it for later usage
    ```
    ./my-cloud-home-go token --to=./token
    ```

- Refresh the token
    ```
    ./my-cloud-home-go token --to="./token"
    token=`cat token`
    ./my-cloud-home-go refresh-token --token=$(echo $token)
    # or
    ./my-cloud-home-go refresh-token --from=./token
    ```

- Get device info
    ```
    ./my-cloud-home-go device-info --from=./token
    ```