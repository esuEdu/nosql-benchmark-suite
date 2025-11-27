#!/usr/bin/env bash
set -e

echo "=== Atualizando pacotes ==="
sudo apt update -y
sudo apt upgrade -y

echo "=== Instalando dependências base ==="
sudo apt install -y git curl unzip wget default-jre

echo "=== Instalando Go ==="
if ! command -v go &> /dev/null; then
    wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
    source ~/.bashrc
fi

echo "=== Instalando Docker ==="
if ! command -v docker &> /dev/null; then
    sudo apt install -y ca-certificates curl gnupg lsb-release
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg \
      | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

    echo \
      "deb [arch=$(dpkg --print-architecture) \
      signed-by=/etc/apt/keyrings/docker.gpg] \
      https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
      | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt update
    sudo apt install -y docker-ce docker-ce-cli containerd.io
    sudo usermod -aG docker ubuntu
fi

echo "=== Subindo bancos NoSQL via docker-compose ==="
docker compose -f docker-compose.yml up -d

echo "=== Instalando AWS CLI ==="
if ! command -v aws &> /dev/null; then
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    sudo ./aws/install
fi

echo "=== Setup concluído ==="
echo "Reinicie seu terminal para aplicar PATH do Go e Docker group."
