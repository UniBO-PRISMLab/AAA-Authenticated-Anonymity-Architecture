## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAA**                               |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 542,015 | 542,027 | 542,024 |    34 |       - |
|        *submitEncryptedPID*           | 489,404 | 489,416 | 489,412 |     3 |       - |
|        *submitEncryptedSID*           |       - |       - | 174,223 |     6 |       - |
|        *submitEncryptedWord*          | 356,923 | 497,411 | 387,110 |    68 |       - |
|        *submitRedundantWord*          | 626,663 | 626,699 | 626,680 |    12 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,522 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,415 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,307 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAA**             | 3,217,733 | 3,264,231 | 3,261,413 |  10.9 % |       - |
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

