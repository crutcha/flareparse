export CGO_LDFLAGS="-L$(pwd)/target/release"
export LD_LIBRARY_PATH="$(pwd)/target/release"
go build
./flareparse
