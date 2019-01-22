package bufr

import codes "github.com/mtfelian/go-eccodes"

// GetInt ...
func GetInt(msg codes.Message, key string) int { v, _ := msg.GetLong(key); return int(v) }

// GetInt64 ...
func GetInt64(msg codes.Message, key string) int64 { v, _ := msg.GetLong(key); return v }

// GetFloat64 ...
func GetFloat64(msg codes.Message, key string) float64 { v, _ := msg.GetDouble(key); return v }

// GetString ...
func GetString(msg codes.Message, key string) string { v, _ := msg.GetString(key); return v }
