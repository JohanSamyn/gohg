#
# script for using gocov
# see: https://github.com/axw/gocov
#

package=$1
ceiling=$2
regex=$3

if [[ "$package" = "" -o "$ceiling" = "" -o "$regex" = "" ]]
then
  echo
  echo    Usage: cov.sh "<package> -ceiling=<nn> <regex-to-filter-functions>"
  echo
  echo           -ceiling=nn    Only annotate functions with coverage % below nn.
  echo
  echo    Example:  cov.sh bitbucket.org/gohg/gohg -ceiling=50 .*addOption
  echo
  echo    Run this script from the package folder.
  echo
else
  echo === Deleting existing logfiles...
  if [[ -d covdata ]]
  then
    if [[ -f covdata/coverage.json ]]
    then
      rm covdata/coverage.json
    fi
    if [[ -f covdata/coverage.log ]]
    then
      rm covdata/coverage.log
    fi
    if [[ -f covdata/coverage-annotate.log ]]
    then
      rm covdata/coverage-annotate.log
    fi
  else
    mkdir covdata; chmod 777 covdata
  fi
  echo === Gathering coverage info...
  gocov test $package > covdata/coverage.json
  echo === Creating summary report...
  gocov report covdata/coverage.json > covdata/coverage.log
  echo === Annotating source code...
  gocov annotate $ceiling covdata/coverage.json $regex > covdata/coverage-annotate.log
  echo "=== Done!"
fi
