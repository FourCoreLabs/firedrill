# 🧯firedrill: A malware simulation harness

[![goreleaser](https://github.com/FourCoreLabs/firedrill/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/FourCoreLabs/firedrill/actions/workflows/goreleaser.yml)

TL;DR: firedrill is an open-source library from FourCore Labs to build malware simulations easily. We have built a set of four different attack simulations for you to use and build on top of: Ransomware Simulation, Discovery Simulation, a UAC Bypass and a Persistence Simulation.

Organizations invest a whole lot in their security controls and tooling for the security teams to be efficient. They might put in 10s of man-hours tuning a single security control to their needs building out detection rules, identifying best practice configuration and setting up automation. However, the task shouldn't end right here, it is crucial to test the effectiveness of these security systems against attackers. Usually, this is done with pentesting and red teaming activities done either in-house or by external teams. This results in immense value for the organization as its build confidence in security controls against in the event of a real attack.

![firedrill](https://i.imgur.com/flySzca.png)

Read more about firedrill on our [**blog**](https://fourcore.io/blogs/firedrill-open-source-attack-simulation).

## Ransomware Simulation

The ransomware simulation consists of typical behaviour of a ransomware.

This includes, in this order:
- Encryption of files on the filesystem (only test files dropped by the binary and ).
- Dropping a ransom note on the desktop.
- Changing the system wallapaper through registry keys (and restoring it after some time).

Sandbox Analysis

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/21b95de03a83883b67fe14d9d517782f73276649378fbb4fca632c89410c2ba9/61dff7ee07ae9c2e3f3842e4) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/50b1305b-16f1-4f12-8df7-97cdaf468cec) |

## Discovery Simulation

The discovery simulation consists of simulation of a malware executing three techniques from the Discovery tactic in MITRE ATT&CK, performing reconnaisance of system information which is used for further exploiting the system:

This includes, in this order:
- Discovering the running processes on the system.
- Discovering the peripherals present on the system.
- Discovering the softwares installed on the system with their respective versions.

Sandbox Analysis 

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/c8fcd8419bf11385bdddc9cfd8017226493365ff97d2232f9283fbe6309830bc/61dff860d9a3de1d1f04a1fb) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/0e7f5d02-1f69-43a5-b6b3-65c3cffcd21d) |

## UAC Bypass Simulation

The UAC Bypass simulation consists of malware using the fodhelper.exe utility available from Windows 10 to achieve local privilege escalation by creating a registry structure to execute arbitrary commands with adminstrator privileges:

This includes, in this order:
- Create a new registry structure in `HKCU:\Software\Classes\ms-settings\` and start `notepad.exe` with admin privileges bypassing UAC.

Sandbox Analysis

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/98ee778d81174276c74ef2039163b48479b9b1d798770ea434d8d54cb35390b0) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/5a60925a-7195-4b4c-8f23-5db9dfbbea5a) |


## Registry Run Key Persistence Simulation

This is a simulation of a persistence techniques which use registry Run keys to achieve persistence for arbitrary payloads. These keys include: `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`, `EY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce`, `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices`, `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce`.

This includes, in this order:
- Add a value in the registry key at `HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` to execute a sample payload embedded in the binary.
- Delete the value from `HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` to bring back it to it's original state for a safe simulation.

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/353aa45090090f298af8b1d7135b33ea03c7b5b431c31367e9468366aff227b2) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/9623c141-4927-43a3-acb9-38c65b6c9a5e) |

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
$ runkeyregistry.exe
```

## Linux/bash
```
$ GOOS=windows make ransomware # and so on
```
