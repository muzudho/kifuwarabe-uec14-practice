# Step [O1o1o0g1o0] はじめに

いきなり　囲碁プログラム　を書くのは大変なので、練習をする  

👇 また、技術的でない内容を含むブログを別の場所に 開設する  

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会☆（＾ｑ＾）＜その１＞](http://grayscale2.dou-jin.com/go/20220829)  

# Step [O1o1o0g1o1o0] ソースの置き場所（リモート）

👇 まず、ソースの置き場所を決めておく  

📖 [リポジトリ](https://github.com/muzudho/kifuwarabe-uec14-practice)  

# Step [O1o1o0g1o2o0] ソースの置き場所（ローカル）

Go言語では ローカルPCのどこにソースを置くかは自分で設定して決めておく。  
私は以下の場所にした  

`C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice`  

# Step [O1o1o0g2o0] 公式のサンプルプログラムを真似る

👇 とりあえずGo言語の公式のサンプルプログラムを一通りやろう。  

📖 [Documentation](https://go.dev/doc/)  

# Step [O1o1o0g2o1o_1o_10o0] Go言語のインストール

１年ぐらい経つと新しいバージョンが出ていたりするので、Go言語の最新バージョンをダウンロードしておくことにする。  
Windows向けには、 `.msi` ファイルが配られているので インストールは楽だろう  

古いバージョンの Go言語は削除することになる  

Go言語のインストール先は `C:\Program Files\Go\` とした  

# Step [O1o1o0g2o1o_1o0] Go言語のバージョン確認

👇 以下のカレントディレクトリーから、コマンドを入力してほしい  

```shell
👉  📂 kifuwarabe-uec14-practice
```

Input:  

```shell
go version
```

Output:  

```plaintext
go version go1.19 windows/amd64
```

# Step [O1o1o0g2o1o0] モジュール作成

👇 以下のコマンドを入力してほしい  

Input:  

```shell
go mod init example/main
```

Output:  

```plaintext
go: creating new go.mod: module example/main
go: to add module requirements and sums:
        go mod tidy
```

👇 すると 以下のファイルが自動生成される  

```shell
    📂 kifuwarabe-uec14-practice
👉  └── 📄 go.mod
```

```go
module example/main

go 1.19
```

# Step [O1o1o0g3o0] ハローワールド

Go言語では フォルダーを、ファイルを小分けにするただの入れ物として使うには向いていない。  
大量のファイルをトップレベルに置いておくことも気にしないことにする  

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
    ├── 📄 go.mod
👉  └── 📄 main.go
```

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" { // [O1o1o0g3o0]
		fmt.Println("Hello, World!")

		// ここに分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run .
```

Output:  

```plaintext
go run . {programName}
```

Input:  

```shell
go run . hello
```

Output:  

```plaintext
Hello, World!
```

Go言語を実行する環境は整ったようだ  

# Step [O1o1o0g4o0] クォート

👇 以下の既存ファイルを編集してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
    ├── 📄 go.mod
👉  └── 📄 main.go
```

👇 抜粋  

```go
import (
	// ...略...


	// * 以下を追加する
	"rsc.io/quote"
)


// ...略...


		// ここに分岐を挟んでいく

	// * 以下を追加する
	} else if name == "quote" { // [O1o1o0g4o0]
		fmt.Println(quote.Go())
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
go run . quote
```

Output:  

```plaintext
Don't communicate by sharing memory, share memory by communicating.
```

以上のようなメッセージが出てくる。  
このサンプルプログラムは何をやっているのか分からない。  

# Step [O1o1o0g5o0] あいさつ

📖 [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)  
