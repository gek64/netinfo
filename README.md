# Netinfo

## Features

- Send network information to a remote server or file
- Gateway for receiving network information

## Usage

```
Usage:
netinfo [OPTION]...

Startup:
  -l
  -h
  -v
-------------------------------------------
Send Mode:
  -id <string>
  -interval <time>
  -sendmode <mode>
    mode value: file, s3, webdav, netinfo

  Send Mode file:
    -file <file>

  Send Mode s3:
    -endpoint <url>
    -access_key_id <string>
    -secret_access_key <string>
    -object_path <bucket/object>
    -skip-certificate-verify

  Send Mode webdav:
    -endpoint <url>
    -webdav_user <string>
    -webdav_password <string>
    -skip-certificate-verify

  Send Mode netinfo:
    -endpoint <url>
    -netinfo_user <string>
    -netinfo_password <string>
    -skip-certificate-verify
-------------------------------------------
Receive Mode:
  -listen <address:port>
  -receivemode <mode>
    mode value: netinfo
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/netinfo https://github.com/gek64/netinfo/releases/latest/download/netinfo-linux-386
chmod +x /usr/local/bin/netinfo

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/netinfo https://github.com/gek64/netinfo/releases/latest/download/netinfo-freebsd-amd64
chmod +x /usr/local/bin/netinfo
```

## Install Service

### Linux(systemd)

```sh
curl -Lo /etc/systemd/system/netinfo.service https://github.com/gek64/netinfo/raw/main/configs/netinfo.service
systemctl enable netinfo && systemctl restart netinfo && systemctl status netinfo
```

### Linux(openrc)

```sh
curl -Lo /etc/init.d/netinfo https://github.com/gek64/netinfo/raw/main/configs/netinfo.openrc
chmod +x /etc/init.d/netinfo
rc-update add netinfo && rc-service netinfo restart && rc-service netinfo status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/netinfo https://github.com/gek64/netinfo/raw/main/configs/netinfo.rcd
chmod +x /usr/local/etc/rc.d/netinfo
service netinfo enable && service netinfo restart && service netinfo status
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
# start netinfo server at 127.0.0.1:1996
netinfo -server localhost:1996

# start netinfo client
netinfo -client http://localhost:1996 -interval 15m -id center

# check info
curl http://localhost:1996/all
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
