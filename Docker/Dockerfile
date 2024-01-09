FROM ubuntu:latest 

# add your keyname
ARG KEY_NAME <add your keyname>

# add the nickname for your validator on the network
ARG MONIKER <add your moniker>

# add the network id
ARG NETWORK_ID qom_9624-1

# arbitrary validator metadata
ARG DETAILS <details as string>

# add amount you're delegating
ARG TOKEN_DELEGATION <amount>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update &&\ 
	apt-get install -y apt-utils git git-extras software-properties-common vim \
	wget ca-certificates curl build-essential libssl-dev make openssh-server snap

RUN snap install go --classic

# builds out qomd
RUN git clone https://github.com/QOM-One/QomApp-Testnet
WORKDIR /Canto-Testnet/cmd/qomd
RUN go install -tags ledger ./... &&\
    mv $HOME/go/bin/qomd /usr/bin/

# generate keys
RUN qomd keys add $KEY_NAME
    # qomd keys add $KEY_NAME --recover &&\
    # qomd keys add $KEY_NAME --ledger

# initialize node
RUN qomd init $MONIKER --chain-id qom_9624-1

# get genesis file
RUN wget https://github.com/QOM-One/QomApp-Testnet-v2/raw/main/genesis.json -P $HOME/.qomd/config/
RUN sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0001aqom"/g' $HOME/.qomd/config/app.toml
RUN wget https://github.com/QOM-One/QomApp-Testnet-v2/raw/main/Docker/qomd.service -P $HOME/systemd/system/

# run qomd
WORKDIR /Canto-Testnet/cmd/qomd
RUN sudo systemctl daemon-reload &&\
    sudo systemctl enable qomd.service &&\
    systemctl start qomd && journalctl -u qomd -f

ENTRYPOINT ["/bin/bash"]
