# GENTX & HARDFORK INSTRUCTIONS

### Install & Initialize 

* Install qomd binary

* Initialize canto node directory 
```bash
qomd init <node_name> --chain-id <chain_id>
```
* Download the [genesis file](https://github.com/QOM-One/QomApp/raw/main/Mainnet/genesis.json)
```bash
wget https://github.com/QOM-One/QomApp/raw/main/Mainnet/genesis.json -b $HOME/.qomd/config
```

### Add a Genesis Account
A genesis account is required to create a GENTX

```bash
qomd add-genesis-account <address-or-key-name> aqom --chain-id <chain-id>
```
### Create & Submit a GENTX file + genesis.json
A GENTX is a genesis transaction that adds a validator node to the genesis file.
```bash
qomd gentx <key_name> <token-amount>aqom --chain-id=<chain_id> --moniker=<your_moniker> --commission-max-change-rate=0.01 --commission-max-rate=0.10 --commission-rate=0.05 --details="<details here>" --security-contact="<email>" --website="<website>"
```
* Fork [Canto](https://github.com/QOM-One/QomApp)

* Copy the contents of `${HOME}/.qomd/config/gentx/gentx-XXXXXXXX.json` to `$HOME/Canto/Mainnet/gentx/<yourvalidatorname>.json`

* Copy the genesis.json file `${HOME}/.qomd/config/genesis.json` to `$HOME/Canto/Mainnet/Genesis-Files/`

* Create a pull request to the main branch of the [repository](https://github.com/QOM-One/QomApp/Mainnet/gentx)

### Restarting Your Node

You do not need to reinitialize your Canto Node. Basically a hard fork on Cosmos is starting from block 1 with a new genesis file. All your configuration files can stay the same. Steps to ensure a safe restart

1) Backup your data directory. 
* `mkdir $HOME/canto-backup` 

* `cp $HOME/.qomd/data $HOME/canto-backup/`

2) Remove old genesis 

* `rm $HOME/.qomd/genesis.json`

3) Download new genesis

* `wget`

4) Remove old data

* `rm -rf $HOME/.qomd/data`

5) Create a new data directory

* `mkdir $HOME/.qomd/data`

If you do not reinitialize then your peer id and ip address will remain the same which will prevent you from needing to update your peers list.

7) Download the new binary
```
cd $HOME/Canto
git checkout <branch>
make install
mv $HOME/go/bin/qomd /usr/bin/
```


6) Restart your node

* `systemctl restart qomd`

## Emergency Reversion

1) Move your backup data directory into your .qomd directory 

* `mv HOME/canto-backup/data $HOME/.canto/`

2) Download the old genesis file

* `wget https://github.com/QOM-One/QomApp/raw/main/Mainnet/genesis.json -b $HOME/.qomd/config/`

3) Restart your node

* `systemctl restart qomd`