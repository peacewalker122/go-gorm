package middleware

//
//type CustomBinder struct{}
//
//func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
//	// First try to bind using the default Echo binder
//	if err := c.Bind(i); err != nil {
//		// If the default binder fails, try to bind using the custom binder
//		// Iterate over the struct fields
//		for i := 0; i < reflect.TypeOf(i).NumField(); i++ {
//			// Get the struct field
//			field := reflect.ValueOf(i).Elem().Field(i)
//
//			// Get the query tag value
//			queryTag := reflect.TypeOf(i).Field(i).Tag.Get("query")
//
//			// If the field has a query tag
//			if queryTag != "" {
//				// Get the query parameter value from the context
//				queryParam := c.QueryParam(queryTag)
//
//				// If the query parameter is empty, set the default value
//				if queryParam == "" {
//					defaultValue := reflect.TypeOf(i).Field(i).Tag.Get("default")
//					if defaultValue != "" {
//						queryParam = defaultValue
//					}
//				}
//
//				// Convert the query parameter value to the field's type
//				fieldValue := reflect.New(field.Type()).Elem()
//				if err := echo.QueryParamsBinder(map[string][]string{queryTag: {queryParam}}, fieldValue.Addr().Interface()); err != nil {
//					return echo.NewHTTPError(http.StatusBadRequest, "Invalid query parameter")
//				}
//
//				// Set the field value to the converted value
//				field.Set(fieldValue)
//			}
//		}
//	}
//
//	return nil
//}
