## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAA**                               |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* |       - |       - | 631,439 |    34 |       - |
|        *submitEncryptedPID*           | 489,392 | 489,416 | 489,408 |     3 |       - |
|        *submitEncryptedSID*           | 174,255 | 174,267 | 174,261 |     6 |       - |
|        *submitEncryptedWord*          | 348,063 | 488,078 | 378,049 |    68 |       - |
|        *submitRedundantWord*          | 626,697 | 626,721 | 626,710 |    12 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,557 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,415 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,371 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAA**             | 3,298,311 | 3,344,809 | 3,341,991 |  11.1 % |       - |
| **MockAAALib**      |         - |         - |   472,628 |   1.6 % |       - |
| **MockUIPRegistry** |   681,737 |   728,223 |   686,902 |   2.3 % |       - |

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

