# github-action-practice

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