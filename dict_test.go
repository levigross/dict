package dict

import (
	"testing"
)

func TestFirstByte(t *testing.T) {
	if len(DictionaryBytes) != 3193759 {
		t.Fatal("Len of the bytes should be 3193759")
	}
	if len(DictionaryString) != 312095 {
	t.Fatal("Len of string should be 312095")
	}
}