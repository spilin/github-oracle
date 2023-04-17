# Github Orcacle backend

```bash
solc --abi --bin oracle.sol -o build
abigen --bin=./build/GHOracle.bin --abi=./build/GHOracle.abi --pkg=oracle --out=contracts/GHOracle.go
```