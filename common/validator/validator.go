package validator

import (
	"fmt"
	"github.com/go-playground/locales/en_US"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"net/http"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(en_US.New())
	trans, _ := uni.GetTranslator("en_US")

	err := enTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), http.StatusUnprocessableEntity
		}
	}
	return "", 0
}
