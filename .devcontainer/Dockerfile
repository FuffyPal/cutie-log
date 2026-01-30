FROM golang:1.21-bookworm

# SQLite ve gopsutil için gerekli araçlar
RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    sqlite3 \
    procps \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspaces/cutie-log