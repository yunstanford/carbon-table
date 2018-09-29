SCRIPT_PATH=$(pwd)

go version

if [ $? != 0 ]; then
	echo "[ERROR] Go doesn't exist."
	exit 1
fi

echo "[INFO] found GOPATH: $GOPATH"

export REPO_PATH=$GOPATH/src/github.com/yunstanford

mkdir -p $REPO_PATH

echo "[INFO] set REPO_PATH: $REPO_PATH"
echo "[INFO] create dir: $REPO_PATH"

cd $REPO_PATH
echo "[INFO] cloning carbon-table from https://github.com/yunstanford/carbon-table.git"

git clone https://github.com/yunstanford/carbon-table.git

cd $GOPATH/src/github.com/yunstanford/carbon-table

echo "[INFO] compiling carbon-table..."

make build

echo "[INFO] Finished, find binary from $REPO_PATH/carbon-table/build/carbon-table"
