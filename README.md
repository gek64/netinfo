# Netinfo

## Features

- Display local IP information
- Setup as a server to tell client its ip information

## Usage

```
Usage:                                   
         netinfo {Command} [Option]      
                                         
        Command:                         
         -client           : start client
         -server           : start server
         -h                : show help   
         -v                : show version
                                         
        Option:                          
         -interval       <IP>     : set client interval
         -description    <Port>   : set client description

        Example:
         1) netinfo -show
         2) netinfo -server localhost:1996
         3) netinfo -client http://localhost:1996/record -interval 6h -description home_opnsense
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
curl -Lo /etc/systemd/system/netinfo.service https://github.com/gek64/netinfo/raw/main/configs/netinfo.service
systemctl enable netinfo && systemctl start netinfo
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/netinfo https://github.com/gek64/netinfo/raw/main/configs/netinfo
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
# start netinfo server at 127.0.0.1:1996
netinfo -server localhost:1996

# start netinfo client
netinfo -client http://localhost:1996/record -interval 1h -description home_opnsense

# check info
curl http://localhost:1996/record/all
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
