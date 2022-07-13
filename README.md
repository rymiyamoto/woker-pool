# worker-pool

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
