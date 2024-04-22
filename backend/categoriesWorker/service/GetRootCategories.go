package service

import (
	"log/slog"
	"protos/parser"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

func (c *CategoryService) GetRootCategories(req *parser.RootCategoriesRequest, srv parser.CategoryParser_GetRootCategoriesServer) error {
	// Get root categories from the database
	slog.Debug("Incoming GetRootCategories request", "request", req)
	categories, err := c.db.DBGetRootCategoryChildren(req.Market)
	if err != nil {
		errText := "failed to get root categories from database"
		slog.Error(errText, "err", err)
		return sendErrorStatus_GetRootCategories(srv, errText)
	}
	slog.Debug("Sending root categories from database", "request", req)
	// Iterate over the categories and send them to the caller
	for _, category := range categories {
		resp := &parser.CategoryResponse{
			Message: &parser.CategoryResponse_Category{
				Category: category,
			},
		}
		if err := srv.Send(resp); err != nil {
			slog.Error("failed to send category to caller", "err", err)
			return err
		}
	}
	return nil
}

func sendErrorStatus_GetRootCategories(srv parser.CategoryParser_GetRootCategoriesServer, errText string) error {
	resp := &parser.CategoryResponse{
		Message: &parser.CategoryResponse_Status{
			Status: &statuspb.Status{
				Code:    int32(codes.Internal),
				Message: errText,
			},
		},
	}
	return srv.Send(resp)
}
