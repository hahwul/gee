#!/bin/sh -l

out=$(echo $1 | /app/gee $2)
echo "output=$out" >> $GITHUB_ENV
