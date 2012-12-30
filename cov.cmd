@echo off
::
:: usage: cov <package> <logfile-prefix>
:: examples:
::		cov . gohg
::		cov gohg gohg
::
echo === Deleting existing logfiles...
if exist data\. (
if exist data\%2-coverage.json (del data\%2-coverage.json)
if exist data\%2-coverage.log (del data\%2-coverage.log)
if exist data\%2-coverage-annotate.log (del data\%2-coverage-annotate.log)
goto new
)
mkdir data
:new
echo === Gathering coverage info...
gocov test %1 > data\%2-coverage.json
echo === Creating logfile...
gocov report data\%2-coverage.json > data\%2-coverage.log
echo === Annotating functions...
gocov annotate data\%2-coverage.json %2.* > data\%2-coverage-annotate.log
echo === Done!
