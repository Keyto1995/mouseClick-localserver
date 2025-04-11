# 构建前端
function Build-Web {
  Set-Location web
  pnpm install
  pnpm run build
  Set-Location ..
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
