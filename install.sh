echo "Buiding as td"
go build -o td
echo "Copying as td to $GOPATH"
cp td $GOPATH/bin
