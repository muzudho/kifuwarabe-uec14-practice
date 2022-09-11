// BOF [O1o1o0g10o5o0]

package main

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger(logFile *os.File) *zap.Logger {
	// 設定
	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON形式
			zapcore.Lock(os.Stderr),        // 出力先は標準エラー
			zapcore.DebugLevel),            // ログレベル
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
	// ファイル
	logFile, _ := os.OpenFile("example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close() // ログファイル使用済み時にファイルを閉じる

	// `log` インターフェースを直接使うなら
	log.SetOutput(logFile) // ロガーにログファイルを紐づけ
	log.Println("Hello, world!")

	// カスタマイズしたロガーを使うなら
	var logger = createLogger(logFile)

	// Run
	// 糖衣構文のインターフェースを取得
	sugar := logger.Sugar()
	var url = "http://tic.warabenture.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// EOF [O1o1o0g10o5o0]
