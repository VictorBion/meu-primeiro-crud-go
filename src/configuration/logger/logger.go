package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
    log *zap.Logger

    LOG_OUTPUT = "LOG_OUTPUT"
    LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
    logConfig := zap.Config{
		//Define onde os logs serão enviados. Pode ser "stdout" (terminal), "stderr", ou um caminho de arquivo. A função getOutputLogs() lê a variável de ambiente LOG_OUTPUT para decidir isso.
        OutputPaths: []string{getOutputLogs()},
		//Define o nível mínimo de log que será exibido (ex: info, error, debug). A função getLevelLog() lê a variável LOG_LEVEL e retorna o nível desejado. NewAtomicLevelAt permite mudar o nível dinamicamente depois, se quiser.
        Level:       zap.NewAtomicLevelAt(getLevelLog()),
        Encoding:    "json",
		//Define como os campos do log serão formatados. Aqui você personaliza os nomes e estilos dos dados que aparecem no log.
        EncoderConfig: zapcore.EncoderConfig{
			//O nome do campo que vai mostrar o nível do log (info, error, etc.) no JSON.
            LevelKey:     "level",
			//O nome do campo que vai mostrar a data e hora do log
            TimeKey:      "time",
			//O nome do campo que vai mostrar a mensagem principal do log.
            MessageKey:   "message",
			//Define o formato da data/hora. ISO8601 é um padrão internacional (2025-10-27T04:58:00Z).
            EncodeTime:   zapcore.ISO8601TimeEncoder,
			//Define que o nível do log será exibido em letras minúsculas (info, error, debug).
            EncodeLevel:  zapcore.LowercaseLevelEncoder,
			//➡️ Adiciona o local do código que gerou o log (arquivo e linha), em formato curto (main.go:42).
            EncodeCaller: zapcore.ShortCallerEncoder,
        },
    }

    var err error
    log, err = logConfig.Build()
    if err != nil {
        panic(err)
    }
}

func Info(message string, tags ...zap.Field) {
    log.Info(message, tags...)
    _ = log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
    tags = append(tags, zap.NamedError("error", err))
    log.Error(message, tags...)
    _ = log.Sync()
}

func getOutputLogs() string {
    output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
    if output == "" {
        return "stdout"
    }
    return output
}

func getLevelLog() zapcore.Level {
    switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
    case "info":
        return zapcore.InfoLevel
    case "error":
        return zapcore.ErrorLevel
    case "debug":
        return zapcore.DebugLevel
    default:
        return zapcore.InfoLevel
    }
}
