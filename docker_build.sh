#!/bin/bash

version=$1

docker build -t jw-sys:$version .

# shellcheck disable=SC2086
docker tag jw-sys:$version www.jwdouble.top:10443/k8s/jw-sys:$version

# shellcheck disable=SC2086
docker push www.jwdouble.top:10443/k8s/jw-sys:$version