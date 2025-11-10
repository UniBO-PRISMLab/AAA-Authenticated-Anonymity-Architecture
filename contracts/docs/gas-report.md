## Methods
| **Symbol** | **Meaning**                                                                              |
| :--------: | :--------------------------------------------------------------------------------------- |
|    **◯**   | Execution gas for this method does not include intrinsic gas overhead                    |
|    **△**   | Cost was non-zero but below the precision setting for the currency display (see options) |

|                                       |     Min |     Max |     Avg | Calls | eur avg |
| :------------------------------------ | ------: | ------: | ------: | ----: | ------: |
| **AAA**                               |         |         |         |       |         |
|        *seedPhraseGenerationProtocol* | 541,971 | 541,983 | 541,981 |    17 |       - |
|        *submitEncryptedPID*           | 149,306 | 149,318 | 149,314 |     6 |       - |
|        *submitEncryptedSID*           | 333,403 | 333,427 | 333,418 |     8 |       - |
|        *submitEncryptedWord*          | 315,017 | 456,450 | 350,637 |    58 |       - |
|        *submitSAC*                    |       - |       - |  54,900 |     8 |       - |
|        *submitSACRecord*              |       - |       - |  53,902 |     4 |       - |
| **MockUIPRegistry**                   |         |         |         |       |         |
|        *addNode*                      |       - |       - |  74,500 |     5 |       - |
|        *onlyNodeFn*                   |       - |       - |  23,465 |     2 |       - |
|        *onlyOwnerFn*                  |       - |       - |  23,393 |     2 |       - |
|        *removeNode*                   |       - |       - |  37,325 |     2 |       - |

## Deployments
|                     |       Min |      Max  |       Avg | Block % | eur avg |
| :------------------ | --------: | --------: | --------: | ------: | ------: |
| **AAA**             | 2,861,424 | 2,907,922 | 2,903,879 |   9.7 % |       - |
| **MockAAALib**      |         - |         - |   472,628 |   1.6 % |       - |
| **MockUIPRegistry** |   599,642 |   646,128 |   604,807 |     2 % |       - |

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

