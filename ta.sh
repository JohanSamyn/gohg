clear

shel=`ps -p $$ | tail -1 | awk '{print $NF}'`

if [[ "$shel" = "bash" ]]; then
  arg="-p "
else # probably zsh
  arg="?"
fi

echo ----------------------------------------
echo running: go test
echo ----------------------------------------
# %* is for passing in -v f.i.
go test $*
#read -p "Press [enter] to continue ..."
read $arg"Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/stats.go
echo ----------------------------------------
go run examples/stats.go
read $arg"Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/readme-test.go
echo ----------------------------------------
go run examples/readme-test.go
read $arg"Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example1.go
echo ----------------------------------------
go run examples/example1.go
read $arg"Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example2.go
echo ----------------------------------------
go run examples/example2.go
read $arg"Press [enter] to continue ..."

echo ----------------------------------------
echo running: go run examples/example3.go
echo ----------------------------------------
go run examples/example3.go
