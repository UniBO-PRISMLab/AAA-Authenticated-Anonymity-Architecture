FROM node:22-alpine

RUN apk add --no-cache curl

WORKDIR /usr/src/app

COPY . .

RUN npm ci
RUN npx hardhat clean
RUN npx hardhat compile

RUN chmod +x scripts/start.sh

EXPOSE 8545
ENTRYPOINT ["./scripts/start.sh"]