package service

import (
	"log/slog"
	"protos/parser"
)

func (c *CategoryService) GetRootCategories(req *parser.RootCategoriesRequest, srv parser.CategoryParser_GetRootCategoriesServer) error {
	// Get root categories from the database
	categories, err := c.db.DBGetRootCategoryChildren(req.Market)
	if err != nil {
		slog.Error("failed to get root categories from database", err)
		return err
	}
	// Iterate over the categories and send them to the caller
	for _, category := range categories {
		resp := &parser.CategoryResponse{
			Message: &parser.CategoryResponse_Category{
				Category: category,
			},
		}
		if err := srv.Send(resp); err != nil {
			slog.Error("failed to send category to caller", err)
			return err
		}
	}
	return nil
}
