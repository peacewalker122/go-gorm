#!/bin/bash
export $(cat .env | xargs)

# Parse command-line arguments
COMMAND=$1
ARG=$2
ARG2=$3

# Run the appropriate migrate command with the specified options
case $COMMAND in
  "create")
    migrate create -ext sql -dir migration -seq $ARG
    ;;
  *)
    migrate -database "${DATABASE_URL}" -path migration "$COMMAND" $ARG $ARG2
    ;;
esac
