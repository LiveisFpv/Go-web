package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var ErrNotStruct = errors.New("wrong argument given, should be a struct")
var ErrInvalidValidatorSyntax = errors.New("invalid validator syntax")
var ErrInvalidLength = errors.New("invalid length")
var ErrValidateForUnexportedFields = errors.New("validation for unexported field is not allowed")

type ValidationError struct {
	Err error
}

type ValidationErrors []ValidationError

func New() ValidationErrors {
	return make(ValidationErrors, 0)
}
func (v ValidationErrors) Error() string {
	var s string
	s = ""
	for _, val := range v {
		s += val.Err.Error() + "\n"
	}
	return s
}
func valid_len(le string, el reflect.Value) error {
	le1, err := strconv.Atoi(le)
	if err != nil {
		return errors.Wrap(ErrInvalidValidatorSyntax, "invalid length validator syntax")
	} else {
		if el.Kind() == reflect.Int {
			return errors.Wrapf(ErrInvalidValidatorSyntax, "invalid length validator syntax")
		} else if el.Kind() == reflect.String {
			if el.Len() != le1 {
				return errors.Wrapf(ErrInvalidLength, "length of field %s should be equal to %d", el.Type(), le1)
			}
		}
	}
	return nil
}
func valid_min(min string, el reflect.Value) error {
	min1, err := strconv.Atoi(min)
	if err != nil {
		return errors.Wrap(ErrInvalidValidatorSyntax, "invalid length validator syntax")
	} else {
		if el.Kind() == reflect.Int {
			if int(el.Int()) < min1 {
				return errors.Wrapf(ErrInvalidLength, "length of field %s should be greater than or equal to %d", el.Type(), min1)
			}
		} else if el.Kind() == reflect.String {
			if el.Len() < min1 {
				return errors.Wrapf(ErrInvalidLength, "length of field %s should be greater than or equal to %d", el.Type(), min1)
			}
		}
	}
	return nil
}
func valid_max(max string, el reflect.Value) error {
	max1, err := strconv.Atoi(max)
	if err != nil {
		return errors.Wrap(ErrInvalidValidatorSyntax, "invalid length validator syntax")
	} else {
		if el.Kind() == reflect.Int {
			if int(el.Int()) > max1 {
				return errors.Wrapf(ErrInvalidLength, "length of field %s should be less than or equal to %d", el.Type(), max1)
			}
		} else if el.Kind() == reflect.String {
			if el.Len() > max1 {
				return errors.Wrapf(ErrInvalidLength, "length of field %s should be less than or equal to %d", el.Type(), max1)
			}
		}
	}
	return nil
}
func valid_in(in string, el reflect.Value) error {
	inSlice := strings.Split(in, ",")
	found := false
	for _, v := range inSlice {
		if el.Kind() == reflect.String {
			if v == string(el.String()) {
				found = true
				break
			}
		} else if el.Kind() == reflect.Int {
			val, _ := strconv.Atoi(v)
			if int(el.Int()) == val {
				found = true
				break
			}
		}
	}
	if !found {
		return fmt.Errorf("value of field %s should be one of %s", el.Type(), in)
	}
	return nil
}
func parse(tag, val string) string {
	if strings.Contains(val, tag) {
		return strings.Split(val, tag)[1]
	}
	return ""
}
func Validate(v any) error {
	ValidationErrors := New()
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Struct {
		fmt.Println(typ.Kind())
		return errors.Wrap(ErrNotStruct, "invalid type")
	} else {
		for i := 0; i < val.NumField(); i++ {
			el := val.Field(i)
			for _, tag := range strings.Split(val.Type().Field(i).Tag.Get("validate"), " ") {
				le := parse("len:", tag)
				min := parse("min:", tag)
				max := parse("max:", tag)
				in := parse("in:", tag)
				if len(le+min+max+in) == 0 {
					continue
				} else {
					if le != "" {
						err := valid_len(le, el)
						if err != nil {
							ValidationErrors = append(ValidationErrors, ValidationError{err})
							continue
						}
					} else if min != "" {
						err := valid_min(min, el)
						if err != nil {
							ValidationErrors = append(ValidationErrors, ValidationError{err})
							continue
						}
					} else if max != "" {
						err := valid_max(max, el)
						if err != nil {
							ValidationErrors = append(ValidationErrors, ValidationError{err})
							continue
						}
					} else if in != "" {
						err := valid_in(in, el)
						if err != nil {
							ValidationErrors = append(ValidationErrors, ValidationError{err})
							continue
						}
					}
				}
				fmt.Println(el, le, min, max, in)
			}
			if typ.Field(i).PkgPath != "" && val.Field(i).IsZero() {
				ValidationErrors = append(ValidationErrors, ValidationError{
					Err: errors.Wrap(ErrValidateForUnexportedFields, "validation for unexported field is not allowed"),
				})
			}
		}
		if len(ValidationErrors) > 0 {
			return ValidationErrors
		}
	}
	return nil
}
