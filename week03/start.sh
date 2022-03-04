#!/bin/bash
docker run -dit \
 --name HttpServer -h HttpServer \
 -p 8080:8080 \
 -v /data/logs/:/data/logs \
 --restart=always  httpserver:v0.0.1
