#!/bin/bash

input=$(cat input.txt)
size=${#input}

# Just do 26 different rounds of part 1, ezpz
for chr in {a..z}; do
	tmpInput=${input//$chr}
	tmpInput=${tmpInput//${chr^^}}
	for i in {1..1000}; do
		for j in {a..z}; do
			tmpInput=${tmpInput//${j^^}$j}
			tmpInput=${tmpInput//$j${j^^}}
		done
	done
	if [ ${#tmpInput} -lt $size ]; then
		$size=${#tmpInput}
	fi
done

echo $size
