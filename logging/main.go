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
