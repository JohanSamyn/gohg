@echo off
cls

echo ----------------------------------------
echo running: go test
echo ----------------------------------------
:: %* is for passing in -v f.i.
go test %*
pause

echo ----------------------------------------
echo running: go run examples\stats.go
echo ----------------------------------------
go run examples\stats.go
pause

echo ----------------------------------------
echo running: go run examples\readme-test.go
echo ----------------------------------------
go run examples\readme-test.go
pause

echo ----------------------------------------
echo running: go run examples\example1.go
echo ----------------------------------------
go run examples\example1.go
