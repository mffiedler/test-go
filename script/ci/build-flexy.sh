#!/bin/bash

rm -fv ./flexy/flexy.test
rm -rfv ./bulid/
mkdir -p ./build/
ginkgo build flexy/
mv -v flexy/flexy.test ./build/

#TO run
# ginkgo -v -focus="\[Main\] Flexy" ./build/flexy.test
## OR,
# ginkgo -v -focus="\[Main\] Flexy" ./flexy