#!/usr/bin/env bash

# shellcheck disable=SC2046
kill -9 $(cat pid.txt)
rm pid.txt
