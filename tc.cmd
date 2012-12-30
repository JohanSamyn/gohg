::
:: test one command
::
:: usage: testc verify
::
@echo off
if "%1" == "all" (
  go test %2 %3 %4 %5 %6 %7 %8 %9
  goto end
)
if "%1" == "version" (
  go test %2 %3 %4 %5 %6 %7 %8 %9 client.go commands.go options.go util.go setup_and_teardown_test.go %1.go %1_test.go
  goto end
)
:: Need version.go to avoid compile error, cause Version() is called in client.go.
go test %2 %3 %4 %5 %6 %7 %8 %9 client.go commands.go options.go util.go version.go setup_and_teardown_test.go %1.go %1_test.go
goto end

:end
