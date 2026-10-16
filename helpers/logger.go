package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// UpstashWriter implements io.Writer and pushes logs to an Upstash Redis list via the REST API
type UpstashWriter struct {
	url    string
	token  string
	key    string
	client *http.Client
}

// run a single redis command against the Upstash REST API
func (u *UpstashWriter) do(args ...string) error {
	body, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", u.url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+u.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res struct {
		Error string `json:"error"`
	}
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != http.StatusOK || res.Error != "" {
		return fmt.Errorf("upstash returned %d: %s", resp.StatusCode, res.Error)
	}
	return nil
}

func (u *UpstashWriter) Write(p []byte) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("UpstashWriter Write Crashed: ", r)
		}
	}()
	if err := u.do("LPUSH", u.key, strings.TrimRight(string(p), "\n")); err != nil {
		fmt.Println("UpstashWriter Write failed: ", err)
		return 0, err
	}
	return len(p), nil
}

func newCore(writer zapcore.WriteSyncer) zapcore.Core {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)
	return zapcore.NewCore(jsonEncoder, writer, zapcore.InfoLevel)
}

// InitLogger logs to stdout and to the Upstash list `key`
func InitLogger(restURL, token, key string) (*zap.Logger, error) {
	if restURL == "" || token == "" {
		return nil, errors.New("UPSTASH_REDIS_REST_URL or UPSTASH_REDIS_REST_TOKEN is not set")
	}
	writer := &UpstashWriter{
		url:    strings.TrimRight(restURL, "/"),
		token:  token,
		key:    key,
		client: &http.Client{Timeout: 10 * time.Second},
	}
	if err := writer.do("PING"); err != nil {
		return nil, err
	}

	// Write to both Upstash and stdout so docker logs works
	upstashSync := zapcore.AddSync(writer)
	stdoutSync := zapcore.AddSync(os.Stdout)
	multiSync := zapcore.NewMultiWriteSyncer(upstashSync, stdoutSync)

	logger := zap.New(newCore(multiSync), zap.AddCaller())
	Logger = logger
	return logger, nil
}
