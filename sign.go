package uniSign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
)

type Tool struct {
	AuthSecretKey string
}

func (tool Tool) GetLoginSignature(params map[string]any, nonce string, timestamp int64) string {
	paramsStr := tool.getParamsString(params)
	mac := hmac.New(sha256.New, []byte(tool.AuthSecretKey+nonce))
	mac.Write([]byte(strconv.Itoa(int(timestamp)) + paramsStr))
	hexSignature := mac.Sum(nil)
	signature := hex.EncodeToString(hexSignature)
	return strings.ToUpper(signature)
}

func (tool Tool) GetS2SSignature(params map[string]any, timestamp int64) string {
	paramsStr := tool.getParamsString(params)
	mac := hmac.New(sha256.New, []byte(tool.AuthSecretKey))
	mac.Write([]byte(strconv.Itoa(int(timestamp)) + "\n" + paramsStr))
	hexSignature := mac.Sum(nil)
	signature := hex.EncodeToString(hexSignature)
	return signature
}

func (tool Tool) getParamsString(params map[string]any) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramsStr = ""
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		v := params[k]
		var convertStr string
		switch t := v.(type) {
		case int:
			convertStr = strconv.Itoa(t)
		case int64:
			convertStr = strconv.FormatInt(t, 10)
		case int32:
			convertStr = strconv.FormatInt(int64(t), 10)
		case int16:
			convertStr = strconv.FormatInt(int64(t), 10)
		case int8:
			convertStr = strconv.FormatInt(int64(t), 10)
		case uint:
			convertStr = strconv.FormatUint(uint64(t), 10)
		case uint64:
			convertStr = strconv.FormatUint(t, 10)
		case uint32:
			convertStr = strconv.FormatUint(uint64(t), 10)
		case uint16:
			convertStr = strconv.FormatUint(uint64(t), 10)
		case uint8:
			convertStr = strconv.FormatUint(uint64(t), 10)
		case float32:
			convertStr = strconv.FormatFloat(float64(t), 'f', -1, 32)
		case float64:
			convertStr = strconv.FormatFloat(t, 'f', -1, 64)
		case bool:
			convertStr = strconv.FormatBool(t)
		case string:
			convertStr = t
		}
		if len(convertStr) > 0 {
			if i != 0 {
				paramsStr += "&"
			}
			paramsStr += k
			paramsStr += "="
			paramsStr += convertStr
		}
	}
	return paramsStr
}
