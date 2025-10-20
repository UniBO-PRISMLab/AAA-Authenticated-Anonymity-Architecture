## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAAContract**                       |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 614,568 | 614,580 | 614,579 |    14 |       - |
|        *storeEncryptedPID*            |       - |       - | 612,953 |     2 |       - |
|        *storeEncryptedSID*            |       - |       - | 390,654 |     6 |       - |
|        *submitEncryptedWord*          | 718,436 | 910,717 | 772,840 |    33 |       - |
|        *submitRedundantEncryptedWord* |       - |       - | 278,734 |     2 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,499 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,393 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,324 |     2 |       - |

## Deployments
|                     |     Min |    Max  |       Avg | Block % | eur avg |
| :------------------ | ------: | ------: | --------: | ------: | ------: |
| **AAAContract**     |       - |       - | 2,557,644 |   8.5 % |       - |
| **MockAAALib**      |       - |       - |   472,628 |   1.6 % |       - |
| **MockUIPRegistry** | 489,686 | 536,172 |   495,749 |   1.7 % |       - |

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

