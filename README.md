# worker-pool

refs https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923

## Getting Started

事前に以下のものを入れてください

| 項目                                                       | 目的                    |
| ---------------------------------------------------------- | ----------------------- |
| [direnv](https://github.com/direnv/direnv)                 | .envrc 利用のため       |
| [golangci-lint](https://github.com/golangci/golangci-lint) | linter                  |
| [pre-commit](https://pre-commit.com)                       | commit 時の事前チェック |

## Setup

```sh
# 環境変数に追加
$ cp .envrc.dev .envrc
$ vim .envrc
$ direnv allow .

# pre-commit の追加
$ pre-commit install
```
