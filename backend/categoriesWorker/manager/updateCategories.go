package manager

import (
	"categoriesWorker/db"
	handlerservice "categoriesWorker/handlerService"
	"context"
	"errors"
	"io"
	"log/slog"
	"protos/parser"
)

func UpdateMarketCategories() error {
	// Connections
	client, err := handlerservice.ConnectToParser()
	if err != nil {
		slog.Error("failed to connect to parser", err)
		return err
	}
	d, err := db.NewDBConnectionManager()
	if err != nil {
		slog.Error("failed to connect to database", err)
		return err
	}

	// Get reqs with empty ParseDate
	reqs, err := d.DBGetCategoriesWithEmptyParseDate()
	if err != nil {
		slog.Error("failed to get categories from database", err)
		return err
	}

	// Iterate over the categories
	for _, req := range reqs {
		// Send GetSubCategories to parser
		stream, err := client.GetSubCategories(context.Background(), req)
		if err != nil {
			slog.Error("failed to get subcategories from parser", err)
			return err
		}

		// Iterate over the stream
		for {
			categoryResponse, err := stream.Recv()
			if err == io.EOF {
				// Update parent category parse time to NOW() after receiving all subcategories from the stream
				if err := d.DBSetCategoryUpdateTime(req.CategoryUrl); err != nil {
					slog.Error("failed to update category parse date in database", err)
				}
				break
			}
			if err != nil {
				slog.Error("failed to receive category from stream", err)
				return err
			}

			// Use a type assertion to get the Category from the CategoryResponse
			if subCategory, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
				// Insert new subcategory from stream
				go func(cat *parser.Category, market parser.Markets) {
					if err := d.DBSaveCategory(cat, market); err != nil {
						slog.Error("failed to save subcategory to database", err)
					}
				}(subCategory.Category, req.Market)
			} else {
				slog.Error("received a non-Category message")
				return errors.New("received a non-Category message")
			}
		}
	}

	return nil
}
