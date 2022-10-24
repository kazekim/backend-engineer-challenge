#!/usr/bin/env bash

cd grscankafka && mockery --name=Client --outpkg grscankafkamocks --output mocks
cd ../grgitscanner && mockery --name=Client --outpkg grgitscannermocks --output mocks
cd ../besqlx && mockery --name=Client --outpkg besqlxmocks --output mocks