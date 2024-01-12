package paramValidator

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var t ut.Translator

func registerTraverse(tag, message string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		return trans.Add(tag, message, true)
	}
}

func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.Field()
	}
	return msg
}

func translateKV(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Param(), fe.Field())
	if err != nil {
		return fe.Field()
	}
	return msg
}

func T(err error) (result string) {
	errs := err.(validator.ValidationErrors)

	for _, v := range errs.Translate(t) {
		result += v
		result += " "
	}
	return result
}

func RespondWithParamError(c *gin.Context, err error) {
	response := &res.BaseResponse{
		Code:    -1,
		Message: T(err),
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

func registerValidator(v *validator.Validate, t ut.Translator) {
	v.RegisterValidation(StringIsMobileTag, stringIsMobile)
	v.RegisterTranslation(StringIsMobileTag, t, registerTraverse(StringIsMobileTag, "need a correctly formatted {0}"), translate)

	v.RegisterValidation(StringNotBothEmptyTag, notBothEmpty)
	v.RegisterTranslation(StringNotBothEmptyTag, t, registerTraverse(StringNotBothEmptyTag, "{0} and {1} can't both be empty"), translateKV)
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		e := en.New()
		uni := ut.New(e, e)
		t, ok = uni.GetTranslator("en")
		if !ok {
			panic("init validate message translator error")
		}
		en_translations.RegisterDefaultTranslations(v, t)
		registerValidator(v, t)
	}
}
