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

以降の文章では、あなたのリポジトリに読み替えてほしい  

# Step [O1o1o0g2o_1o0] Visual Studio Code を使う

がんばって、 `Visual Studio Code` を使えるようにしておいてほしい  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

# Step [O1o1o0g2o0] 公式のサンプルプログラムを真似る

👇 とりあえずGo言語の公式のサンプルプログラムを一通りやろう。  

📖 [Documentation](https://go.dev/doc/)  

# Step [O1o1o0g2o1o_1o_10o0] Go言語のインストール

１年ぐらい経つと新しいバージョンが出ていたりするので、Go言語の最新バージョンをダウンロードしておくことにする。  
Windows向けには、 `.msi` ファイルが配られているので インストールは楽だろう  

古いバージョンの Go言語は削除することになる  

Go言語のインストール先は `C:\Program Files\Go\` とした  

# Step [O1o1o0g2o1o_1o_11o0] カレントディレクトリー

Visual Studio Code を開く  

私の場合、以下のディレクトリーへ移動する  

```shell
cd C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice
```

# Step [O1o1o0g2o1o_1o0] Go言語のバージョン確認

👇 以下のカレントディレクトリーから、以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

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

# Step [O1o1o0g2o1o_1o0] 設定 - go.work ファイル

Visual Studio Code のワークスペースを使いたいので、設定を行う  

📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  

👇 以下のカレントディレクトリーから、以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
👉  📂 kifuwarabe-uec14-practice
```

Input:  

```shell
go work init
go work use tools tools/gopls
```

👇 すると 以下のファイルが自動生成される  

```shell
    📂 kifuwarabe-uec14-practice
👉  └── 📄 go.work
```

```go
go 1.19

```

# Step [O1o1o0g2o1o_1o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集（無ければ新規作成）してほしい  

```shell
    📂 kifuwarabe-uec14-practice
  	├── 📄 go.work
👉  └── 📄 .gitignore
```

例えば冒頭に追加  

```plaintext
# * ここから、以下を追加
# (^q^)

go.work
# * ここまで


# ...略...
```

# Step [O1o1o0g2o1o0] モジュール作成

👇 以下のコマンドを入力してほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice
#           --------------------------------------------
#           1
# 1. この部分はあなたのレポジトリに合わせて変えてほしい
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
👉  ├── 📄 go.mod
	└── 📄 go.work
```

```go
module github.com/muzudho/kifuwarabe-uec14-practice

go 1.19
```

また、モジュールを追加したあとには以下のコマンドを入力してほしい  

```shell
go mod tidy
```

ソースから推論して必要なセットアップが自動的に行われるだろう  

👇 以下のファイルが自動生成される（コードを書き始めたばかりのときは、無いこともある）  

```shell
    📂 kifuwarabe-uec14-practice
👉  ├── 📄 go.sum
  	├── 📄 go.mod
	└── 📄 go.work
```

# Step [O1o1o0g3o0] ハローワールド

Go言語では フォルダーを、ファイルを小分けにするただの入れ物として使うには向いていない。  
大量のファイルをトップレベルに置いておくことも気にしないことにする  

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
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

		// この上に分岐を挟んでいく

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

# Step [O1o1o0g4o_1o0] クォート

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get rsc.io/quote
```

# Step [O1o1o0g4o0] クォート

👇 以下の既存ファイルを編集してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
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


	// * 以下を追加する
	} else if name == "quote" { // [O1o1o0g4o0]
		fmt.Println(quote.Go())

		// この上に分岐を挟んでいく
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

👇 以下のフォルダーを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
👉 	├── 📂 greetings
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
  	└── 📄 main.go
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd greetings
```

`📂 greetings` へ カレントディレクトリーが移動した  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init greetings
```

👇 以下のファイルが新規作成された  

```plaintext
    📂 kifuwarabe-uec14-practice
	├── 📂 greetings
👉	│	└── 📄 go.mod
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
  	└── 📄 main.go
```

```go
module greetings

go 1.19
```

👇 続けて

Input:  

```shell
go mod tidy
```

# Step [O1o1o0g5o1o0] あいさつ

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📄 go.mod
👉	│	└── 📄 greetings.go
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
  	└── 📄 main.go
```

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
```

# Step [O1o1o0g5o2o0] あいさつ

👇 以下の既存ファイルを編集してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📄 go.mod
	│	└── 📄 greetings.go
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
👉	└── 📄 main.go
```

```go
import (
	// ...略...


	// * 以下を追加する
	g "github.com/muzudho/kifuwarabe-uec14-practice/greetings"
)


// ...略...


	// * 以下を追加する
	} else if name == "hello" { // [O1o1o0g5o2o0]
		g.Hello("Nanashino Gonbee")

		// この上に分岐を挟んでいく
```

# 参考にした記事

## Go言語と Visual Studio Code

### go.work

📖 [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)  
📖 [VSCodeを開いたときのgoplsのmoduleに関するエラー](https://qiita.com/__tou__tou/items/d0867ed231e8bacd1b9e)  
📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  

おわり