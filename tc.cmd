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
go test %2 %3 %4 %5 %6 %7 %8 %9 client.go commands.go options.go util.go setup_and_teardown_test.go %1.go %1_test.go
:end
