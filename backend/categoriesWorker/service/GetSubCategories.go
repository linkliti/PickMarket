package service

import (
	"log/slog"
	"protos/parser"
)

func (c *CategoryService) GetSubCategories(req *parser.SubCategoriesRequest, srv parser.CategoryParser_GetSubCategoriesServer) error {
	// Get subcategories from the database
	categories, err := c.db.DBGetCategoryChildren(req.CategoryUrl, req.Market)
	if err != nil {
		slog.Error("failed to get subcategories from database", "err", err)
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
			slog.Error("failed to send category to caller", "err", err)
			return err
		}
	}
	return nil
}
