package handlerservice

import (
	"categoriesWorker/db"
	"log/slog"
	"protos/parser"
)

func (c *CategoryService) GetSubCategories(req *parser.SubCategoriesRequest, srv parser.CategoryParser_GetSubCategoriesServer) error {
	// Connections
	d, err := db.NewDBConnection(req.Market)
	if err != nil {
		slog.Error("failed to connect to database", err)
		return err
	}

	// Get subcategories from the database
	categories, err := d.DBGetCategoryChildren(req.CategoryUrl)
	if err != nil {
		slog.Error("failed to get subcategories from database", err)
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
