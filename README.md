# Netinfo
Netinfo
- Display local IP information
- Setup as a server to tell the client its ip information 
- Written in golang 

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
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/netinfo.git

cd netinfo

go build -v -trimpath -ldflags "-s -w"
```

## QA

### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This report occurred after `Windows 10 21h2`. This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. For more information, see https://go.dev/doc/faq#virus

### Q: Why should I clone `https://github.com/gek64/gek.git` before building
A: I don’t want the project to depend on a certain cloud service provider, and this is also a good way to avoid network problems.


## License

**GNU Lesser General Public License v2.1**

See `LICENSE` for details
