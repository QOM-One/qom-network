#!/bin/bash

KEY="mykey"
CHAINID="qom_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t qom-datadir.XXXXX)

echo "create and add new keys"
./qomd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init qom with moniker=$MONIKER and chain-id=$CHAINID"
./qomd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./qomd add-genesis-account \
"$(./qomd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aqom,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./qomd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./qomd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./qomd validate-genesis --home $DATA_DIR

echo "starting qom node $i in background ..."
./qomd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started qom node"
tail -f /dev/null