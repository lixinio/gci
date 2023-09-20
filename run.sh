#!/bin/bash

exec 0>&-
./dist/gci diff pkg/utils --skip-generated --skip-vendor
