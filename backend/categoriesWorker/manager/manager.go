package manager

// func UpdateMarketRootCategories() error {
// 	d, err := db.NewDBConnection()
// 	if err != nil {
// 		return fmt.Errorf("failed to create a new database connection: %v", err)
// 	}

// 	markets, err := d.DBGetMarketsWithEmptyUpdateTime()
// 	if err != nil {
// 		return fmt.Errorf("failed to get markets with empty update time: %v", err)
// 	}

// 	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
// 	conn, err := grpc.Dial(parserAddr, grpc.WithInsecure())
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to parser service: %v", err)
// 	}
// 	defer conn.Close()

// 	client := parser.NewCategoryParserClient(conn)

// 	for _, marketName := range markets {
// 		// Create a RootCategoriesRequest for the gRPC call
// 		req := &parser.RootCategoriesRequest{
// 			Market: parser.Markets(parser.Markets_value[marketName]),
// 		}

// 		// Call the GetRootCategories method from the gRPC client
// 		resp, err := client.GetRootCategories(context.Background(), req)
// 		if err != nil {
// 			return fmt.Errorf("failed to get root categories from parser service for market %s: %v", marketName, err)
// 		}

// 		// Handle the response
// 		if status := resp.GetStatus(); status != nil {
// 			// Handle the error status
// 			continue // Log the error and continue with the next market
// 		}

// 		// Extract the category data from the response
// 		categories := resp.GetCategory()

// 		// Save the categories to the database
// 		for _, category := range categories {
// 			err := d.DBSaveCategory(*category)
// 			if err != nil {
// 				return fmt.Errorf("failed to save category for market %s: %v", marketName, err)
// 			}
// 		}

// 		// Update the market's update time in the database
// 		err = d.DBUpdateMarketUpdateTime(marketName)
// 		if err != nil {
// 			return fmt.Errorf("failed to update market update time for market %s: %v", marketName, err)
// 		}
// 	}

// 	return nil
// }
