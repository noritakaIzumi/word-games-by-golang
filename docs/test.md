# テスト

## 実行対象ファイル

`/pkg` 以下のファイルとなります。

## 実行方法

以下のコマンドはリポジトリ直下で実行してください。

- 全ファイル実行

```shell script
go test -v -race -coverpkg=./pkg/... ./test/...
```

- カバレッジレポートを出力したい場合

```shell script
go test -v -race -coverpkg=./pkg/... -coverprofile=coverage.out ./test/... \
  && go tool cover -html=coverage.out -o coverage.html
```

- 特定のパッケージをテストする場合

```shell script
go test -v -race -coverpkg=./pkg/パッケージ名 ./test/パッケージ名_test
```

## テストスクリプトについて

### 対象

対象は `/pkg` 内のソースコードとします。この範囲のコードについては、テストコードを記述してください。
外部パッケージとして実装してください。

- 例

`linked_list` パッケージのテストは `linked_list_test` パッケージとして実装します。

### カバレッジについて

C1 カバレッジ 80 % 以上となるようにテストコードを記述してください。
**通常のテストツールでは C0 カバレッジしか計測できないため、条件分岐については以下のようなルールで実装してください。**

- `if-else` 分岐

`if` のみの場合、 `if` を通っていないテストケースが見逃されやすいため、何らかの `else` 句を用意するか `if` 句内では `return` してください。

```go
package main
func main() {
    a := 1
    //goland:noinspection ALL
    if a == 0 {
        // do something...
    } else {
        // do something...
    }
}
```

- レシーバの not nil チェック部分

実際にはレシーバが `nil` になり得ない場合もあるので、カバーできない場合はしなくてもいいです。

- テストの存在しないパッケージ

特定のパッケージについて、関数を実行するファイルが一つもない場合、そもそもカバレッジが算出されません。
新たにパッケージを作成した場合は、その時テストを書かない場合でもパッケージ内の関数を実行するようなテストコードを用意しておいてください。

> `pkg/io` 及び `test/io_test` パッケージを参考にしてください。
