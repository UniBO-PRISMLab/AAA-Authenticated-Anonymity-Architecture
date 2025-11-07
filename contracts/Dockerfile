# FROM node:22-alpine

# RUN apk add --no-cache curl

# WORKDIR /usr/src/app

# COPY . .

# RUN npm ci
# RUN npx hardhat clean
# RUN npx hardhat compile

# RUN chmod +x scripts/start.sh

# EXPOSE 8545
# ENTRYPOINT ["./scripts/start.sh"]

FROM ghcr.io/foundry-rs/foundry:latest
USER root
RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY . .

RUN npm ci && npx hardhat clean && npx hardhat compile

COPY scripts/start.sh ./scripts/start.sh
RUN chmod +x ./scripts/start.sh

EXPOSE 8545

ENTRYPOINT ["./scripts/start.sh"]
