::
:: script for running go test
::
@echo off
if "%1" == "" (
  goto usage
)
if "%1" == "all" (
  go test %2 %3 %4 %5 %6 %7 %8 %9
  goto end
)
if "%1" == "version" (
  go test %2 %3 %4 %5 %6 %7 %8 %9 client.go options.go add.go setup_and_teardown_test.go %1.go %1_test.go
  goto end
)
if "%1" == "init" (
  go test %2 %3 %4 %5 %6 %7 %8 %9 client.go options.go add.go version.go setup_and_teardown_test.go %1.go %1_test.go
  goto end
)
if "%1" == "commit" (
  go test %2 %3 %4 %5 %6 %7 %8 %9 client.go options.go add.go version.go setup_and_teardown_test.go tip.go %1.go %1_test.go
  goto end
)
:: Need version.go to avoid compile error, cause Version() is called in client.go.
:: Need init.go for tests
go test %2 %3 %4 %5 %6 %7 %8 %9 client.go options.go init.go add.go version.go setup_and_teardown_test.go %1.go %1_test.go
goto end

:usage
echo.
echo   usage:    'gt ^<command^> [-v]' or 'gt all [-v]'
echo.
echo   examples: gt verify
echo             gt all -v
echo.
echo   Run this script from the package folder itself.
echo.

:end
