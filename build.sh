#!/bin/bash

Toolbox=("go")
OS=("windows" "linux" "darwin")
Architecture=("amd64" "386" "arm" "arm64")
Proxy=""
ProgramName=""

BaseFolder=$(basename "$PWD")

# 检查工具链
function CheckToolbox() {
  pass=true
  for i in "${Toolbox[@]}"; do
    if [ ! "$(command -v "$i")" ]; then
      echo -e "\033[31m找不到$i\033[0m"
      pass=false
    fi
  done

  ## 返回值判断
  if $pass; then
    return 0
  else
    echo -e "\033[31m未能满足运行条件,请安装以上依赖后再启动程序\033[0m"
    return 1
  fi
}

function Init() {
  # 检查依赖库是否安装
  if ! CheckToolbox; then
    exit 1
  fi
  # 获取编译文件名
  GetProgramName
  # 获取Args
  while [[ $# -gt 0 ]]; do
    case "$1" in
    "-h" | "--help")
      Show_help
      break
      ;;
    "-n" | "--name")
      ProgramName=$2
      shift
      ;;
    "-p" | "--proxy")
      Proxy=$2
      shift
      ;;
    esac
    shift
  done

}

# 显示帮助信息
function Show_help() {
  cat <<-EOF
    Usage:
    $0 [Arguments]

    Arguments:
      -p | --proxy proxy_url  : use proxy
      -n | --name  staticName : set static file name
      -h | --help             : show help

    Example:
      1) $0 -p http://127.0.0.1:1080
      2) $0 -n myapp
      3) $0 -h
EOF
  exit 0
}

# 编译
function Build() {
  for i in "${OS[@]}"; do
    for j in "${Architecture[@]}"; do
      export "GOOS=$i"
      export "GOARCH=$j"

      if [ "$i" == "windows" ]; then
        staticFile="bin/$ProgramName-$i-$j.exe"
      else
        staticFile="bin/$ProgramName-$i-$j"
      fi

      if [ "$Proxy" != "" ]; then
        export http_proxy=$Proxy
        export https_proxy=$Proxy
      fi
      go build -x -trimpath -ldflags "-s -w" -o "$staticFile"
    done
  done
}

# 获取程序名称
function GetProgramName() {
  if [ "$ProgramName" == "" ]; then
    ProgramName="$BaseFolder"
  fi
}

Init "$@"
Build
