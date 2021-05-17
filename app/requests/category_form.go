package requests

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/category"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:")) //handle other error
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})
}

// 验证表单，返回 errs 长度等于零即通过
func ValidateCategoryForm(data category.Category) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string {
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少2个字",
			"max_cn:分类名称长度不能超过8个字",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options {
		Data: &data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}

	// 4.开始验证
	return govalidator.New(opts).ValidateStruct()
}
