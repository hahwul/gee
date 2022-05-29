#!/bin/sh -l

out=$(echo $1 | /app/gee $2)
echo "::set-output name=output::$out"
