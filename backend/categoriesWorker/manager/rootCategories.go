package manager

import (
	"context"
	"io"
	"log/slog"
	"protos/parser"
	"sync"
)

func (m *Manager) UpdateRootCategories() error {
	// Get markets without parseDate
	markets, err := m.db.DBGetMarketsWithoutParseDate()
	if err != nil {
		return err
	}
	var marketStrings []string
	for _, market := range markets {
		marketStrings = append(marketStrings, market.String())
	}
	slog.Debug("empty markets from db", "markets", marketStrings)
	// Create a buffered channel for markets
	marketChan := make(chan parser.Markets, len(markets))
	// Create a worker pool
	var wg sync.WaitGroup
	for i := 0; i < m.workpoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for market := range marketChan {
				slog.Info("updating root categories for market", "market", market.String())
				// Create a RootCategoriesRequest for the market
				req := &parser.RootCategoriesRequest{Market: market}
				// Call GetRootCategories for the market
				stream, err := m.parsClient.GetRootCategories(context.Background(), req)
				if err != nil {
					slog.Error("failed to get root categories from parser", "err", err)
					continue
				}
				// Iterate over the stream
				for {
					categoryResponse, err := stream.Recv()
					if err == io.EOF {
						// Set parseDate to NOW() for the market after receiving all categories from the stream
						if err := m.db.DBSetMarketParseDate(market); err != nil {
							slog.Error("failed to set parseDate for market", "err", err)
						}
						slog.Info("updated market", "market", market.String())
						break
					}
					if err != nil {
						slog.Error("failed to receive category from stream", "err", err)
						break
					}
					// Use a type assertion to get the Category from the CategoryResponse
					if category, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
						// Save the category in the database
						if err := m.db.DBSaveCategory(category.Category, market); err != nil {
							slog.Error("failed to save category to database", "err", err)
						}
					} else {
						slog.Error("received a non-Category message")
						if status, ok := categoryResponse.Message.(*parser.CategoryResponse_Status); ok {
							slog.Error("status", "status", status.Status)
						}
					}
				}
			}
		}()
	}
	// Send markets to the channel
	for _, market := range markets {
		marketChan <- market
	}
	// Close the channel and wait for all workers to finish
	close(marketChan)
	wg.Wait()
	return nil
}
