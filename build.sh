#!/bin/bash

# 构建前端
function build_web {
  cd web
  pnpm install
  pnpm run build
  cd ..
}

# 构建 Go 程序
function build_go {
  # 确保目录存在
  mkdir -p static
  # 清空目标目录
  rm -rf static/*
  # 复制到 Go 的静态资源目录
  cp -r web/dist/* static/
  go build -o localserver
}

# 主执行逻辑
echo "开始构建..."
build_web
echo "web构建完成！"
build_go
echo "go构建完成！"

