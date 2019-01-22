package bufr

import codes "github.com/mtfelian/go-eccodes"

func getInt(msg codes.Message, key string) int         { v, _ := msg.GetLong(key); return int(v) }
func getInt64(msg codes.Message, key string) int64     { v, _ := msg.GetLong(key); return v }
func getFloat64(msg codes.Message, key string) float64 { v, _ := msg.GetDouble(key); return v }
func getString(msg codes.Message, key string) string   { v, _ := msg.GetString(key); return v }
