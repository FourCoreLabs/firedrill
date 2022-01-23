# 🧯firedrill: A malware simulation harness

Organizations invest a whole lot in their security controls and tooling for the security teams to be efficient. They might put in 10s of man-hours tuning a single security control to their needs building out detection rules, identifying best practice configuration and setting up automation. However, the task shouldn't end right here, it is crucial to test the effectiveness of these security systems against attackers. Usually, this is done with pentesting and red teaming activities done either in-house or by external teams. This results in immense value for the organization as its build confidence in security controls against in the event of a real attack.

![firedrill](https://i.imgur.com/flySzca.png)

We have built [firedrill](https://github.com/FourCoreLabs/firedrill), a malware simulation harness to help security teams simplify and improve the security validation process. As the name suggests, it is alike a fire drill for your security infrastructure. You can build safe and reproducible malware simulations in the Go programming language and produce native executables binaries for Windows, Mac and Linux. These binaries can produce execution behaviours that can represent the activities of a malware. With the flexibility of Go, we can build a variety of simulations including Ransomware attacks, simulations of techniques in the popular MITRE ATT&CK matrix and many more. We are open sourcing firedrill under the MIT License and inviting security practitioners and Go developers to collaborate with us to identify impactful simulations for their own organizations.

In our first release, we are releasing a ransomware simulation and a discovery simulation. You can build it yourself, with instructions on our [repository](https://github.com/FourCoreLabs/firedrill) or you can just grab the `windows/amd64` binary releases for both simulation from GitHub releases. Both of the provided simulations are designed to be safe, to produce a reproducible runtime execution beahviour of a malware. It does not perform any destructive actions on the system.

## Ransomware Simulation

The ransomware simulation consists of typical behaviour of a ransomware.

This includes, in this order:
- Encryption of files on the filesystem (only test files dropped by the binary and ).
- Dropping a ransom note on the desktop.
- Changing the system wallapaper through registry keys (and restoring it after some time).

## Discovery Simulation

The ransomware simulation consists of simulation of a malware executing three techniques from the Discovery tactic in MITRE ATT&CK, performing reconnaisance of system information which is used for further exploiting the system:

This includes, in this order:
- Discovering the running processes on the system.
- Discovering the peripherals present on the system.
- Discovering the softwares installed on the system with their respective versions.
Malware simulation harness. Build native binaries for Windows, Linux and Mac simulating malicious behaviours. Test the effectiveness of your endpoint security controls against malware.

## Usage

- Requires Go 1.17+, GNU make.

## Windows

Ransomware Simulation
```
$ make ransomware
$ ransomware.exe
```

Discovery Simulation
```
$ make discovery
$ discovery.exe
```

UAC Bypass Simulation
```
$ make uac_bypass
$ uac_bypass.exe
```

Registry Run Key Simulation
```
$ make registry_run
$ registry_run.exe
```

## Linux/bash
```
$ GOOS=windows make ransomware # and so on
```
