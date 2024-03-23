package manager

import (
	"categoriesWorker/db"
	"context"
	"errors"
	"io"
	"protos/parser"
)

type Manager struct {
	parsClient parser.CategoryParserClient
	db         *db.Database
}

func NewManager(parsClient parser.CategoryParserClient, db *db.Database) *Manager {
	return &Manager{parsClient: parsClient, db: db}
}

func (m *Manager) UpdateRootCategories() error {
	// Get markets without parseDate
	markets, err := m.db.DBGetMarketsWithoutParseDate()
	if err != nil {
		return err
	}
	// Iterate over the markets
	for _, market := range markets {
		// Create a RootCategoriesRequest for the market
		req := &parser.RootCategoriesRequest{Market: market}
		// Call GetRootCategories for the market
		stream, err := m.parsClient.GetRootCategories(context.Background(), req)
		if err != nil {
			return err
		}
		// Iterate over the stream
		for {
			categoryResponse, err := stream.Recv()
			if err == io.EOF {
				// Set parseDate to NOW() for the market after receiving all categories from the stream
				if err := m.db.DBSetMarketParseDate(market); err != nil {
					return err
				}
				break
			}
			if err != nil {
				return err
			}
			// Use a type assertion to get the Category from the CategoryResponse
			if category, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
				// Save the category in the database
				if err := m.db.DBSaveCategory(category.Category, market); err != nil {
					return err
				}
			} else {
				return errors.New("received a non-Category message")
			}
		}
	}
	return nil
}

func (m *Manager) UpdateSubCategories() error {
	// Get markets without parseDate
	markets, err := m.db.DBGetMarketsWithoutParseDate()
	if err != nil {
		return err
	}
	// Iterate over the markets
	for _, market := range markets {
		// Get categories without parseDate for the current market
		categories, err := m.db.DBGetCategoriesWithoutParseDate(market)
		if err != nil {
			return err
		}
		// Iterate over the categories
		for _, categoryUrl := range categories {
			// Create a SubCategoriesRequest for the category
			req := &parser.SubCategoriesRequest{Market: market, CategoryUrl: categoryUrl}
			// Call GetSubCategories for the category
			stream, err := m.parsClient.GetSubCategories(context.Background(), req)
			if err != nil {
				return err
			}
			// Iterate over the stream
			for {
				categoryResponse, err := stream.Recv()
				if err == io.EOF {
					// Set parseDate to NOW() for the category after receiving all subcategories from the stream
					if err := m.db.DBSetCategoryParseDate(categoryUrl); err != nil {
						return err
					}
					break
				}
				if err != nil {
					return err
				}
				// Use a type assertion to get the Category from the CategoryResponse
				if category, ok := categoryResponse.Message.(*parser.CategoryResponse_Category); ok {
					// Save the subcategory in the database
					if err := m.db.DBSaveCategory(category.Category, market); err != nil {
						return err
					}
				} else {
					return errors.New("received a non-Category message")
				}
			}
		}
	}
	return nil
}
