#!/bin/bash

input=$(cat input.txt)

for i in {1..1000}; do
	for j in {a..z}; do
		input=${input//${j^^}$j}
		input=${input//$j${j^^}}
	done
done

echo ${#input}
