@echo off
::
:: script for using gocov
:: see: https://github.com/axw/gocov
::

set package=%1
set ceiling=%2
set ceilingvalue=%3
set regex=%4
goto validate

:usage
echo.
echo   Usage: cov ^<package^> -ceiling=^<nn^> ^<regex-to-filter-functions^>
echo.
echo          -ceiling=nn    Only annotate functions with coverage %% below nn.
echo.
echo   Example:  cov bitbucket.org/gohg/gohg -ceiling=50 .*addOption
echo.
echo   Run this script from the package folder itself.
echo.
goto end

:validate
if "%package%" == "" (goto usage)
if "%ceiling%" == "" (goto usage)
if "%ceilingvalue%" == "" (goto usage)
if "%regex%" == "" (goto usage)

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
gocov test %package% > covdata\coverage.json
echo === Creating summary report...
gocov report covdata\coverage.json > covdata\coverage.log
echo === Annotating source code...
gocov annotate %ceiling%=%ceilingvalue% covdata\coverage.json %regex% > covdata\coverage-annotate.log
echo === Done!

:end
