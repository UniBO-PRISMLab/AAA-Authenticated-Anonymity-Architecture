## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAAContract**                       |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 541,980 | 542,004 | 542,003 |    16 |       - |
|        *submitEncryptedPID*           |       - |       - | 149,295 |     4 |       - |
|        *submitEncryptedSID*           | 333,400 | 333,448 | 333,432 |     6 |       - |
|        *submitEncryptedWord*          | 634,920 | 788,301 | 673,194 |    50 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,522 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,415 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,307 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAAContract**     | 2,356,394 | 2,402,892 | 2,397,726 |     8 % |       - |
| **MockAAALib**      |         - |         - |   472,628 |   1.6 % |       - |
| **MockUIPRegistry** |   583,653 |   630,139 |   588,818 |     2 % |       - |

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

