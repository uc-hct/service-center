#!/bin/sh
set -e
echo $GOPATH
echo $(pwd)
export COVERAGE_PATH=$(pwd)
cd $COVERAGE_PATH
#$1
for d in $(go list ./... | grep server); do
    cd $GOPATH/src/$d
    if [ $(ls | grep _test.go | wc -l) -gt 0 ]; then
        go test -cover -covermode atomic -coverprofile coverage.out  
        if [ -f coverage.out ]; then
            sed '1d;$d' coverage.out >> $GOPATH/src/github.com/ServiceComb/service-center/coverage.txt
        fi
    fi
done
