FROM ubuntu

RUN DEBIAN_FRONTEND=noninteractive apt update -y \
    && DEBIAN_FRONTEND=noninteractive apt install -y curl jq