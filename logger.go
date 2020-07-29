package zaplogger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(c *zapConfig) *zap.Logger {
	syncer := zapcore.AddSync(os.Stdout)
	if c.sample {
		c.opts = append(c.opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSamplerWithOptions(core, time.Second, 100, 100)
		}))
	}
	c.opts = append(c.opts, zap.AddCallerSkip(1), zap.ErrorOutput(syncer))
	return zap.New(zapcore.NewCore(c.encoder, syncer, c.level)).WithOptions(c.opts...)
}
