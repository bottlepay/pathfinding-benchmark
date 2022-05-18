#!/bin/bash

set -o pipefail

if [ -n "$WAIT_FOR_LND" ]; then
    echo waiting for other lnd $WAIT_FOR_LND

    while true; do
        lncli --rpcserver=node-$WAIT_FOR_LND --tlscertpath=/cfg/$WAIT_FOR_LND/tls.cert --macaroonpath=/cfg/$WAIT_FOR_LND/admin.macaroon --network=regtest getinfo | jq -e '.synced_to_chain == true'
        if [ $? -eq 0 ]; then
            break
        fi
        sleep 1
    done
fi

lnd $@