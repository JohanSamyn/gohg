@echo off
::set GOPATH=%GOPATH%;C:\DEV\go
set GOPATH=C:\DEV\go
C:\Windows\System32\cmd.exe /c start "Godoc Server http://localhost:6161" "C:\Programs\Go\bin\godoc.exe" -goroot="C:\Programs\Go\." -http=localhost:6161 -tabwidth=4
