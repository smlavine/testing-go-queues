#!/bin/bash

function run()
{
	time ./testing-go-queues -i 100000 $@ >/dev/null
	echo "^^^^ $@"
}

function bench()
{
	run $@ queue
	run $@ list
	run $@ slice
	echo '~~~~~~~~~~'
}

go build

bench -n 10
bench -n 10 -r

bench -n 100
bench -n 100 -r

bench -n 1000
bench -n 1000 -r

bench -n 10000
bench -n 10000 -r
