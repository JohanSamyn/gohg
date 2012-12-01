@echo off
set GOPATH=%GOPATH%;C:\DEV\go
C:\Windows\System32\cmd.exe /c start "Godoc Server http://localhost:6061" "C:\Programs\Go\bin\godoc.exe" -http=localhost:6061 -goroot="C:\Programs\Go\." -tabwidth=4
