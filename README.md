# Word Games

以下の勉強目的で作成しています。

- Go 言語
- 連結リスト

## 必要な環境について

Go または Docker をご用意ください。

## 実行方法

### Go を使用する場合

**アプリ名** は **これまでに実装したアプリの一覧** の項目名を参照してください。

```shell script
cd /path/to/word_games
go run cmd/main.go アプリ名
```

### Docker を使用する場合

```shell script
cd /path/to/word_games
./build/run_app.sh アプリ名
```

## 開発する方へ

コードを書き始める前に `/doc` 配下の資料をお読みください。
新しいアプリの開発が終了したら、下記 **これまでに実装したアプリの一覧** に追記してください。

## これまでに実装したアプリの一覧

### `is_palindrome`

単語や文字列が左右対称かどうかを判断します。

```text
$ go run cmd/main.go is_palindrome
Please enter some word: pop
The word "pop" is palindrome!
```

### `get_common_ending`

2 つの単語に共通する語尾を取得します。

```text
$ go run cmd/main.go get_common_ending
Please enter first word: common
Please enter second word: salmon
The word "common" and the word "salmon" have the common ending "mon".
```
