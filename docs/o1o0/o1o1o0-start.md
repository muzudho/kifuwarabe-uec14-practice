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

この 📄 `go.work` ファイルは１個だけ作る。  
このファイルが置いてあると Go言語は `ワークスペースズ モード` になる

`ワークスペースズ モード` になると、  
今まで覚えた Go言語の `go run .` コマンドが使えなくなったり、  
今まで通っていたパッケージへのパスが通らなくなる。  
このような制限を解除する方法はあとで説明する  

また、せっかく作った 📂 `greetings`、 📂 `greetings\japanese`、 📂 `web-service-gin` だが、  
まだ ２番目、３番目、４番目のワークスペース にはできない。あとで説明する  

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
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------

go.work

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
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

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
👉  ├── 📄 go.sum
  	└── 📄 go.work
```

# Step [O1o1o0g3o0] サンプルプログラム - ハローワールド

Go言語では フォルダーを、ファイルを小分けにするただの入れ物として使うには向いていない。  
大量のファイルをトップレベルに置いておくことも気にしないことにする  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
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
go work use .
```

👇 以下のファイルが自動で編集されている  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
👉	├── 📄 go.work
  	└── 📄 main.go
```

```go
// ...略...


// * 以下の行が自動追加
use .
```

📄 `go.work` に `use .` の行が追加されたことで、カレントディレクトリーは Go言語のワークスペースになった。  
制限のかかっていた `go run .` コマンドが 再び使えるようになる  

👇 以下のファイルが自動生成されている  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
👉	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
github.com/ugorji/go v1.2.7 h1:qYhyWUUd6WbiM+C6JZAUkIJt/1WrjzNHY9+KCIjVqTo=
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

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . hello
```

Output:  

```plaintext
Hello, World!
```

Go言語を実行する環境は整ったようだ  

# Step [O1o1o0g4o_1o0] サンプルプログラム - クォート

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get rsc.io/quote
```

## Step [O1o1o0g4o0] クォート

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
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

# Step [O1o1o0g5o0] サンプルプログラム - あいさつ

👇 以下のサンプルプログラムの実装を行う  

📖 [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)  

## Step [O1o1o0g5o1o_1o0] カレントディレクトリーの移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd greetings
```

```plaintext
  	📂 kifuwarabe-uec14-practice
👉	├── 📂 greetings
	│	└── 📂 japanese
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

`📂 greetings` へ カレントディレクトリーが移動した  

## Step [O1o1o0g5o1o_2o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/greetings
```

👇 以下のファイルが新規作成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
👉	│	└── 📄 go.mod
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
module github.com/muzudho/kifuwarabe-uec14-practice/greetings

go 1.19
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
```

## Step [O1o1o0g5o1o0] あいさつ

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
👉	│	└── 📄 welcome.go
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g5o1o0]

package greetings

import "fmt"

// GetMessage returns a greeting for the named person.
func GetMessage(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// EOF [O1o1o0g5o1o0]
```

## Step [O1o1o0g5o2o_1o0] ワークスペースズ モードへの登録

👇 ここでカレントディレクトリーはまだ 📂 `greetings` だとする  

```plaintext
  	📂 kifuwarabe-uec14-practice
👉	├── 📂 greetings
	└── 📂 web-service-gin
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	└── 📄 go.mod
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
👉	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...

// * 以下が自動で削除
// use .
// * 以下が自動で追加
use (
	.
	./greetings
)
```

これにより、 📄 `main.go` から 📄 `greetings/welcome.go` ファイル（パッケージ）にパスを通すことができるようになる  

## Step [O1o1o0g5o2o0] あいさつ

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
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

## Step [O1o1o0g5o2o1o0] あいさつ

👇 カレントディレクトリーを戻してほしい  

```plaintext
👉  📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	└── 📂 web-service-gin
```

👇 そのためのコマンド例  

Input:  

```shell
cd ..
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
go run . greetings
```

* `go mod tidy` で `go.sum` が自動的に整理される

Output:  

```shell
Hi, Nanashino Gonbee. Welcome!
```

# Step [O1o1o0g6o0] サンプルプログラム - RESTful API

👇 以下のサンプルプログラムの実装を行う  

📖 [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)  

## Step [O1o1o0g6o1o_1o0] カレントディレクトリーの移動

👇 カレントディレクトリーを移動してほしい

```plaintext
    📂 kifuwarabe-uec14-practice
👉 	└── 📂 web-service-gin
```

👇 そのためのコマンド例  

Input:  

```shell
cd web-service-gin
```

## Step [O1o1o0g6o1o_1o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/web-service-gin
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
👉 	│	└── 📄 go.mod
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/web-service-gin

go 1.19
```

## Step [O1o1o0g6o1o_2o0] ワークスペースズ モードへの登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
 	│	└── 📄 go.mod
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
👉	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...

// * 以下が自動で削除
// use (
// 	   .
// 	   ./greetings
// )
// * 以下が自動で追加
use (
	.
	./greetings
	./web-service-gin
)
```

## Step [O1o1o0g6o1o_3o0] カレントディレクトリーを戻す

👇 カレントディレクトリーを移動してほしい  

```plaintext
👉	📂 kifuwarabe-uec14-practice
	└── 📂 web-service-gin
 		└── 📄 go.mod
```

👇 そのためのコマンド例  

```shell
cd ..
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod tidy
```

👇 これにより、以下の既存ファイルが自動で編集、または生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
👉	│	├── 📄 go.mod
👉	│	└── 📄 go.sum
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

内容は行数が多いので省略  

## Step [O1o1o0g6o1o0] データの作成

👇 カレントディレクトリーを移動してほしい  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 web-service-gin
```

👇 そのためのコマンド  

```shell
cd web-service-gin
```

## Step [O1o1o0g6o2o0] ビューおよびルート作成 - web-service-gin/main.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	└── 📄 main.go
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g6o2o0]

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

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// EOF [O1o1o0g6o2o0]
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go get .
```

必要なパッケージが自動的にダウンロードされるのだろう  

## Step [O1o1o0g6o3o_1o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run .
```

Output:  

```plaintext
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:8080
```

## Step [O1o1o0g6o3o_2o0] Webアクセス

👇 ブラウザを開いてアクセスしてほしい  

📖 [http://localhost:8080/albums](http://localhost:8080/albums)  

👇 画面に以下のように表示されるだろう  

```json
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
```

👇 ブラウザを開いてアクセスしてほしい  

📖 [http://localhost:8080/albums/2](http://localhost:8080/albums/2)  

👇 画面に以下のように表示されるだろう  

```json
{
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
}
```

`[Ctrl] + [C]` キーでサーバーを止めてほしい  

## Step [O1o1o0g6o3o0] カレントディレクトリーを戻す

👇 カレントディレクトリーを移動してほしい  

```plaintext
👉	📂 kifuwarabe-uec14-practice
	└── 📂 web-service-gin
```

👇 そのためのコマンドの例  

```shell
cd ..
```

# Step [O1o1o0g7o0] サンプルプログラム - ジェネリクス

👇 以下のサンプルプログラムの実装を行う  

📖 [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)  

## Step [O1o1o0g7o1o0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
👉	├── 📂 generics
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

## Step [O1o1o0g7o2o0] カレントディレクトリーを移動

👇 コマンドラインのプロンプトを見てほしい  

```shell
C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice>
# ------------------------------------------------------------------
# 11
# -------------------------------------------------------------------
# 10
#
# 10. （`C:` も含めて）これがプロンプト
# 11. カレントディレクトリー
```

👇 現在のカレントディレクトリーが以下とする  

```plaintext
👉	📂 kifuwarabe-uec14-practice
 	└── 📂 generics
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd generics
```

👇 コマンドラインのプロンプトを見てほしい  

```shell
C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice\generics>
```

👇 カレントディレクトリーは移動した  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 generics
```

## Step [O1o1o0g7o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/generics
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 generics
👉	│	└── 📄 go.mod
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/generics

go 1.19
```

## Step [O1o1o0g7o4o0] ワークスペースズ モードへの登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 generics
	│	└── 📄 go.mod
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
👉	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...

// * 以下が自動で削除
// use (
// 	   .
// 	   ./greetings
//     ./web-service-gin
// )
// * 以下が自動で追加
use (
	.
	./generics
	./greetings
	./web-service-gin
)
```

## Step [O1o1o0g7o5o0] プログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 generics
	│	├── 📄 go.mod
👉	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g7o5o0]

package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// EOF [O1o1o0g7o5o0]
```

## Step [O1o1o0g7o6o0] カレントディレクトリーを移動

👇 カレントディレクトリーを移動してほしい  

```plaintext
👉	📂 kifuwarabe-uec14-practice
	└── 📂 generics
```

👇 そのためのコマンドの例  

```shell
cd ../
```

## Step [O1o1o0g7o7o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
```

👇 これにより、以下の既存ファイルが自動で編集、または生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
👉 	├── 📄 go.mod
👉 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

内容は行数が多いので省略  

## Step [O1o1o0g7o8o0] カレントディレクトリーを移動

👇 カレントディレクトリーを移動してほしい  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 generics
```

👇 そのためのコマンドの例  

```shell
cd generics
```

## Step [O1o1o0g7o9o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run .
```

Output:  

```shell
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
Generic Sums with Constraint: 46 and 62.97
```

# Step [O1o1o0g8o0] サンプルプログラム - Fussing

👇 以下のサンプルプログラムの実装を行う  

📖 [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)  

## Step [O1o1o0g8o1o0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
👉	├── 📂 fuzz
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

## Step [O1o1o0g8o2o0] カレントディレクトリーを移動

👇 現在のカレントディレクトリーが以下とする  

```plaintext
👉	📂 kifuwarabe-uec14-practice
 	└── 📂 fuzz
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd fuzz
```

👇 カレントディレクトリーは移動した  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 fuzz
```

## Step [O1o1o0g8o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/fuzz
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
👉	│	└── 📄 go.mod
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/fuzz

go 1.19
```

## Step [O1o1o0g8o4o0] ワークスペースズ モードへの登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
👉	│	└── 📄 go.mod
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...

// * 以下が自動で削除
// use (
// 	   .
//     ./generics
// 	   ./greetings
//     ./web-service-gin
// )
// * 以下が自動で追加
use (
	.
	./fuzz
	./generics
	./greetings
	./web-service-gin
)
```

## Step [O1o1o0g8o5o0] プログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
👉	│	└── 📄 main.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g8o5o0]

package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

// EOF [O1o1o0g8o5o0]
```

## Step [O1o1o0g8o5o1o0] プログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
👉	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g8o5o1o0]

package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// EOF [O1o1o0g8o5o1o0]
```

## Step [O1o1o0g8o6o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run .
```

Output:  

```plaintext
original: "The quick brown fox jumped over the lazy dog"
reversed: "god yzal eht revo depmuj xof nworb kciuq ehT", err: <nil>
reversed again: "The quick brown fox jumped over the lazy dog", err: <nil>
```

## Step [O1o1o0g8o7o0] テスト

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go test
```

Output:  

```shell
PASS
ok      github.com/muzudho/kifuwarabe-uec14-practice/fuzz       0.219s
```

# Step [O1o1o0g9o0] サンプルプログラム - Wiki

👇 以下のサンプルプログラムの実装を行う  

📖 [Writing Web Applications](https://go.dev/doc/articles/wiki/)  

## Step [O1o1o0g9o1o0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
👉	├── 📂 gowiki
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

## Step [O1o1o0g9o2o0] カレントディレクトリーを移動

👇 現在のカレントディレクトリーが以下とする  

```plaintext
👉	📂 kifuwarabe-uec14-practice
 	└── 📂 gowiki
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd gowiki
```

👇 カレントディレクトリーは移動した  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 gowiki
```

## Step [O1o1o0g9o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/gowiki
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
👉	│	└── 📄 go.mod
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/gowiki

go 1.19
```

## Step [O1o1o0g9o4o0] ワークスペースズ モードへの登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	└── 📄 go.mod
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
👉	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...

// * 以下が自動で削除
// use (
// 	   .
//     ./fuzz
//     ./generics
// 	   ./greetings
//     ./web-service-gin
// )
// * 以下が自動で追加
use (
	.
	./fuzz
	./generics
	./gowiki
	./greetings
	./web-service-gin
)
```

## Step [O1o1o0g9o5o0] Vol1 プログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
👉	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g9o5o0]

package main

import (
	"flag"
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "vol1" { // [O1o1o0g9o5o0]
		p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
		p1.save()
		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.Body))

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o5o0]
```

### Step [O1o1o0g9o6o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . vol1
```

Output:  

```shell
This is a sample Page.
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
👉	│	├── 📄 TestPage.txt
	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

## Step [O1o1o0g9o7o_1o0] Vol2 プログラム作成 - simple.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
👉	│	├── 📄 simple.go
	│	├── 📄 TestPage.txt
	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g9o7o_1o0]

package main

import (
	"fmt"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// EOF [O1o1o0g9o7o_1o0]
```

先頭行に `//go:build ignore` と書くと Go言語のビルド対象から外れるようだが、  
ビルド対象に含めることにした  

### Step [O1o1o0g9o7o_2o0] プログラム編集

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
	│	├── 📄 simple.go
	│	├── 📄 TestPage.txt
👉	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

👇 抜粋  

```go
// ...略...


	// * 以下を追加する
	} else if name == "vol2" { // [O1o1o0g9o7o_2o0]
		http.HandleFunc("/", SimpleHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))

		// この上に分岐を挟んでいく
```

### Step [O1o1o0g9o7o_3o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . vol2
```

### Step [O1o1o0g9o7o_4o0] Webアクセス

👇 ブラウザを開いてアクセスしてほしい  

📖 [http://localhost:8080/albums](http://localhost:8080/albums)  

👇 画面に以下のように表示されるだろう  

```plaintext
Hi there, I love albums!
```

### Step [O1o1o0g9o7o_5o0] Webアクセス

👇 ブラウザを開いてアクセスしてほしい  

📖 [http://localhost:8080/monkeys](http://localhost:8080/monkeys)  

👇 画面に以下のように表示されるだろう  

```plaintext
Hi there, I love monkeys!
```

## Step [O1o1o0g9o7o0] Vol3 プログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
	│	├── 📄 simple.go
	│	├── 📄 TestPage.txt
👉	│	├── 📄 view.go
	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g9o7o0]

package main

import (
	"fmt"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// EOF [O1o1o0g9o7o0]
```

### Step [O1o1o0g9o8o0] プログラム編集

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
	│	├── 📄 simple.go
	│	├── 📄 TestPage.txt
	│	├── 📄 view.go
👉	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

👇 抜粋  

```go
// ...略...


	// * 以下を追加する
	} else if name == "vol3" { // [O1o1o0g9o8o0]
		http.HandleFunc("/view/", ViewHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))

		// この上に分岐を挟んでいく
```

### Step [O1o1o0g9o9o_1o0] ページ作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
	│	├── 📄 simple.go
👉	│	├── 📄 test.txt
	│	├── 📄 TestPage.txt
	│	├── 📄 view.go
	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
Hello world
```

### Step [O1o1o0g9o9o0] ビルド

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go build .
```

カレントディレクトリーにある `*.go` ファイルをビルドしました  

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	│	├── 📄 go.mod
	│	├── 📄 main.go
	│	└── 📄 reverse_test.go
	├── 📂 generics
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 gowiki
	│	├── 📄 go.mod
👉	│	├── 📄 gowiki.exe
	│	├── 📄 simple.go
	│	├── 📄 TestPage.txt
	│	├── 📄 view.go
	│	└── 📄 wiki.go
	├── 📂 greetings
	│	├── 📂 japanese
	│	├── 📄 go.mod
	│	└── 📄 welcome.go
	├── 📂 web-service-gin
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	└── 📄 main.go
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

### Step [O1o1o0g9o10o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
# Windows
gowiki.exe vol3
```

### Step [O1o1o0g9o11o0] Webアクセス

👇 ブラウザを開いてアクセスしてほしい  

📖 [http://localhost:8080/view/test](http://localhost:8080/view/test)  

👇 画面に以下のように表示されるだろう  

```plaintext
test

Hello world
```

# Step [O1o1o0g10o0] ロギング

👇 公式のサンプルを一旦離れて　ロギングを練習する  

📖 [Zap](https://github.com/uber-go/zap)  

## Step [O1o1o0g10o1o0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
👉	├── 📂 logging
	├── 📂 web-service-gin
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

## Step [O1o1o0g10o2o0] カレントディレクトリーを移動

👇 コマンドラインのプロンプトを見てほしい  

```shell
C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice>
# ------------------------------------------------------------------
# 11
# -------------------------------------------------------------------
# 10
#
# 10. （`C:` も含めて）これがプロンプト
# 11. カレントディレクトリー
```

👇 現在のカレントディレクトリーが以下とする  

```plaintext
👉	📂 kifuwarabe-uec14-practice
 	└── 📂 logging
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd logging
```

👇 コマンドラインのプロンプトを見てほしい  

```shell
C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-uec14-practice\logging>
```

👇 カレントディレクトリーは移動した  

```plaintext
	📂 kifuwarabe-uec14-practice
👉	└── 📂 logging
```

## Step [O1o1o0g10o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14-practice/logging
```

Output:  

```plaintext
go: creating new go.mod: module github.com/muzudho/kifuwarabe-uec14-practice/logging
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
	├── 📂 logging
👉	│	└── 📄 go.mod
	├── 📂 web-service-gin
  	├── 📄 .gitignore
 	├── 📄 go.mod
 	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14-practice/logging

go 1.19
```

## Step [O1o1o0g10o4o0] ワークスペースズ モードへの登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go work use .
```

👇 これにより、以下の既存ファイルが自動で編集された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
	├── 📂 logging
	│	└── 📄 go.mod
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
👉	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// ...略...


use (
// ...略...


	./logging


// ...略...
)
```

## Step [O1o1o0g10o5o_1o0] インストール

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u go.uber.org/zap
```

## Step [O1o1o0g10o5o_2o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
	├── 📂 logging
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 web-service-gin
👉 	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...

*.log

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O1o1o0g10o5o0] メインプログラム作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
	├── 📂 logging
	│	├── 📄 go.mod
👉	│	└── 📄 main.go
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g10o5o0]

package main

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createLoggerForConsole - ロガー作成，コンソール用
func createLoggerForConsole(logFile *os.File) *zap.Logger {
	// 設定，コンソール用
	var config1 = zapcore.EncoderConfig{
		MessageKey: "message",

		// LevelKey:    "level",
		// EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder, // 日本時間のタイムスタンプ

		// CallerKey:    "caller",
		// EncodeCaller: zapcore.ShortCallerEncoder,
	}
	// config1.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// 設定，ファイル用
	var config2 = zap.NewProductionEncoderConfig()
	config2.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config1), // コンソール形式
			zapcore.Lock(os.Stderr),            // 出力先は標準エラー
			zapcore.DebugLevel),                // ログレベル
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config2), // コンソール形式
			zapcore.AddSync(logFile),           // 出力先はログファイル
			zapcore.DebugLevel),                // ログレベル
	)

	// ロガーのビルド
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// createLoggerForJson - ロガー作成，複数行JSON形式
func createLoggerAsJson(logFile *os.File) *zap.Logger {
	// 設定
	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON形式
			zapcore.AddSync(logFile),       // 出力先はログファイル
			zapcore.DebugLevel),            // ログレベル
	)

	// ロガーのビルド
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// func createLogger2(logFile *os.File) *zap.Logger {
// 	// ロガー作成
// 	// logger, _ := zap.NewDevelopment() // 開発用
// 	logger, _ := zap.NewProduction() // 製品用
// 	// ロガー破棄時にバッファーをフラッシュする
// 	defer logger.Sync() // flushes buffer, if any

// 	return logger
// }

func main() {
	// ファイル，基本
	basicLogFile, _ := os.OpenFile("example-basic.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer basicLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// `log` インターフェースを直接使うなら
	log.SetOutput(basicLogFile) // ロガーにログファイルを紐づけ
	log.Println("Hello, world!")

	// ファイル，コンソール用
	logcFile, _ := os.OpenFile("example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logcFile.Close() // ログファイル使用済み時にファイルを閉じる

	// ファイル，複数行JSON形式
	logjFile, _ := os.OpenFile("example-json.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logcFile.Close() // ログファイル使用済み時にファイルを閉じる

	// カスタマイズしたロガーを使うなら
	var logc = createLoggerForConsole(logcFile).Sugar()
	var logj = createLoggerAsJson(logjFile).Sugar()

	// Run
	var url = "http://tic.warabenture.com"
	logc.Infof("Failed to fetch URL: %s", url)
	logj.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
}

// EOF [O1o1o0g10o5o0]
```

## Step [O1o1o0g10o6o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run .
```

Output:  

```plaintext
2022-09-11T14:36:31.819+0900    Failed to fetch URL: http://tic.warabenture.com
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14-practice
	├── 📂 fuzz
	├── 📂 generics
	├── 📂 gowiki
	├── 📂 greetings
	├── 📂 logging
👉	│	├── 📄 example-basic.log
👉	│	├── 📄 example-json.log
👉	│	├── 📄 example.log
	│	├── 📄 go.mod
	│	└── 📄 main.go
	├── 📂 web-service-gin
  	├── 📄 .gitignore
  	├── 📄 go.mod
  	├── 📄 go.sum
	├── 📄 go.work
	├── 📄 go.work.sum
  	└── 📄 main.go
```

👇 📄 `example-basic.log`  

```plaintext
2022/09/11 14:30:26 Hello, world!
```

👇 📄 `example-json.log`  

```plaintext
{"level":"info","ts":"2022-09-11T14:30:26.223+0900","caller":"logging/main.go:93","msg":"failed to fetch URL","url":"http://tic.warabenture.com","attempt":3,"backoff":1}
```

👇 📄 `example.log`  

```plaintext
2022-09-11T14:30:26.222+0900	info	logging/main.go:92	Failed to fetch URL: http://tic.warabenture.com
```

* 作成されるログファイルは JSON形式ではない。 ワンライナーのJSONが複数行並ぶ

# 次の記事

📖 [Go [O1o1o0] 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# 参考にした記事

📖 [Go [O1o1o0] 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会](https://qiita.com/muzudho1/items/cea62be01f7418bbf150) - この記事  

## Go言語

### モジュール

📖 [[Golang] Module](https://dev-yakuza.posstree.com/en/golang/module/)  

### マルチ ワークスペース

📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  
📖 [Go 1.18集中連載 Workspacesモードを試してみた](https://future-architect.github.io/articles/20220216a/)  

### ロギング

📖 [How to log to stdout or stderr based on log level using uber-go/zap?](https://stackoverflow.com/questions/68472667/how-to-log-to-stdout-or-stderr-based-on-log-level-using-uber-go-zap)  
📖 [5 structured logging packages for Go](https://blog.logrocket.com/5-structured-logging-packages-for-go/)  
📖 [Structured Logging in Golang with Zap – Blazing Fast Logger](https://codewithmukesh.com/blog/structured-logging-in-golang-with-zap/)  
📖 [【Go】ログライブラリzapを使う時はglobal loggerが便利！](https://qiita.com/yagi_eng/items/2957ef04cebcdeaae1d6)  
📖 [GolangのLogger「zap」の最低限の設定方法](https://tomokazu-kozuma.com/minimum-setting-method-of-golangs-logger-zap/)  
📖 [zapでログレベルでログの出力先を振り分ける方法](https://qiita.com/emonuh/items/cb3a730979dba7d76920)  
📖 [golangの高速な構造化ログライブラリ「zap」の使い方](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)  
📖 [GoのロギングライブラリzapのTips](https://christina04.hatenablog.com/entry/golang-zap-tips)  
📖 [golang zap v1.0.0 でログの日付をJSTで表示する方法](https://qiita.com/fuku2014/items/c6501c187c8161336485)  
📖 [Using Zap - Creating custom loggers](https://blog.sandipb.net/2018/05/03/using-zap-creating-custom-loggers/)  

## Go言語と Visual Studio Code

### go.work

📖 [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)  
📖 [VSCodeを開いたときのgoplsのmoduleに関するエラー](https://qiita.com/__tou__tou/items/d0867ed231e8bacd1b9e)  
📖 [Setting up your workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)  

おわり