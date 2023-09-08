package main

import (
	"context"
	viewersapi "server/kitex_gen/viewersapi"
)

// ViewerServiceImpl implements the last service interface defined in the IDL.
type ViewerServiceImpl struct{}

// Getuniqueviewernames implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	response := &viewersapi.Response{
		Viewerslist: []string{
			"Steven",
			"Jane",
			"Elliot",
			"Lyle",
			"Lydia",
			"Bill",
			"Gale",
			"Wendy",
			"Todd",
			"Jack",
		},
	}
	return response, nil
}
