#!/usr/bin/env bash
scp main.go root@www.ydl.com:/root
ssh root@www.ydl.com "cd /root && go run main.go"