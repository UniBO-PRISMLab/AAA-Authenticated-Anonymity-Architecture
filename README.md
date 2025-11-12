# Authenticated-Anonymity-Architecture (AAA)

[![Contracts Tests](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/hardhat-test.yaml/badge.svg)](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/hardhat-test.yaml)
[![Push nip-backend image to ghcr.io](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/nip-backend-ghcr.yaml/badge.svg)](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/nip-backend-ghcr.yaml)
[![acm-link](https://img.shields.io/badge/doi/10.1145/3677525.3678676-black.svg?style=plain&logo=ACM&logoColor=white)](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://dl.acm.org/doi/10.1145/3677525.3678676&ved=2ahUKEwi19Namm9aQAxVbgP0HHbK4BGkQFnoECBoQAQ&usg=AOvVaw3jxZaNiXsVdYoM-rXoTAVL)

The Authenticated Anonymity Architecture (AAA) is a blockchain-based solution designed to provide robust and ethical authenticated anonymous identities, enabling users to maintain anonymity while allowing for deanonymization in cases of criminal activity.

## Glossary

- PID (Public Identity Data): anonymous token that identifies the user. The real identity of the user is carried by the PID.
- NIP (National Identity Provider): national institution that releases a PID after verification of the real identity of a person.
- UIP (Union of Identity Provide): the network of official national identity NIPs of each contry.
- SID (Secret Identity Data): hash of the concatenation of the hashes of the 24-words hash.
- PAC (Public Authentication Code): one-time code used to authenticate the user.
- SAC (Secret Authentication Code): one-time code used to authenticate an anonymous user.

## How To Run

```sh
docker compose up --no-attach contracts
```

`--no-attach` option will suppress contracts logs which can be noisy.
