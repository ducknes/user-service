package domain

import "errors"

var (
	NoDocuments        = errors.New("no documents found")
	NoDocumentAffected = errors.New("no documents affected")
)
