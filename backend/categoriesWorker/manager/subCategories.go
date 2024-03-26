package manager

import (
	"context"
	"io"
	"log/slog"
	"protos/parser"
	"sync"
)

func (m *Manager) UpdateAllSubCategories() error {
	// Get markets that have categories without parseDate
	markets, err := m.db.DBGetMarketsWithCategoriesWithoutParseDate()
	if err != nil {
		return err
	}
	// Iterate over the markets
	for _, market := range markets {
		slog.Info("updating subcategories for market", "market", market)
		// Continuously call UpdateSubCategories until there are no categories left without a parseDate
		for {
			// Call UpdateSubCategories for the market
			err := m.UpdateSubCategories(market)
			if err != nil {
				return err
			}
			// Check if there are any categories left without a parseDate
			categories, err := m.db.DBGetCategoriesWithoutParseDate(market)
			if err != nil {
				return err
			}
			// If there are no categories left without a parseDate, break the loop
			if len(categories) == 0 {
				break
			}
		}
	}
	return nil
}

func (m *Manager) UpdateSubCategories(market parser.Markets) error {
	// Get categories without parseDate for the current market
	categories, err := m.db.DBGetCategoriesWithoutParseDate(market)
	if err != nil {
		return err
	}
	// Create a buffered channel for categories
	categoryChan := make(chan string, len(categories))
	// Create a worker pool
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for categoryUrl := range categoryChan {
				// Create a SubCategoriesRequest for the category
				req := &parser.SubCategoriesRequest{Market: market, CategoryUrl: categoryUrl}
				// Call GetSubCategories for the category
				stream, err := m.parsClient.GetSubCategories(context.Background(), req)
				if err != nil {
					slog.Error("failed to get subcategories from parser", err)
					continue
				}
				// Iterate over the stream
				for {
					categoryResponse, err := stream.Recv()
					if err == io.EOF {
						// Set parseDate to NOW() for the category after receiving all subcategories from the stream
						if err := m.db.DBSetCategoryParseDate(categoryUrl); err != nil {
							slog.Error("failed to set parseDate for category", err)
						}
						break
					}
					if err != nil {
						slog.Error("failed to receive category from stream", err)
						break
					}
					// Use a type assertion to get the Category from the CategoryResponse
					if category, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
						// Save the subcategory in the database
						if err := m.db.DBSaveCategory(category.Category, market); err != nil {
							slog.Error("failed to save category to database", err)
						}
					} else {
						slog.Error("received a non-Category message")
					}
				}
			}
		}()
	}
	// Send categories to the channel
	for _, categoryUrl := range categories {
		categoryChan <- categoryUrl
	}
	// Close the channel and wait for all workers to finish
	close(categoryChan)
	wg.Wait()
	return nil
}
