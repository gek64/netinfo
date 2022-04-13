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
git clone https://github.com/gek64/gek.git
git clone https://github.com/gek64/netinfo.git
cd netinfo
go build -v -trimpath -ldflags "-s -w"
```

## QA
### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. For more information, see https://go.dev/doc/faq#virus

## License
- **GNU Lesser General Public License v2.1**
- See `LICENSE` for details
