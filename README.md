# Netinfo
Netinfo
- Display local IP information
- setup as a server to tell the client its ip information 
- written in golang 

## Usage
```
Version:
  1.00

Usage:
  netinfo {Command} [Option]

Command:
  -s                : start http server
  -h                : show help
  -v                : show version

Option:
  -a  <IP>          : set server IP
  -p  <Port>        : set server port

Example:
  1) netinfo
  2) netinfo -s
  3) netinfo -s -a 127.0.0.1 -p 1996
  4) netinfo -h
  5) netinfo -v
```

## Build
### Example
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/netinfo.git

cd netinfo

go build -v -trimpath -ldflags "-s -w"
```
