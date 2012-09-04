@echo off

cls
cd C:\DEV\go\src\gohg
go build gohg.go
choice /c yn /m "Continue?"
if errorlevel == 255 goto end
if errorlevel == 2 goto end
if errorlevel == 1 goto test
if errorlevel == 0 goto end

:test
cd C:\DEV\go\src\gohg\test
go build test.go
choice /c yn /m "Continue?"
cls & test

:end
