package logging

import "time"

type Log struct {
  CorrelationID string    `json:"sericy.correlationId"`
  ClassName     string    `json:"sericy.className"`
  TimeStsamp    time.Time `json:"sericy.timestamp"`
  Message       string    `json:"sericy.message"`
  Exception     string    `json:"sericy.exception"`
  LogLevel      string    `json:"sericy.logLevel"`
  ExecutionTime int64     `json:"sericy.executionTime"`
}
