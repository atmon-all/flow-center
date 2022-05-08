#!/bin/bash
BUILD_HOME=$(cd `dirname $0`; pwd)

# clean all pb code files.
GRPC_HOME=$(cd $BUILD_HOME; cd ..; pwd)/grpc
rm -rf $GRPC_HOME/*.go

# generate code by pb file in grpc directory.
cd "$GRPC_HOME"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

# generate code by pb file in sub directory.
SUB_FILES=`ls $GRPC_HOME`
for CURRENT in $SUB_FILES
do
  # build path.
  SUB_DIR=$GRPC_HOME/$CURRENT

  if [ ! -d  $SUB_DIR ]; then
    continue
  fi

  # clean all pb code files.
  rm -rf $SUB_DIR/*.go

  # generate code by pb file in grpc directory.
  cd "$SUB_DIR"
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
done
