#!/usr/bin/env bash

set -e

create_mock() {
    rm -rf mock
    mkdir mock

    FILES=$(find . -type f -name \*.go | grep -v "./vendor/")
    for F in ${FILES[@]}; do
        if grep -e 'type.*interface' ${F} >/dev/null 2>&1; then
            F_UNPREFIX=`echo "${F}" | sed 's/^\.\///'`
            echo "Create mock of ${F_UNPREFIX}"
            mockgen -source ${F_UNPREFIX} -destination mock/${F_UNPREFIX}
        fi
    done
}

test_app() {
    go test ./... -v -count=1
}

build_app() {
    go build
}

case $1 in
    mock)
        create_mock
        ;;
    test)
        test_app
        ;;
    build)
        build_app
        ;;
    suite)
        create_mock
        test_app
        build_app
        ;;
    *)
        echo "$0 <command>"
        echo ""
        echo "<command>"
        echo "  mock  ... Create mock"
        echo "  test  ... Test app"
        echo "  build ... Build app"
        echo "  suite ... Run mock, test and build"
esac
