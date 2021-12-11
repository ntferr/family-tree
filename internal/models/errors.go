package models

import "errors"

var (
	ErrOnConnectDB = errors.New("Error on connect to Database")

	ErrToGetParamRelationshipType = errors.New("Error on get relationship type on params")

	ErrToGetParamPerson = errors.New("Need insert an id or name by query param")
)
