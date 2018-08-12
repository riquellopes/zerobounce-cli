Zerobounce Cli
==============

The [zerobounce](https://www.zerobounce.net/) [cli](https://en.wikipedia.org/wiki/Command-line_interface) is a new option to check up information about the email by line command. After install the lib create your account at [here](https://www.zerobounce.net/), get your token application and define the environment. 

#### Export API_KEY_ZERO:
```sh
$ export API_KEY_ZERO={ZEROBOUNCE_TOKEN}
```
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
