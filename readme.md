# bt
a wrapper for [bluetoothctl](https://wiki.archlinux.org/index.php/bluetooth) with command completion and easy connection to previously paired devices via a configuration file

# install
> go get github.com/mrgarelli/bt

# configure
```
---
devices:
  echo: "Echo Plus-CNX"
  headset: "Tribit XFree Tune"
```

# run
> ./bt

### output:

```
a package for managing bluetooth devices with bluetoothctl

Usage:
  bt [flags]
  bt [command]

Available Commands:
  beDiscoverable make the current device discoverable
  connect        <device>: connect to a device
  ctrl           simply opens bluetoothctl
  disconnect     <device>: disconnect from a device
  help           Help about any command
  info           get info for all paired devices
  scan           scans for bluetooth devices

Flags:
      --completion string   generate shell completion
      --config string       config file (default is $HOME/.bt.yaml)
  -h, --help                help for bt

Use "bt [command] --help" for more information about a command.
```
