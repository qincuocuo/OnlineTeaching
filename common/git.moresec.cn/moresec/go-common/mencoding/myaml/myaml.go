package myaml

import (
	"gopkg.in/yaml.v3"

	"git.moresec.cn/moresec/go-common/internal/json"
	"git.moresec.cn/moresec/go-common/mconv"
	"git.moresec.cn/moresec/go-common/merrors/merror"
)

func Encode(value interface{}) (out []byte, err error) {
	if out, err = yaml.Marshal(value); err != nil {
		err = merror.Wrap(err, `encode value to yaml failed`)
	}
	return
}

func Decode(value []byte) (interface{}, error) {
	var (
		result map[string]interface{}
		err    error
	)
	if err = yaml.Unmarshal(value, &result); err != nil {
		err = merror.Wrap(err, `decode yaml failed`)
		return nil, err
	}
	return mconv.MapDeep(result), nil
}

func DecodeTo(value []byte, result interface{}) (err error) {
	err = yaml.Unmarshal(value, result)
	if err != nil {
		err = merror.Wrap(err, `encode yaml to value failed`)
	}
	return
}

func ToJson(value []byte) (out []byte, err error) {
	var (
		result interface{}
	)
	if result, err = Decode(value); err != nil {
		return nil, err
	} else {
		if out, err = json.Marshal(result); err != nil {
			err = merror.Wrap(err, `convert yaml to json failed`)
		}
		return out, err
	}
}
