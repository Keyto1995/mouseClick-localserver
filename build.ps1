# 构建前端
function Build-Web {
  Write-Output "开始构建web..."

  Set-Location web
  pnpm install
  pnpm run build
  Set-Location ..

  Write-Output "web构建完成！"
}

# 构建 Go 程序
function Build-Go {
  Write-Output "开始构建Go程序..."

  # 确保目录存在
  New-Item -Path "static" -ItemType Directory -Force | Out-Null
  # 清空目标目录
  Remove-Item -Path "static/*" -Recurse -Force -ErrorAction Ignore
  # 复制到 Go 的静态资源目录
  Copy-Item -Path "web/dist/*" -Destination "static/" -Recurse -Force
  go build -o localserver.exe

  Write-Output "Go构建完成！"
}

# 主执行逻辑
Write-Output "开始构建..."
Build-Web
Build-Go
