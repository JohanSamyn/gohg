@echo off
::
:: usage: cov <package> <logfile-prefix>
:: examples:
::		cov . gohg
::		cov gohg gohg
::
echo === Deleting existing logfiles...
if exist %2-coverage.json (del %2-coverage.json)
if exist %2-coverage.log (del %2-coverage.log)
if exist %2-coverage-annotate.log (del %2-coverage-annotate.log)
echo === Gathering coverage info...
gocov test %1 > %2-coverage.json
echo === Creating logfile...
gocov report %2-coverage.json > %2-coverage.log
echo === Annotating functions...
gocov annotate %2-coverage.json %2.* > %2-coverage-annotate.log
echo === Done!
