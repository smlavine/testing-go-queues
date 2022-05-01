#!/bin/bash

function run()
{
	time ./testing-go-queues $@ >/dev/null
	echo "^^^^ $@"
}

function bench()
{
	run $@ list
	run $@ queue
	run $@ slice
	echo '~~~~~~~~~~'
}

go build

bench -i 1000 -n 10
bench -i 1000 -n 10 -r

bench -i 1000 -n 100
bench -i 1000 -n 100 -r

bench -i 1000 -n 1000
bench -i 1000 -n 1000 -r

bench -i 1000 -n 10000
bench -i 1000 -n 10000 -r


bench -n 1000
bench -n 1000 -r

bench -n 10000
bench -n 10000 -r

bench -n 100000
bench -n 100000 -r

bench -n 1000000
bench -n 1000000 -r

bench -n 10000000
bench -n 10000000 -r

bench -n 100000000
bench -n 100000000 -r
