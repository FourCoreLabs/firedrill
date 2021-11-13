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
