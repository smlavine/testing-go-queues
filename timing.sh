#!/bin/bash

go build

echo "slice"
time ./queues-test -n 100000 slice >/dev/null

echo "list"
time ./queues-test -n 100000 list >/dev/null

echo "-r slice"
time ./queues-test -n 100000 -r slice >/dev/null

echo "-r list"
time ./queues-test -n 100000 -r list >/dev/null
