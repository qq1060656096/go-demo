package biz

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

const (
	bodyTypeIsNone = "none"
	bodyTypeIsRaw = "raw"
	bodyTypeIsFormData = "formData" //multipart/form-data
	bodyTypeIsFormParams = "urlencoded"// urlencoded
)

type Body struct {
	Type string
	Raw  string
	data interface{}
}

func BodyTypeIsNone(body *Body) bool {
	if body.Type == bodyTypeIsNone {
		return true
	}
	return false
}

func BodyTypeIsRaw(body *Body) bool {
	if body.Type == bodyTypeIsRaw {
		return true
	}
	return false
}

func BodyTypeIsFormData(body *Body) bool {
	if body.Type == bodyTypeIsFormData {
		return true
	}
	return false
}

func BodyTypeIsFormParams(body *Body) bool {
	if body.Type == bodyTypeIsFormParams {
		return true
	}
	return false
}
func (b *Body) GetRaw() error {
	if !BodyTypeIsRaw(b) {
		return nil
	}
	return nil
}

func (b *Body) GetFormData() error {
	if !BodyTypeIsFormData(b) {
		return nil
	}
	var fb bytes.Buffer
	w := multipart.NewWriter(&fb)
	l := strings.Split(b.Raw, "\n")
	filePrefix := "file:"
	for _, v := range l {
		isFile := false
		var rv []string
		// 不是文件
		if !strings.HasPrefix(v, filePrefix) {
			isFile = true
			v = strings.TrimPrefix(v, filePrefix)
		}
		rv = strings.SplitN(v, ":", 2)
		if len(rv) != 2 {
			return errors.New(fmt.Sprintf("字段解析错误: %s", rv))
		}
		name := rv[0]
		value := rv[1]
		if !isFile {
			w.WriteField(name, value)
			continue
		}
		f, err := os.Open(value)
		if err != nil {
			return errors.New(fmt.Sprintf("file类型字段解析错误: %s %v", rv, err))
		}
		// 会defer泄露，但是form 数据量小，就先这样写
		defer f.Close()
		_, err = w.CreateFormFile(name, f.Name())
		if err != nil {
			return errors.New(fmt.Sprintf("file类型字段解析错误: %s %v", rv, err))
		}
	}
	b.data = fb
	return nil
}

func (b *Body) GetFormParams() error {
	if !BodyTypeIsFormParams(b) {
		return nil
	}
	l := strings.Split(b.Raw, "\n")
	values := make(url.Values)
	for _, v := range l {
		var rv []string
		rv = strings.SplitN(v, ":", 2)
		if len(rv) != 2 {
			return errors.New(fmt.Sprintf("字段解析错误: %s", rv))
		}
		name := rv[0]
		value := rv[1]
		values.Add(name, value)
	}
	b.data = values
	return nil
}

func (b *Body) GetData() interface{} {
	b.GetRaw()
	b.GetFormData()
	b.GetFormParams()
	return b.data
}