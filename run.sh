#!/bin/bash

if [[ $1 == "" ]]
then
  echo "usage: run.sh lnd | cln"
  exit 1
fi

if [[ $1 != 'cln' && $1 != 'lnd' ]]
then
  echo "unknown target"
  exit 1
fi

# Select target: 'cln' or 'lnd'
export TARGET=$1

# Generate up to date docker compose file from graph.yml
go run ./cmd/gencluster/... > docker-compose.yml

# Make sure that all leftovers from a previous run are removed. 
docker-compose down -v --remove-orphans

# Spin up the stack and output logs as a foreground process. Grep filters the
# output to only show the test results.
docker-compose up --build | grep testrunner # --abort-on-container-exit
