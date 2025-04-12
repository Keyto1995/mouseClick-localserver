# 构建前端
function Build-Web {
  Set-Location web
  pnpm install
  pnpm run build

  Set-Location ..
  # 确保目录存在
  New-Item -Path "static" -ItemType Directory -Force | Out-Null
  # 清空目标目录
  Remove-Item -Path "static/*" -Recurse -Force -ErrorAction Ignore
  # 复制到 Go 的静态资源目录
  Copy-Item -Path "web/dist/*" -Destination "static/" -Recurse -Force
}

# 构建 Go 程序
function Build-Go {
  Build-Web
  go build -o localserver.exe
}

# 主执行逻辑
if ($MyInvocation.InvocationName -ne '.') {
  Build-Go
}
