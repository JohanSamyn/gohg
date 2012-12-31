
@echo off

set newgopath=%GOPATH%
set mod=0

if "%newgopath%" == "" (
  set newgopath=C:\Programs\GoLib;C:\DEV\go
  set mod=1
) else (
  echo.%newgopath% | findstr /C:"C:\Programs\GoLib" 1>nul
  if errorlevel 1 (
    set newgopath=%newgopath%;C:\Programs\GoLib
    set mod=1
  )
  echo.%newgopath% | findstr /C:"C:\DEV\go" 1>nul
  if errorlevel 1 (
    set newgopath=%newgopath%;C:\DEV\go
    set mod=1
  )
)

if %mod% == 1 (
  set GOPATH=%newgopath%
)
