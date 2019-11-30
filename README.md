# PoC: Go Binary Diff Tool

This short script aims to use [bsdiff](http://www.daemonology.net/bsdiff/)'s
implementations in **C** and **Go** to create a diff/patch file from two related
binaries and successfully regenerate the latest from the earliest + the diff file.

## Requirements

Build the `bspatch.c` file in order to recreate the latest file.

    cd bsdiff-4.3 && make

## Usage

Create diff file:

    go run main.go earliest-file.ext latest-file.ext diff-file.ext

Recreate latest file from diff:

    bsdiff-4.3/bspatch earliest-file.ext generated-latest-file.ext diff-file.ext
