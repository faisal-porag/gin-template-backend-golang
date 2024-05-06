package utils

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog/log"
)

func ValidateStructData(data interface{}) error {
	validate := validator.New()
	uni := ut.New(en.New())
	translator, _ := uni.GetTranslator("en")

	validatorErr := enTranslations.RegisterDefaultTranslations(validate, translator)
	if validatorErr != nil {
		log.Error().Err(validatorErr)
		return errors.New(validatorErr.Error())
	}

	err := validate.Struct(data)
	if err != nil {
		//for _, err := range err.(validator.ValidationErrors) {
		//	fmt.Printf("Key: '%s' Error: %s\n", err.StructField(), err.Translate(translator))
		//}
		log.Error().Err(err).Msgf(
			"Key: '%s' Error: %s\n",
			err.(validator.ValidationErrors)[0].StructField(),
			err.(validator.ValidationErrors)[0].Translate(translator),
		)
		return errors.New(err.(validator.ValidationErrors)[0].Translate(translator))
	}
	return err
}
