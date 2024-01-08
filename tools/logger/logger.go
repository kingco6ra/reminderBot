package logger

// import "context"

// func initLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
// 	loggingLevel := zap.InfoLevel
// 	if cfg.Project.Debug {
// 		loggingLevel = zap.DebugLevel
// 	}

// 	consoleCore := zapcore.NewCore(
// 		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
// 		os.Stderr,
// 		zap.NewAtomicLevelAt(loggingLevel),
// 	)

// 	gelfCore, err := gelf.NewCore(
// 		gelf.Addr(cfg.Telemetry.GraylogPath),
// 		gelf.Level(loggingLevel),
// 	)
// 	if err != nil {
// 		logger.FatalKV(ctx, "initLogger() error", "err", err)
// 	}

// 	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

// 	sugaredLogger := notSugaredLogger.Sugar()
// 	logger.SetLogger(sugaredLogger.With(
// 		"service", cfg.Project.Name,
// 	))

// 	return func() {
// 		err := notSugaredLogger.Sync()
// 		if err != nil {
// 			logger.FatalKV(ctx, "initLogger() error", "err", err)
// 		}
// 	}
// }
