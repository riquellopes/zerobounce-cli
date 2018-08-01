Zerobounce Cli
==============

The [zerobounce](https://www.zerobounce.net/) cli is an option to check up information about the email in line command.

#### Usage Examples

Type zerobounce-cli to list all available commands or zerobounce-cli help COMMAND for specific command.

**Get the validation result for one email address.**
```sh
    $ zerobounce-cli validate contato@henriquelopes.com.br
```

**Get the number of credits available in your ZeroBounce account.**
```sh
    $ zerobounce-cli getcredits
```

**Get the validation result for one email address, using an IP.**
```
    $ zerobounce-cli validatewithip contato@henriquelopes.com.br 127.0.0.1
```
