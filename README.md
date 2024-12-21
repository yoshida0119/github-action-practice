# github-action-practice

## 機能一覧
- [PRにコメントを追加する](./.github/workflows/pr_comment.yml)
- [実行結果を環境変数に入れる](./.github/workflows/result_output.yml)

## 開発
### コマンド
- github actionの確認

```shell
# 実行中
gh run watch
# 実行後
gh run view
```

## 記述
- workflowを手動で実行する方法

```yaml
on: 
  workflow_dispatch:
```

```shell
# このコマンドで実行できる
gh workflow run param.yml -f greeting=hello
```

- scheduleの設定

```yaml
on: 
  schedule:
    - cron: '*/15 * * * *' # 15分ごと
```

- 環境変数
  - varsはgithubに設定されている環境変数
```yaml
env:
  EXAMPLE: workflow sample text
    ...
jobs:
    steps:
      - run: echo "${EXAMPLE}"
      - uses: actions/checkout@v4
      - run: echo ${{ vars.USERNAME }}
```
