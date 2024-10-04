package utils

import (
	"base-api/constants"
	"fmt"
	"strings"
)

func ErrDuplicate(m string) error {
	errs := strings.Split(m, constants.PGDuplicateConstraint)
	s := trimQuote(errs[1])
	return fmt.Errorf(fmt.Sprintf(constants.ErrDuplicate, s))
}

func ErrHttpClient(m string) error {
	return fmt.Errorf(fmt.Sprintf(constants.ErrHttpClient, m))
}

func ErrQueryParamsRequired(m string) error {
	return fmt.Errorf(fmt.Sprintf(constants.ErrQueryParamsRequired, m))
}

func ErrIncompleteProfile(m string) error {
	return fmt.Errorf(fmt.Sprintf(constants.ErrIncompleteProfile, m))
}
