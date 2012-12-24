@echo off
gocov test %1 > %2-coverage.json
gocov report %2-coverage.json > %2-coverage.log
gocov annotate %2-coverage.json %2.* > %2-coverage-annotate.log
