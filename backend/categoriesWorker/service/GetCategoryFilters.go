package service

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"protos/parser"
)

func (c *CategoryService) GetCategoryFilters(req *parser.FiltersRequest, srv parser.CategoryParser_GetCategoryFiltersServer) error {
	// Try to get filters from the database
	filters, err := c.db.DBGetFilters(req.CategoryUrl)
	if err != nil {
		// If it fails, get them from the parser
		stream, err := c.parsClient.GetCategoryFilters(context.Background(), req)
		if err != nil {
			slog.Error("failed to get filters from parser", "err", err)
			return err
		}
		// Slice to hold the pointers to filters for saving to the database
		var filtersToSave []*parser.Filter
		// Iterate over the stream
		for {
			filterResponse, err := stream.Recv()
			if err == io.EOF {
				// Save the filters to the database after receiving all filters from the stream
				go func(filtersToSave []*parser.Filter) {
					if err := c.db.DBSaveFilters(filtersToSave, req.CategoryUrl); err != nil {
						slog.Error("failed to save filters to database", "err", err)
					}
				}(filtersToSave)
				break
			}
			if err != nil {
				slog.Error("failed to receive filter from stream", "err", err)
				return err
			}
			// Use a type assertion to get the Filter from the FilterResponse
			if filter, ok := filterResponse.Message.(*parser.FilterResponse_Filter); ok {
				// Create a new FilterResponse to send to the caller
				resp := &parser.FilterResponse{
					Message: filter,
				}
				// Send the FilterResponse to the caller
				if err := srv.Send(resp); err != nil {
					slog.Error("failed to send filter to caller", "err", err)
					return err
				}
				// Add the pointer to the filter to the slice
				filtersToSave = append(filtersToSave, filter.Filter)
			} else {
				slog.Error("received a non-Filter message")
				return errors.New("received a non-Filter message")
			}
		}
	} else {
		// If getting filters from the database succeeds, send them to the caller
		for _, filter := range filters {
			resp := &parser.FilterResponse{
				Message: &parser.FilterResponse_Filter{
					Filter: filter,
				},
			}
			if err := srv.Send(resp); err != nil {
				slog.Error("failed to send filter to caller", "err", err)
				return err
			}
		}
	}
	return nil
}
