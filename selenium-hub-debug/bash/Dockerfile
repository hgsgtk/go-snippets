FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    bash \
    bash-completion \
    vim \
    nano \
    curl \
    wget \
    git \
    procps \
    netcat \
    telnet \
    tree \
    jq \
    socat \
    && rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-c"]

WORKDIR /workspace

CMD ["bash"]
