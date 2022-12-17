## WD My Cloud Home CLI

Aim of that project is to implement translation layer between WD My Cloud Home Rest API and fuse. Thanks to that it will bepossible to mount WD My Cloud Home device in the same way like any other files/disks/NAS devices.

Besides that majority of the app functionality will be also available to the user in the form of CLI commands (based on cobra library). This also allows to reduce number of tests needed because each use case, each command can be tested from the terminal beforehand.

Learn more about <strong>WD My Cloud Home</strong> from here: https://www.mycloud.com

Strongly inspired by:
- https://github.com/mnencia/mchfuse
- https://github.com/uname-yang/mycloudhome

> Thoughest part was done by authors of those libraries - documentation for WD My Cloud REST Api does not exist anywhere in public. 

> Security concerns at least for now are totally not in scope of the development(credentials in configuration files). User responsibility is to keep them secret.

> Application relies of existence of configuration json file named config.json stored in default location (place from application were started). You can override it easily by passing its custom path and custom name through the command line. Configuration file is meant to store sensitive values like secrets, password so please keep it secret. 

## Other resources you may find useful
- https://home.mycloud.com
- https://github.com/spf13/cobra
- https://github.com/spf13/viper