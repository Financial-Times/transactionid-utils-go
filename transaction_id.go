package transactionidutils

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/dchest/uniuri"
)

type tidKeyType int

const (
	//TransactionIDHeader is the request header to look for
	TransactionIDHeader = "X-Request-Id"
	//transactionIDKey is the key used to store the value on the context
	transactionIDKey tidKeyType = iota
)

// GetTransactionIDFromRequest will look on the request
// for an 'X-Request-Id' header, and use that value as the returned transactionID.
// If none is found, one will be autogenerated, with a 'tid_' prefix and a random
// ten character string and it will be set as request header.
func GetTransactionIDFromRequest(req *http.Request) string {
	transactionID := req.Header.Get(TransactionIDHeader)
	if transactionID == "" {
		transactionID = NewTransactionID()
		req.Header.Set(TransactionIDHeader, transactionID)
	}
	return transactionID
}

// NewTransactionID generates a new random transaction ID conforming to the FT spec
func NewTransactionID() string {
	return fmt.Sprintf("tid_%s", uniuri.NewLen(10))
}

// TransactionAwareContext  will take the
// context passed in and store the transactionID on it
func TransactionAwareContext(ctx context.Context, transactionID string) context.Context {
	return context.WithValue(ctx, transactionIDKey, transactionID)
}

// GetTransactionIDFromContext  will look for a transactionID
// value on the context and return it if found. If none is found, return empty
// string and an error
func GetTransactionIDFromContext(ctx context.Context) (string, error) {
	// ctx.Value returns nil if ctx has no value for the key;
	// string type assertion returns ok=false for nil.
	transactionID, ok := ctx.Value(transactionIDKey).(string)
	if ok {
		return transactionID, nil
	}
	return "", fmt.Errorf("no transactionID found")
}
