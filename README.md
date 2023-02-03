# wg-keygen-rep

wireguard keypair generator with Salt string

## Usage

```
% ./wg-keygen-rep --salt example-salt-string
priv: MGm3G4CsUB17qmGVEyqhjW/l9ervkv33iMeSnnmEbXw=
pub: +Gf1tUb5OMk+UD7pWZPQkD2K/CqS+NZd85Z7CGmpCFo=

# can regenerate same keys with same salt string
% ./wg-keygen-rep --salt example-salt-string
priv: MGm3G4CsUB17qmGVEyqhjW/l9ervkv33iMeSnnmEbXw=
pub: +Gf1tUb5OMk+UD7pWZPQkD2K/CqS+NZd85Z7CGmpCFo=
```

## Help

```
% ./wg-keygen-rep -h     
Usage:
  wg-keygen-rep [OPTIONS]

Application Options:
  -s, --salt=    salt string for generating private key
      --json     output with JSON format
  -v, --version  Show version

Help Options:
  -h, --help     Show this help message
```

## Install

Please download release page

