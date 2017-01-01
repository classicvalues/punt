package punt

import (
	"encoding/json"
	"github.com/jeromer/syslogparser"
)

type Transformer interface {
	Transform(parts syslogparser.LogParts) (map[string]interface{}, error)
}

// Doesn't perform any transformation or parsing on the syslog structure
type DirectTransformer struct{}

func NewDirectTransformer(config map[string]interface{}) *DirectTransformer {
	return &DirectTransformer{}
}

func (b *DirectTransformer) Transform(parts syslogparser.LogParts) (map[string]interface{}, error) {
	return parts, nil
}

// Parses the log line as JSON and merges it into the syslog structure
type UnpackMergeTransformer struct{}

func NewUnpackMergeTransformer(config map[string]interface{}) *UnpackMergeTransformer {
	return &UnpackMergeTransformer{}
}

func (u *UnpackMergeTransformer) Transform(parts syslogparser.LogParts) (map[string]interface{}, error) {
	err := json.Unmarshal([]byte(parts["content"].(string)), parts)
	return parts, err
}

// Parses the log line as JSON and uses it as the core structure (ignoring syslog data)
type UnpackTakeTransformer struct{}

func NewUnpackTakeTransformer(config map[string]interface{}) *UnpackTakeTransformer {
	return &UnpackTakeTransformer{}
}

func (u *UnpackTakeTransformer) Transform(parts syslogparser.LogParts) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(parts["content"].(string)), &data)
	return data, err
}
