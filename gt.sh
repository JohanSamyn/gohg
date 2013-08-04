#
# script for running go test
#

if [ "$1" == "" ]
then
  echo
  echo "  usage:    'gt.sh <command> [-v]' or 'gt.sh all [-v]'"
  echo
  echo "  examples:    gt.sh summary"
  echo "               gt.sh all"
  echo
else
  if [ "$1" == "all" ]
  then
    go test $2 $3 $4 $5 $6 $7 $8 $9
  else
    if [ "$1" == "client" ]
    then
      go test %2 %3 %4 %5 %6 %7 %8 %9 options.go version.go identify.go %1.go %1_test.go
    else
      if [ "$1" == "version" ]
      then
        go test $2 $3 $4 $5 $6 $7 $8 $9 client.go options.go add.go setup_and_teardown_test.go $1.go $1_test.go
      else
        if [ "$1" == "init" ]
        then
          go test $2 $3 $4 $5 $6 $7 $8 $9 client.go options.go add.go version.go setup_and_teardown_test.go $1.go $1_test.go
        else
          if [ "$1" == "commit" ]
          then
            # uses tip in test
            go test $2 $3 $4 $5 $6 $7 $8 $9 client.go options.go add.go version.go setup_and_teardown_test.go tip.go $1.go $1_test.go
          else
            # Need version.go to avoid compile error, cause Version() is called in client.go.
            go test $2 $3 $4 $5 $6 $7 $8 $9 client.go options.go add.go version.go setup_and_teardown_test.go init.go $1.go $1_test.go
          fi
        fi
      fi
    fi
  fi
fi
