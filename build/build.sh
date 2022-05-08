#!/bin/bash
BUILD_HOME=$(cd `dirname $0`; pwd)

# clean bin files.
if [ -d "$BUILD_HOME/bin" ]; then
  rm -rf $BUILD_HOME/bin
fi

# create binary directory.
mkdir $BUILD_HOME/bin

# build.
PROJECT_HOME=$(cd $BUILD_HOME; cd ..; pwd)
cd $PROJECT_HOME
go build -o $BUILD_HOME/bin/flow-center

# copy configuration files.
mkdir -p $BUILD_HOME/bin/config
cp $BUILD_HOME/config.json $BUILD_HOME/bin/config