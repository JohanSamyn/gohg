clear

echo ----------------------------------------
echo running: go test
echo ----------------------------------------
# %* is for passing in -v f.i.
go test $*
read -p "Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/stats.go
echo ----------------------------------------
go run examples/stats.go
read -p "Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/readme-test.go
echo ----------------------------------------
go run examples/readme-test.go
read -p "Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example1.go
echo ----------------------------------------
go run examples/example1.go
read -p "Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example2.go
echo ----------------------------------------
go run examples/example2.go
read -p "Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example3.go
echo ----------------------------------------
go run examples/example3.go
