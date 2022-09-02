package datatypes

import (
	"encoding/json"
	"testing"
	"time"

	. "gorm.io/gorm/utils/tests"
)

func TestJSONEncoding(t *testing.T) {
	date := DateTime(time.Now())
	b, err := json.Marshal(date)
	if err != nil {
		t.Fatalf("failed to encode datatypes.Date: %v", err)
	}

	var got DateTime
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("failed to decode to datatypes.Date: %v", err)
	}

	AssertEqual(t, date, got)
}
