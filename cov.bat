@echo off
::
:: script for using gocov
:: see: https://github.com/axw/gocov
::

if "%1" == "" (
	goto usage
)

echo === Deleting existing logfiles...
if exist covdata\. (
if exist covdata\coverage.json (del covdata\coverage.json)
if exist covdata\coverage.log (del covdata\coverage.log)
if exist covdata\coverage-annotate.log (del covdata\coverage-annotate.log)
goto docov
)
mkdir covdata

:docov
echo === Gathering coverage info...
gocov test %1 > covdata\coverage.json
echo === Creating summary report...
gocov report covdata\coverage.json > covdata\coverage.log
echo === Annotating functions...
gocov annotate covdata\coverage.json .* > covdata\coverage-annotate.log
echo === Done!
goto end

:usage
echo.
echo usage: cov ^<package^>
echo
echo Run this script from the package folder.
echo.

:end
