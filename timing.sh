#!/bin/bash

function run()
{
	time ./queues-test -n 100000 $@ >/dev/null
	echo "^^^^ $@"
}

go build

run list
run slice

run -r list
run -r slice

