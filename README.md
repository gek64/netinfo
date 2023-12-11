# Netinfo

## Features

- Send network information to a remote server or file
- Gateway for receiving network information

## Example

```sh
# Show local network information
netinfo list

# Start nconnect server
netinfo receive -listen localhost:1996

# Send local network information to a file
netinfo send file -filepath="./center.json"
## Send local network information to a file and encrypt the file
netinfo send file -filepath="./center.json" -encryption_key="admin123"
## Loop sending local network information to a file and encrypt the file
netinfo send file -filepath="./center.json" -encryption_key="admin123" -interval="5m"

# Send local network information to s3 server
netinfo send s3 -endpoint="https://s3.amazonaws.com" -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Send local network information to minio s3 server
netinfo send s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Send local network information to minio s3 server and encrypt the file
netinfo send s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123"
## Loop Send local network information to minio s3 server and encrypt the file
netinfo send s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123" -interval="5m"

# Send local network information to webdav server
netinfo send webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json"
## Send local network information to webdav server and encrypt the file
netinfo send webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123"
## Loop Send local network information to webdav server and encrypt the file
netinfo send webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123" -interval="5m"

# Send local network information to nconnect server
netinfo send nconnect -id="center" -endpoint="http://localhost:1996/"
## Loop Send local network information to nconnect server
netinfo send nconnect -id="center" -endpoint="http://localhost:1996/" -interval="5m"
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
curl -Lo /etc/systemd/system/xxx.service https://github.com/gek64/netinfo/raw/main/configs/systemd/xxx.service
systemctl enable xxx.service && systemctl restart xxx.service && systemctl status xxx.service
curl -Lo /etc/systemd/system/xxx.timer https://github.com/gek64/netinfo/raw/main/configs/systemd/xxx.timer
systemctl enable xxx.timer && systemctl restart xxx.timer && systemctl status xxx.timer
```

### Alpine Linux(openrc)

```sh
curl -Lo /etc/init.d/xxx https://github.com/gek64/netinfo/raw/main/configs/openrc/xxx
chmod +x /etc/init.d/xxx
rc-update add xxx && rc-service xxx restart && rc-service xxx status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/xxx https://github.com/gek64/netinfo/raw/main/configs/rc.d/xxx
chmod +x /usr/local/etc/rc.d/xxx
service xxx enable && service xxx restart && service xxx status
```

### OpenWRT(init.d)

```sh
curl -Lo /etc/init.d/xxx https://github.com/gek64/netinfo/raw/main/configs/init.d/xxx
chmod +x /etc/init.d/xxx
service xxx enable && service xxx restart && service xxx status
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/netinfo.git
cd netinfo
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

### For mipsle router

```sh
git clone https://github.com/gek64/netinfo.git
cd netinfo
export GOOS=linux
export GOARCH=mipsle
export GOMIPS=softfloat
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
