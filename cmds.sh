#!/usr/bin/env bash

set -e

ENVSET=`cat <<EOF
windows,386,.exe
windows,amd64,.exe
darwin,386,
darwin,amd64,
linux,386,
linux,amd64,
EOF
`

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
    for ENV in ${ENVSET[@]}; do
        local OS=`echo "${ENV}" | awk -F',' '{ print $1 }'`
        local ARCH=`echo "${ENV}" | awk -F',' '{ print $2 }'`
        local EXT=`echo "${ENV}" | awk -F',' '{ print $3 }'`

        local OUTPUT_NAME="bin/pem-${OS}-${ARCH}${EXT}"

        echo "Build for ${OS}(${ARCH})"
        GOOS=${OS} GOARCH=${ARCH} \
            go build \
                -o ${OUTPUT_NAME} \
                -tags netgo -installsuffix netgo \
                --ldflags '-extldflags "-static"'

    done

    # for OS in ${OSES[@]}; do
    #     for ARCH in ${ARCHS[@]}; do
    #         OUTPUT="pem-${OS}-${ARCH}"
    #         if [ "${OS}" = "windows" ]; then
    #             OUTPUT="bin/${OUTPUT}.exe"
    #         fi

    #         GOOS=${OS} GOARCH=${ARCH} \
    #             go build \
    #                 -o ${OUTPUT} \
    #                 -tags netgo -installsuffix netgo \
    #                 --ldflags '-extldflags "-static"'

    #     done
    # done
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
