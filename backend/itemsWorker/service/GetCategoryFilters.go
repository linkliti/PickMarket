package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"protos/parser"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

func (c *ItemsService) GetCategoryFilters(req *parser.FiltersRequest, srv parser.ItemParser_GetCategoryFiltersServer) error {
	// Get filters from DB
	slog.Debug("Incoming GetCategoryFilters request", "request", req)
	filters, err := c.db.DBGetFilters(req.CategoryUrl, req.Market)
	if err != nil {
		errText := "failed to get filters from database"
		slog.Error(errText, "err", err)
		return sendErrorStatus_GetCategoryFilters(srv, errText)
	}
	// Get from parser
	if len(filters) == 0 {
		slog.Debug("Sending filters from parser", "request", req)
		stream, err := c.parsClient.GetCategoryFilters(context.Background(), req)
		if err != nil {
			slog.Error("failed to get filters from parser", "err", err)
			return err
		}
		var filtersToSave []*parser.Filter
		for {
			filterResponse, err := stream.Recv()
			// Save final array to DB
			if err == io.EOF {
				go func(filtersToSave []*parser.Filter) {
					if err := c.db.DBSaveFilters(filtersToSave, req.CategoryUrl, req.Market); err != nil {
						slog.Error("failed to save filters to database", "err", err)
					}
				}(filtersToSave)
				break
			}
			// Failed message
			if err != nil {
				slog.Error("failed to receive filter from stream", "err", err)
				return err
			}
			// Message
			if filter, ok := filterResponse.Message.(*parser.FilterResponse_Filter); ok {
				resp := &parser.FilterResponse{
					Message: filter,
				}
				if err := srv.Send(resp); err != nil {
					slog.Error("failed to send filter to caller", "err", err)
					return err
				}
				filtersToSave = append(filtersToSave, filter.Filter)
			} else {
				slog.Error("received a non-Filter message")
				return fmt.Errorf("received a non-Filter message")
			}
		}
	} else {
		// Send filters from DB
		slog.Debug("Sending filters from database", "request", req)
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

func sendErrorStatus_GetCategoryFilters(srv parser.ItemParser_GetCategoryFiltersServer, errText string) error {
	resp := &parser.FilterResponse{
		Message: &parser.FilterResponse_Status{
			Status: &statuspb.Status{
				Code:    int32(codes.Internal),
				Message: errText,
			},
		},
	}
	return srv.Send(resp)
}
