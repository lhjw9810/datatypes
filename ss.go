package datatypes

//
//func (date Date) MarshalJSON() ([]byte, error) {
//	t := time.Time(date)
//	if y := t.Year(); y < 0 || y >= 10000 {
//		// RFC 3339 is clear that years are 4 digits exactly.
//		// See golang.org/issue/4556#c15 for more discussion.
//		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
//	}
//
//	b := make([]byte, 0, len("2006-01-02")+2)
//	b = append(b, '"')
//	b = t.AppendFormat(b, "2006-01-02")
//	b = append(b, '"')
//	return b, nil
//}
//
//func (date *Date) UnmarshalJSON(b []byte) error {
//
//	data := b
//	// Ignore null, like in the main JSON package.
//	if string(data) == "null" {
//		return nil
//	}
//	// Fractional seconds are handled implicitly by Parse.
//	var err error
//	t, err := time.Parse(`"`+"2006-01-02"+`"`, string(data))
//	*date = Date(t)
//	return err
//}
