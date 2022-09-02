# Step [O1o1o0g1o0] はじめに

いきなり　囲碁プログラム　を書くのは大変なので、練習をする  

👇 また、技術的でない内容を含むブログを別の場所に 開設する  

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会☆（＾ｑ＾）＜その１＞](http://grayscale2.dou-jin.com/go/20220829)  

アーキテクチャ:  

| What is           | This is            |
| ----------------- | ------------------ |
| OS                | Windows            |
| Editor            | Visual Studio Code |
| Program Language  | Go                 |
| Remote Repository | Git Hub            |

# Step [O1o1o0g1o1o0] ソースの置き場所（リモート）

👇 まず、ソースの置き場所を決めておく  

📖 [リポジトリ](https://github.com/muzudho/kifuwarabe-uec14-practice)  

# Step [O1o1o0g1o2o0] ソースの置き場所（ローカル）

Go言語では ローカルPCのどこにソースを置くかは自分で設定して決めておく。  
サンプルでは　ユーザーホームの下に置いているので、真似る  

👇 以下のコマンドを入力してほしい  

Input:  

```shell
echo %HOMEPATH%
```

Output:  

```plaintext
\Users\むずでょ
```

ユーザーホームのパスが分かった。この下に `go\src` で始まるディレクトリーを作っていく。  
私は以下の場所にした  

`C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice`  

以降の文章では、あなたのリポジトリに読み替えてほしい  

# Step [O1o1o0g2o_1o0] Visual Studio Code を使う

がんばって、 `Visual Studio Code` を使えるようにしておいてほしい  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

# Step [O1o1o0g2o_1o0] Goエクステンションをインストールする

`Visual Studio Code` の `Extensions` から、 `Go` をインストールしておいてほしい  

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

```plaintext
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

# Step [O1o1o0g2o1o_1o0] マルチ ワークスペース

べつに ワークスペースは１つでいいが、  
練習で 複数のワークスペースを作れることを示す  

👇 以下の３つのディレクトリーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
👉	├── 📂 greetings
👉	│	└── 📂 japanese
👉	└── 📂 web-service-gin
```

👇 カレントディレクトリーは以下の通りとする  

```plaintext
👉  📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	└── 📂 web-service-gin
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work init
```

👇 すると 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
👉  └── 📄 go.work
```

```go
go 1.19

```

この 📄 `go.work` ファイルは１個だけ作る  

# Step [O1o1o0g2o1o_1o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集（無ければ新規作成）してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
👉  ├── 📄 .gitignore
  	└── 📄 go.work
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

👇 カレントディレクトリーは以下の通りとする  

```plaintext
👉  📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	└── 📂 web-service-gin
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice
#           --------------------------------------------
#           1
# 1. モジュール名。この部分はあなたのレポジトリに合わせて変えてほしい
```

Output:  

```plaintext
go: creating new go.mod: module example/main
go: to add module requirements and sums:
        go mod tidy
```

👇 すると 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
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
👉	│	└── 📄 welcome.go
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
  	└── 📄 main.go
```

```go
package greetings

import "fmt"

// GetMessage returns a greeting for the named person.
func GetMessage(name string) string {
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
	│	└── 📄 welcome.go
  	├── 📄 go.sum
    ├── 📄 go.mod
	├── 📄 go.work
👉	└── 📄 main.go
```

```go
import (
	// ...略...


	// * 以下は、ファイルの保存時に Visual Studio Code の Go エクステンションによって自動で補完される
	"github.com/muzudho/kifuwarabe-uec14-practice/greetings"
)


// ...略...


	// * 以下を追加する
	} else if name == "greetings" { // [O1o1o0g5o2o0]
		fmt.Println(greetings.GetMessage("Nanashino Gonbee"))

		// この上に分岐を挟んでいく
```

# Step [O1o1o0g5o2o1o0] あいさつ

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
go run . greetings
```

Output:  

```shell
Hi, Nanashino Gonbee. Welcome!
```

# Step [O1o1o0g6o0] RESTful API

📖 [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)  

👇 以下のフォルダーを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
👉 	└── 📂 web-service-gin
```

👇 カレントディレクトリーを移動してほしい

```plaintext
    📂 kifuwarabe-uec14-practice
👉 	└── 📂 web-service-gin
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd web-service-gin
go mod init github.com/muzudho/kifuwarabe-uec14-practice/web-service-gin
```

👇 以下のファイルが自動生成された  

```plaintext
    📂 kifuwarabe-uec14-practice
	└── 📂 web-service-gin
👉 		└── 📄 go.mod
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/web-service-gin

go 1.19
```

👇 カレントディレクトリーを移動してほしい  

```plaintext
👉	📂 kifuwarabe-uec14-practice
	└── 📂 web-service-gin
 		└── 📄 go.mod
```

```shell
cd ..
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work init web-service-gin
```

👇 以下のファイルが新規作成される  

```plaintext
	📂 kifuwarabe-uec14-practice
	├── 📂 web-service-gin
 	│	└── 📄 go.mod
👉	├── 📄 go.work
👉	└── 📄 go.work.sum
```

📄 go.work  

```go
go 1.19

use ./web-service-gin
```

📄 go.work.sum

```go
github.com/ugorji/go v1.2.7 h1:qYhyWUUd6WbiM+C6JZAUkIJt/1WrjzNHY9+KCIjVqTo=
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod tidy
```

# Step [O1o1o0g6o1o0] データの作成

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂 kifuwarabe-uec14-practice
	└── 📂 web-service-gin
 		├── 📄 go.mod
👉 		└── 📄 main.go
```

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go get .
```

必要なパッケージが自動的にダウンロードされるのだろう  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go run .
```

# 参考にした記事

📖 [Go [O1o1o0] 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会](https://qiita.com/muzudho1/items/cea62be01f7418bbf150) - この記事  

## Go言語

### モジュール

📖 [[Golang] Module](https://dev-yakuza.posstree.com/en/golang/module/)  

### マルチ ワークスペース

📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  
📖 [Go 1.18集中連載 Workspacesモードを試してみた](https://future-architect.github.io/articles/20220216a/)  

## Go言語と Visual Studio Code

### go.work

📖 [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)  
📖 [VSCodeを開いたときのgoplsのmoduleに関するエラー](https://qiita.com/__tou__tou/items/d0867ed231e8bacd1b9e)  
📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  

おわり