## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAAContract**                       |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 541,970 | 541,982 | 541,981 |    16 |       - |
|        *submitEncryptedPID*           | 149,283 | 149,295 | 149,289 |     4 |       - |
|        *submitEncryptedSID*           | 333,402 | 333,426 | 333,414 |     6 |       - |
|        *submitEncryptedWord*          | 634,884 | 788,301 | 673,188 |    50 |       - |
|        *submitSAC*                    |       - |       - |  54,900 |     8 |       - |
|        *submitSACRecord*              |       - |       - |  53,901 |     4 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,522 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,415 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,307 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAAContract**     | 2,571,204 | 2,617,702 | 2,613,475 |   8.7 % |       - |
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

