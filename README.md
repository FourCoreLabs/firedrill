# firedrill 🧯

Malware simulation harness. Build native binaries for Windows, Linux and Mac simulating malicious behaviours. Test the effectiveness of your endpoint security controls against malware.

![firedrill](https://i.imgur.com/flySzca.png)

## Usage

- Requires Go 1.17+, GNU make.

Windows

```
$ make ransomware
$ ransomware.exe
```

Linux/bash

```
$ GOOS=windows make ransomware
```

## Ransomware Simulation

The ransomware simulation is present in *cmd/ransomware*, you can build it using `make ransomware` or download the binary from from the latest release [here](https://github.com/FourCoreLabs/firedrill/releases).

This simulation involves three tasks, in this order, which is typical of a ransomware:

- Encryption of files on the filesystem (only test files dropped by the binary and ).
- Dropping a ransom note on the desktop.
- Changing the system wallapaper through registry keys (and restoring it after some time).

The simulation is designed to be safe and does not perform any destructive actions on the system. 
