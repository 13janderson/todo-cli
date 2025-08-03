alias=$1 
if [ -z "$alias" ]; then
  alias="td"
fi

echo "Buiding as $alias"
go build -o $alias
echo "Copying as $alias to $GOPATH"
cp $alias $GOPATH/bin
