package regutils

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

const (
	base      = `SOFTWARE`
	src       = `Firedrill\Config`
	EXPAND_SZ = 2
	BINARY    = 3
	DWORD     = 4
	QWORD     = 11
)

func getFiredrillRegKey() (registry.Key, error) {
	appkey, err := registry.OpenKey(registry.CURRENT_USER, base, registry.CREATE_SUB_KEY)
	if err != nil {
		return 0, err
	}
	defer appkey.Close()

	newKey, _, err := registry.CreateKey(appkey, src, registry.ALL_ACCESS)
	if err != nil {
		return 0, err
	}
	return newKey, nil
}

func RemoveFiredrillRegistry() error {
	appkey, err := registry.OpenKey(registry.LOCAL_MACHINE, base, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer appkey.Close()
	return registry.DeleteKey(appkey, src)
}

func GetConfigFromRegistry(Name string) (interface{}, error) {
	var (
		newKey registry.Key
		err    error
		value  interface{}
	)
	newKey, err = getFiredrillRegKey()
	if err != nil {
		panic(err)
	}
	defer newKey.Close()

	var buf []byte
	_, valtype, readErr := newKey.GetValue(Name, buf)
	if readErr != nil {
		return nil, readErr
	}

	switch valtype {
	case EXPAND_SZ:
		value, _, _ = newKey.GetStringValue(Name)
	case BINARY:
		value, _, _ = newKey.GetBinaryValue(Name)
	case DWORD, QWORD:
		value, _, _ = newKey.GetIntegerValue(Name)
	default:
		return nil, errors.New("unknown registry key data type")
	}
	return value, nil
}

func SaveConfigToRegistry(Name string, Data interface{}) error {
	var (
		newKey registry.Key
		err    error
	)
	newKey, err = getFiredrillRegKey()
	if err != nil {
		panic(err)
	}
	defer newKey.Close()

	switch v := Data.(type) {
	case string:
		err = newKey.SetExpandStringValue(Name, v)
	case uint32:
		err = newKey.SetDWordValue(Name, v)
	case uint64:
		err = newKey.SetQWordValue(Name, v)
	case []byte:
		err = newKey.SetBinaryValue(Name, v)
	default:
		return errors.New("unknown data type sent to registry")
	}
	return err
}
