# Netinfo

## Features
- Display local IP information
- Setup as a server to tell client its ip information

## Usage
```
Usage:
  netinfo {Command} [Option]

Command:
  -server           : start http server
  -h                : show help
  -v                : show version

Option:
  -address <IP>     : set server IP
  -port    <Port>   : set server port
  -netdb            : use net ip database to get ip info

Example:
  1) netinfo
  2) netinfo -server
  3) netinfo -server -address 127.0.0.1 -port 1996 -netdb
  4) netinfo -h
  5) netinfo -v
```

## Install
```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/netinfo https://github.com/gek64/netinfo/releases/latest/download/netinfo-linux-amd64
chmod +x /usr/local/bin/netinfo

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/netinfo https://github.com/gek64/netinfo/releases/latest/download/netinfo-freebsd-amd64
chmod +x /usr/local/bin/netinfo
```


## Install Service
### Linux(systemd)
```sh
curl -Lo /etc/systemd/system/netinfo.service https://github.com/gek64/netinfo/raw/main/service/netinfo.service
systemctl enable netinfo && systemctl start netinfo
```
### FreeBSD(rc.d)
```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/netinfo https://github.com/gek64/netinfo/raw/main/service/netinfo
chmod +x /usr/local/etc/rc.d/netinfo
server netinfo enable && server netinfo start
```


## Compile
### How to compile if prebuilt binaries are not found
```sh
git clone https://github.com/gek64/netinfo.git
cd netinfo
go build -v -trimpath -ldflags "-s -w"
```

## Test
```sh
# start netinfo service at 127.0.0.1:1996 and search ip info use net ip database
netinfo -server -address 127.0.0.1 -port 1996 -netdb

# use X-Forwarded-For to give ip that you want to search for info
curl -H "X-Forwarded-For: 8.8.8.8" http://127.0.0.1:1996

# use X-Real-Ip to give ip that you want to search for info
curl -H "X-Real-Ip: 8.8.8.8" http://127.0.0.1:1996
```

## License
- **GPL-3.0 License**
- See `LICENSE` for details
