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

func UpdateMarketRootCategories() error {
	// Connections
	client := handlerservice.ConnectToParser()
	d, err := db.NewDBConnectionManager()
	if err != nil {
		slog.Error("failed to connect to database", err)
		return err
	}

	// Get marketNames with empty marketParseDate
	markets, err := d.DBGetMarketsWithEmptyUpdateTime()
	if err != nil {
		slog.Error("failed to get markets from database", err)
		return err
	}

	// Iterate over the markets
	for _, market := range markets {
		// Prepare the request
		req := &parser.RootCategoriesRequest{
			Market: market,
		}

		// Send GetRootCategories to parser
		stream, err := client.GetRootCategories(context.Background(), req)
		if err != nil {
			slog.Error("failed to get root categories from parser", err)
			return err
		}

		// Iterate over the stream
		for {
			categoryResponse, err := stream.Recv()
			if err == io.EOF {
				// Update marketParseDate to NOW() after receiving all categories from the stream
				if err := d.DBUpdateMarketUpdateTime(market); err != nil {
					slog.Error("failed to update market parse date in database", err)
				}
				break
			}
			if err != nil {
				slog.Error("failed to receive category from stream", err)
				return err
			}

			// Use a type assertion to get the Category from the CategoryResponse
			if category, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
				// Insert new category from stream
				if err := d.DBSaveRootCategory(market, category.Category); err != nil {
					slog.Error("failed to save category to database", err)
					return err
				}
			} else {
				slog.Error("received a non-Category message")
				return errors.New("received a non-Category message")
			}
		}
	}

	return nil
}
