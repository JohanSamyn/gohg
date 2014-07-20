clear

shel=`ps -p $$ | tail -1 | awk '{print $NF}'`

paause() {
  if [ "$shel" = "bash" ]; then
    read -p "Press [enter] to continue ..."
  else # probably zsh
    read \?"Press [enter] to continue ..."
  fi
}

echo ----------------------------------------
echo running: go test
echo ----------------------------------------
# %* is for passing in -v f.i.
go test $*
paause

echo ----------------------------------------
echo running: go run examples/readme-test.go
echo ----------------------------------------
go run examples/readme-test/readme-test.go
paause
echo ----------------------------------------
echo running: go run examples/example1.go
echo ----------------------------------------
go run examples/example1/example1.go
paause

echo ----------------------------------------
echo running: go run examples/example2.go
echo ----------------------------------------
go run examples/example2/example2.go
paause

echo ----------------------------------------
echo running: go run examples/example3.go
echo ----------------------------------------
go run examples/example3/example3.go

echo ----------------------------------------
echo running: go run examples/stats.go
echo ----------------------------------------
go run examples/stats/stats.go
paause
