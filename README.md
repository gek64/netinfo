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
netinfo send file -id="center" -filepath="./center.json"
## Send local network information to a file and encrypt the file
netinfo send file -id="center" -filepath="./center.json" -encryption_key="admin123"
## Loop sending local network information to a file and encrypt the file
netinfo send file -id="center" -filepath="./center.json" -encryption_key="admin123" -interval 5s

# Send local network information to s3 server
netinfo send s3 -id="center" -endpoint="https://s3.amazonaws.com" -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Send local network information to minio s3 server
netinfo send s3 -id="center" -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Send local network information to minio s3 server and encrypt the file
netinfo send s3 -id="center" -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123"
## Loop Send local network information to minio s3 server and encrypt the file
netinfo send s3 -id="center" -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123" -interval 5s

# Send local network information to webdav server
netinfo send webdav -id="center" -endpoint="http://192.168.1.2/" -filepath="/dav/center.json"
## Send local network information to webdav server and encrypt the file
netinfo send webdav -id="center" -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123"
## Loop Send local network information to webdav server and encrypt the file
netinfo send webdav -id="center" -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123" -interval 5s

# Send local network information to nconnect server
netinfo send nconnect -id="center" -endpoint="http://localhost:1996/"
## Loop Send local network information to nconnect server
netinfo send nconnect -id="center" -endpoint="http://localhost:1996/" -interval 5s
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

## License

- **GPL-3.0 License**
- See `LICENSE` for details
