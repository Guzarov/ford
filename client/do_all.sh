#!/bin/bash

docker build . -t ford-web
echo "Build complete!"

docker run -p 8080:80 ford-web