@echo off
cls

echo ----------------------------------------
echo running: go test
echo ----------------------------------------
:: %* is for passing in -v f.i.
go test %*
pause

echo ----------------------------------------
echo running: go run examples\readme-test.go
echo ----------------------------------------
go run examples\readme-test\readme-test.go
pause

echo ----------------------------------------
echo running: go run examples\example1.go
echo ----------------------------------------
go run examples\example1\example1.go
pause

echo ----------------------------------------
echo running: go run examples\example2.go
echo ----------------------------------------
go run examples\example2\example2.go
pause

echo ----------------------------------------
echo running: go run examples\example3.go
echo ----------------------------------------
go run examples\example3\example3.go

echo ----------------------------------------
echo running: go run examples\stats.go
echo ----------------------------------------
go run examples\stats\stats.go
pause
