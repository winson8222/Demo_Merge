package main

import (
	"context"
	postersapi "posterservice/kitex_gen/postersapi"
)

// PosterServiceImpl implements the last service interface defined in the IDL.
type PosterServiceImpl struct{}

// Getuniqueusernames implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames(ctx context.Context) (resp *postersapi.Response, err error) {
	response := &postersapi.Response{
		Posterslist: []string{
			"Michael",
			"Walter",
			"Howard",
			"Charles",
			"Jane",
			"Lyle",
			"James",
			"Wendy",
			"Todd",
			"Jack",
			"Elliot",
		},
	}

	return response, nil
}
