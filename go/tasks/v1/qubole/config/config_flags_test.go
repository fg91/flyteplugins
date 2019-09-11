// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_quboleTokenPath", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("quboleTokenPath"); err == nil {
				assert.Equal(t, string(defaultConfig.QuboleTokenPath), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("quboleTokenPath", testValue)
			if vString, err := cmdFlags.GetString("quboleTokenPath"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.QuboleTokenPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_resourceManagerType", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("resourceManagerType"); err == nil {
				assert.Equal(t, string(defaultConfig.ResourceManagerType), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("resourceManagerType", testValue)
			if vString, err := cmdFlags.GetString("resourceManagerType"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ResourceManagerType)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redisHostPath", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("redisHostPath"); err == nil {
				assert.Equal(t, string(defaultConfig.RedisHostPath), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redisHostPath", testValue)
			if vString, err := cmdFlags.GetString("redisHostPath"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.RedisHostPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redisHostKey", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("redisHostKey"); err == nil {
				assert.Equal(t, string(defaultConfig.RedisHostKey), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redisHostKey", testValue)
			if vString, err := cmdFlags.GetString("redisHostKey"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.RedisHostKey)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redisMaxRetries", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("redisMaxRetries"); err == nil {
				assert.Equal(t, int(defaultConfig.RedisMaxRetries), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redisMaxRetries", testValue)
			if vInt, err := cmdFlags.GetInt("redisMaxRetries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.RedisMaxRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_quboleLimit", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("quboleLimit"); err == nil {
				assert.Equal(t, int(defaultConfig.QuboleLimit), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("quboleLimit", testValue)
			if vInt, err := cmdFlags.GetInt("quboleLimit"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.QuboleLimit)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_lruCacheSize", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("lruCacheSize"); err == nil {
				assert.Equal(t, int(defaultConfig.LruCacheSize), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("lruCacheSize", testValue)
			if vInt, err := cmdFlags.GetInt("lruCacheSize"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.LruCacheSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_lookasideBufferPrefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("lookasideBufferPrefix"); err == nil {
				assert.Equal(t, string(defaultConfig.LookasideBufferPrefix), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("lookasideBufferPrefix", testValue)
			if vString, err := cmdFlags.GetString("lookasideBufferPrefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LookasideBufferPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_lookasideExpirySeconds", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("lookasideExpirySeconds"); err == nil {
				assert.Equal(t, string(defaultConfig.LookasideExpirySeconds.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.LookasideExpirySeconds.String()

			cmdFlags.Set("lookasideExpirySeconds", testValue)
			if vString, err := cmdFlags.GetString("lookasideExpirySeconds"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LookasideExpirySeconds)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
