# AAA - Functional Requirements

`version 0.0.1`

- Authors [Michele Dinelli](mailto:m.dinelli@unibo.it), [Luca Sciullo](mailto:luca.sciullo2@unibo.it), [Lorenzo Gigli](mailto:lorenzo.gigli@unibo.it)
- Date `21/05/2025`

---

The Authenticated Anonymity Architecture (AAA) is a blockchain-based solution designed to provide robust and ethical authenticated anonymous identities, enabling users to maintain anonymity while allowing for deanonymization in cases of criminal activity [[1]](#1).

## Glossary

- PID (Public Identity Data): anonymous token that identifies the user. The real identity of the user is carried by the PID.
- NIP (National Identity Provider): national institution that releases a PID after verification of the real identity of a person.
- SID (Secret Identity Data): hash of the concatenation of the hashes of the 24-words hash.
- PK (Public Key): a cryptographic key used to encrypt/sign payloads in a non repudiable way.
- SK (Secret Key): a cryptographic key used to decrypt/verify payloads.
- SYMK (Symmetric Key): a cryptographic key generally used for symmetric cryptographic communication protocols (e.g., AES, DES)
- PAC (Public Authentication Code): one-time code used to authenticate the user.
- SAC (Secret Authentication Code): one-time code used to authenticate an anonymous user.

## 1. Identity Management

### 1.1 User Registration and Public Identity (PID) Issuance:

- The system SHALL allow a user to submit a request to their National Identity Provider (NIP) with personal data (e.g., ID card, passport) and a Public Key (PK).
- The NIP SHALL be able to verify the user's real identity.
- The NIP SHALL issue a Public Identity Data (PID), an anonymous token identifying the user within the system without explicitly sharing personal information.
- The NIP SHALL save the PID on its local database and send it to the user.
- The NIP SHALL be the unique entity capable of connecting a PID to the real user information.

### 1.2 Seed Phrase Generation and Secret Identity Data (SID) Creation:

- The system SHALL allow a user to send their PID and a PK to a smart contract to request a seed phrase.
- The smart contract SHALL initiate a protocol to generate N random words for the seed phrase.
- The smart contract SHALL encrypt and save the generated words to the blockchain, allowing the user to retrieve them confidentially via their Secret Key (SK).
- The smart contract SHALL generate a Secret Identity Data (SID) token, which is the hash of the concatenation of the hashes of the N generated words.
- The smart contract SHALL save the SID on the blockchain, encrypted with the user's provided Public Key (PK), for confidential recovery by the user.
- The protocol SHALL duplicate each generated word multiple times, saving the record `PID: ENC(PID, wordnumber, word, PKi)` (association between PID, word, and order, encrypted with the public key of a randomly selected UIP node) to the blockchain for redundancy.
- The Union of Identity Providers (UIP) SHALL be able to reconstruct the seed phrase of a user.
- Consensus among UIP nodes SHALL be required to decrypt words and reconstruct the seed phrase.
- The smart contract SHALL create a Symmetric Key (SYMK) using the concatenation of the hash of each word.
- The smart contract SHALL store the record `SID: ENC(PID, SYMK), PK` on the blockchain, representing the PID-SID association encrypted with the symmetric key and public key.

## 2. Authentication Services

### 2.1 Public Authentication Code (PAC) Issuance:

- The system SHALL allow a user to request a PAC from their NIP for services requiring an authenticated public identity.
- The user SHALL send a message to the NIP containing `SIGN(PID, SK)` (PID signed with the private key used to obtain the PID).
- The NIP SHALL verify that the user is the true key holder.
- The NIP SHALL return the PAC to the user and save it in its local repository.
- The system SHALL allow public services to query the NIP to verify if a PAC is associated with an authenticated user.

### 2.2 Secret Authentication Code (SAC) Issuance and Anonymous Identity Authorization:

- The system SHALL allow a user to request a SAC from the UIP for services requiring an authenticated anonymous identity.
- The user SHALL query the NIP by sending a message containing `SID: ENC(SID, SK)` (SID signed with the private key associated with the PK used for SID storage).
- The NIP SHALL retrieve the SID record from the blockchain and verify the signature against the PK saved in the record, certifying user ownership of the SID.
- The NIP SHALL save the mapping `SAC: SID` in its local repository.
- The NIP SHALL save the SAC on the blockchain.
- The smart contract SHALL check the SAC existence on the blockchain.
- The smart contract SHALL save a `SAC: PK` record on the blockchain to store the association between the SAC and the PK of the anonymous account.
- The system SHALL allow a user to create multiple PK-SK pairs (anonymous identities) using their seed phrase.
- The user SHALL send the PK of the anonymous account they intend to use, along with their SAC, to the smart contract for authorization.

### 2.3 Anonymous Service Login:

- The system SHALL allow a user to log in to an anonymous service by providing a record containing `SAC, PK, SIGN(SAC, SK)` (SAC, PK of the anonymous account, and SAC signed with the SK associated with the PK).
- The anonymous service SHALL retrieve the SAC record from the blockchain.
- The anonymous service SHALL verify the signature to confirm the user owns the PK-SK pair.
- The anonymous service SHALL verify the association between the SAC and the anonymous account.
- The anonymous service SHALL provide anonymous access to its features if all checks match.

## 3. Security and Transparency

- The system SHALL securely and robustly connect public and anonymous identities for all legitimate uses.
- The system SHALL provide a simple and reliable way to deanonymize users committing crimes.
- The system SHALL ensure that deanonymization can only be performed by properly authorized actors.
- The system SHALL ensure that deanonymization is conducted in a traceable, transparent, dated, authored, and non-repudiable manner to prevent or limit abuses.
- The system SHALL allow deanonymization only if a necessary consensus is reached among participating organizations, preventing unethical deanonymizations by "evil" actors.
- The system SHALL store private data on the blockchain in an encrypted format to ensure readability only by specific actors.
- The system SHALL use cryptographic systems to ensure data in the secret layer is retrieved and managed only by the owner.
- The system SHALL ensure that the seed phrase generation process is completely verifiable and reproducible by all actors through the initialization of the Pseudo Random Number Generator (PRNG) with the user's PID (\*).

## 4. System Robustness and Fault Tolerance

- The architecture SHALL be resilient and fault-tolerant, even with a large number of managed identities.
- The system SHALL be designed to mitigate "evil attacks" (coalitions of participating evil actors trying to force unethical deanonymization).
- The system SHALL be designed to mitigate "node faults" (coalitions of oppressive countries trying to block ethical deanonymization).

---

(\*) still to investigate since PRNG fed with deterministic input is deterministic.

## References

<a id="1">[1]</a>
Luca Sciullo, Alberto De Marchi, Lorenzo Gigli, Monica Palmirani, and Fabio Vitali. 2024. AAA: A blockchain-based architecture for ethical, robust authenticated anonymity. In Proceedings of the 2024 International Conference on Information Technology for Social Good (GoodIT '24). Association for Computing Machinery, New York, NY, USA, 1. https://doi.org/10.1145/3677525.3678676
