# github-action-practice

## 機能一覧
- [PRにコメントを追加する](./.github/workflows/pr_comment.yml)
- [actionの手動実行](./.github/workflows/param.yml)
- [変数・シークレットキーの設定](./.github/workflows/variable.yml)
- [スケジュールの設定](./.github/workflows/schedule.yml)
- [実行結果を環境変数に入れる](./.github/workflows/result_output.yml)
- [goのCIテスト](./.github/workflows/go_test.yml)
- [条件分岐](./.github/workflows/if.yml)
- [依存関係を自動アップデート](./.github/dependabot.yml)

## 開発
### コマンド
- github actionの確認

```shell
# 実行中
gh run watch
# 実行後
gh run view
# 手動実行
gh workflow run docker_image_publish -f version=0.0.1
```

## gh, gitコマンド
- PR関連

```bash
gh pr create --title "hoge" --body "hoge body"
```

- tagの作成
```bash
git tag vx.x.x -m 'hogehoge'
git push origin vx.x.x
```

- releaseノート
```bash
gh release create v.x.x --title "hoge title" --notes "hoge description"
```


## github package
- githubでパッケージを管理できる
```bash
# githubのユーザー名を使う
export GHCR_USER=$(gh config get -h github.com user)
# build
docker build -t ghcr.io/${GHCR_USER}/example:latest ./
# auth
gh auth refresh --scopes write:packages
gh auth token | docker login ghcr.io -u ${GHCR_USER} --password-stdin
docker push ghcr.io/${GHCR_USER}/example:latest
docker pull ghcr.io/${GHCR_USER}/example:latest
```

# AWS認証系
- gh-oidcリポジトリに作成
  - 認証系のため念の為privateリポジトリで練習
