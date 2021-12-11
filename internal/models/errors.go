package models

import "errors"

var (
	ErrOnConnectDB = errors.New("Error on connect to Database")

	ErrToGetParamRelationshipType = errors.New("Error on get relationship type on params")
)
