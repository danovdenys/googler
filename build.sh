#!/bin/bash
BUILD_PATH=/usr/bin
go build .
sudo mv ./googler $BUILD_PATH/googler

echo Built successfully into $BUILD_PATH/googler