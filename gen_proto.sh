#!/bin/sh
# must execute workspace root dir

docker run --rm -v $(pwd):/work uber/prototool prototool generate
