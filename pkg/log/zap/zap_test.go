/**
  @author: $(USER)
  @data:$(DATE)
  @note:
**/
package zap

import (
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	logger := NewZapLog("demo", false)
	logger.Open()
	defer logger.Close()
	logger.Debug("测试", zap.Error(errors.New("demo errors")))
	logger.Info("测试", zap.Error(errors.New("demo errors")))
	logger.Warn("测试", zap.Error(errors.New("demo errors")))
	logger.Error("测试", zap.Error(errors.New("demo errors")))
}

func TestZapDebug(t *testing.T) {
	logger := NewZapLog("demo", true)
	logger.Open()
	defer logger.Close()
	logger.Debug("测试", zap.Error(errors.New("debug errors")))
	logger.Info("测试", zap.Error(errors.New("debug errors")))
	logger.Warn("测试", zap.Error(errors.New("debug errors")))
	logger.Error("测试", zap.Error(errors.New("debug errors")))
}
