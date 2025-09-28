alias=$1 
if [ -z "$alias" ]; then
  alias="td"
fi

echo "Buiding as $alias"
go build -o $alias
build_path=$GOPATH/bin
echo "Copying as $alias to $build_path"
mkdir -p $build_path
cp $alias $build_path/$alias
