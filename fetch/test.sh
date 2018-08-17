#!/usr/bin/env bash
scp url.go root@www.ydl.com:/root
ssh root@www.ydl.com "cd /root && go run url.go https://www.baidu.com https://www.sourcedev.cc"