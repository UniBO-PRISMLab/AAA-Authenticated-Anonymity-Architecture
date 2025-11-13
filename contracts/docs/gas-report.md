## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAA**                               |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 541,971 | 541,983 | 541,982 |    34 |       - |
|        *submitEncryptedPID*           |       - |       - | 489,416 |     3 |       - |
|        *submitEncryptedSID*           | 174,233 | 174,245 | 174,241 |     6 |       - |
|        *submitEncryptedWord*          | 382,707 | 523,184 | 412,881 |    68 |       - |
|        *submitRedundantWord*          | 628,894 | 629,047 | 628,998 |    12 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,522 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,415 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,307 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAA**             | 3,227,390 | 3,273,888 | 3,271,070 |  10.9 % |       - |
| **MockAAALib**      |         - |         - |   472,628 |   1.6 % |       - |
| **MockUIPRegistry** |   617,955 |   664,441 |   623,120 |   2.1 % |       - |

## Solidity and Network Config
| **Settings**        | **Value**  |
| ------------------- | ---------- |
| Solidity: version   | 0.8.28     |
| Solidity: optimized | true       |
| Solidity: runs      | 800        |
| Solidity: viaIR     | false      |
| Block Limit         | 30,000,000 |
| Gas Price           | -          |
| Token Price         | -          |
| Network             | ETHEREUM   |
| Toolchain           | hardhat    |

