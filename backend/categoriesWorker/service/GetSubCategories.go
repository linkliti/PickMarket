package service

import (
	"log/slog"
	"protos/parser"

	"google.golang.org/grpc/codes"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

func (c *CategoryService) GetSubCategories(req *parser.SubCategoriesRequest, srv parser.CategoryParser_GetSubCategoriesServer) error {
	// Get subcategories from the database
	categories, err := c.db.DBGetCategoryChildren(req.CategoryUrl, req.Market)
	if err != nil {
		errText := "failed to get subcategories from database"
		slog.Error(errText, "err", err)
		return sendErrorStatus_GetSubCategories(srv, errText)
	}
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

func sendErrorStatus_GetSubCategories(srv parser.CategoryParser_GetSubCategoriesServer, errText string) error {
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
