package remote

import "testing"

func TestProtoSerializerUnknownMessageType(t *testing.T) {
	message, err := newProtoSerializer().Deserialize("portability.UnknownMessage", nil)
	if err == nil || message != nil {
		t.Fatalf("expected an error and no message, got %v, %v", message, err)
	}
}
