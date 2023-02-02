# strman (String Manipulator) - to manipulate strings in the CLI

This package gives you opportunity to manipulate strings in local environemnt. Maybe your company restricts you from sharing the data in the open web, or you're afraid of web tools stealing your data for self purpose - String Manipulator is locally installed and is not sharing your information anywhere.

## Contents
* [Installation](#installation)
* [Usage](#usage)
  * [CLI](#cli)
* [Commands](#Commands)
  * [concat](#concat)
  * [random](#random)
  * [reverse](#reverse)
  * [inspect](#inspect)
  * [encode](#encode)
  * [decode](#decode)
  * [hash](#hash)
  * [jsonPretty](#jsonPretty)


## Prerequisites

You will need to `Go >= 1.18` installed on your computer

## Installation
First clone the repository somewhere in your $PATH. A common place would be within your $GOPATH. <br>

### Option 1: Install binary

Build and copy `strman` to your $GOPATH/bin:

```
$ make install
```
### Option 2: Build from source

This command will build the PagerDuty CLI binary, named `strman`. This binary will be created in the root of your project folder.

```
$ make build
```Example:

```
$ mkdir $GOPATH/src/github.com/openshift/pagerduty-short-circuiter
$ cd $GOPATH/src/github.com/openshift/pagerduty-short-circuiter
$ git clone https://github.com/openshift/pagerduty-short-circuiter.git .
To list all commands:
```shell
$ strman help
```

To run command with the full name:
```shell
$ strman <command> [flags] <string>
```

## Commands

### concat
Concats one or more of the given inputs provided and displays it on the output. Use the separator flag to define the separator between strings.
```shell
USAGE
  $ strman concat [--sep=<sep>] <string1> <string2> <string3>...

EXAMPLE
  $ strman concat "test" "hgsddshj" "sdjhsd" --sep=+
  #  test+hgsddshj+sdjhsd
```

### random
Generates a random string of the specified length.
```shell
USAGE
  $ strman random <length>

EXAMPLE
  $ strman random 10
  #  IQoh54Itb5
```

### reverse
Reverses the given string input.
```shell
USAGE
  $ strman reverse <length>

EXAMPLE
  $ strman reverse TestString
  #  gnirtStseT
```

### inspect
Provides more details about the string. Displays the number of characters in the string.
```shell
USAGE
  $ strman inspect [options] string

OPTION


EXAMPLE
  $ strman inspect test234
  #  'test234' has 7 chars

  $ strman inspect test234 --digits
  # 'test234' has 3 digits
```

### encode
Encodes given string in specific format
```shell
USAGE
  $ strman encode --format=<format> <string>

FORMATS
  - base64
  - ascii
  - binary
  - hex
  
EXAMPLE
  $ strman encode --format=base64 "This is my string"
  #  VGhpcyBpcyBteSBzdHJpbmc=
```

### decode
Decodes given string from specific format to UTF-8 string
```shell
USAGE
  $ strman decode --format=<format> <string>

FORMATS
  - base64
  - ascii
  - binary
  - hex

  
EXAMPLE
  $ strman decode --format=base64 VGhpcyBpcyBteSBzdHJpbmc=
  #  This is my string
```

### hash
Hashes string in specified format
```shell
USAGE
  $ strman hash --format=<format> <string>

FORMATS
  - md5
  - sha1
  - sha256
  - sha512
  
EXAMPLE
  $ strman hash --format=sha256 "This is my string"
  #  9da6c02379110815278b615f015f0b254fd3d5a691c9d8abf8141655982c046b
```

### jsonPretty
Makes your ugly JSON string pretty with specified spaces amount
```shell
USAGE
  $ strman jsonPretty [--spaces <n>] <JSON>

EXAMPLE
  $ strman jsonPretty "{\"test\":\"JSON\"}"
  #  {
  #      "test": "JSON"
  #  }
```