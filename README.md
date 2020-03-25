# rc

This is a golang cli tool for sending rest commands to a global cache itach wifi. The rc tool is very much a work and progress and no plans for making this a universal tool. 

# Hardware
Global Cache iTach Flex Wi-Fi (Flex-WF)
https://www.globalcache.com/products/flex/

# Api reference
https://gcapi.docs.apiary.io/#reference

# Ircode Database
https://irdb.globalcache.com/

# iConvert utility
macOS https://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?mt=11&path=mac%2ficonvertgc

Windows https://www.globalcache.com/files/software/iConvert.zip

# usage
```
$ rc help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  rc [command]

Available Commands:
  config      A brief description of your command
  down        Volume down
  hdmi1       A brief description of your command
  hdmi2       A brief description of your command
  help        Help about any command
  mute        Mutes the speakers
  off         Powers off the display
  on          Power ON the display
  tvdown      A brief description of your command
  tvup        A brief description of your command
  up          Volume up

Flags:
      --config string   config file (default is $HOME/.rc.yaml)
  -h, --help            help for rc
  -t, --toggle          Help message for toggle

Use "rc [command] --help" for more information about a command.
```
