#!/bin/sh -l

out=$(echo $1 | /usr/bin/gee $2)
echo "::set-output name=output::$out"