#!/bin/zsh

TARGET=tutorial
PIPELINE_NAME=my-pipe

echo "Setting up Concourse pipeline..."
fly -t "$TARGET" login -c http://localhost:8080 -u test -p test
fly -t "$TARGET" set-pipeline --pipeline "$PIPELINE_NAME" --config pipeline.yml
fly -t "$TARGET" unpause-pipeline --pipeline "$PIPELINE_NAME"