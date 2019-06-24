#!/bin/bash

cd /Users/zhuxi/zx/develop/Ravigation/src/script

ravictl add connectInfo --type agv --name agv1
ravictl add point -i 1 -x 1 -y 2 -z 3
ravictl add cmd -c '{"CgName":"movetocmd1","CmdList":[{"Type":"MoveToCmd","Index":1,"PointList":[]}]}'
ravictl add btncmdmap -o node1 -b btn1 -m movetocmd1