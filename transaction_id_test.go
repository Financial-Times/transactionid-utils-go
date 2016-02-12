package transactionidutils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const knownTransactionID = "KnownTransactionId"

func TestGetKnownTransactionIDFromRequest(t *testing.T) {
	assert := assert.New(t)

	header := http.Header{}
	header.Set(TransactionIDHeader, knownTransactionID)

	req := &http.Request{Header: header}

	transactionID := GetTransactionIDFromRequest(req)

	assert.Equal(knownTransactionID, transactionID, "Didn't get the known transactionID returned")
}

func TestGetGeneratedTransactionIDWhenNoneOnRequest(t *testing.T) {
	assert := assert.New(t)

	header := http.Header{}
	req := &http.Request{Header: header}

	transactionID := GetTransactionIDFromRequest(req)

	assert.NotNil(transactionID)
	assert.Contains(transactionID, "tid", "Didn't get the expected generated prefix")
}

func TestGetDifferentGeneratedTransactionIDs(t *testing.T) {
	assert := assert.New(t)

	header := http.Header{}
	req := &http.Request{Header: header}

	transactionID := GetTransactionIDFromRequest(req)
	req.Header.Del(TransactionIDHeader)
	secondTransactionID := GetTransactionIDFromRequest(req)
	assert.NotEqual(transactionID, secondTransactionID, "Transaction IDs not unique")
}

func TestTransactionAwareContextAddsTransactionIDToContext(t *testing.T) {
	assert := assert.New(t)
	transactionAwareContext := TransactionAwareContext(context.Background(), knownTransactionID)
	assert.Equal(knownTransactionID, transactionAwareContext.Value(TransactionIDKey), "wrong transactionID on context")
}

func TestTransactionAwareContextOverridesAnyExistingTransactionIDToContext(t *testing.T) {
	assert := assert.New(t)
	existingTransactionAwareContext := TransactionAwareContext(context.Background(), "different transaction ID")
	transactionAwareContext := TransactionAwareContext(existingTransactionAwareContext, knownTransactionID)
	assert.Equal(knownTransactionID, transactionAwareContext.Value(TransactionIDKey), "wrong transactionID on context")
}

func TestCanGetTransactionIDFromContext(t *testing.T) {
	assert := assert.New(t)
	transactionAwareContext := TransactionAwareContext(context.Background(), knownTransactionID)
	transactionID, err := GetTransactionIDFromContext(transactionAwareContext)
	assert.NoError(err, "Got unexpected error")
	assert.Equal(knownTransactionID, transactionID, "Didn't get the known transactionID returned")
}

func TestErrorHandlingForNoTransactionIDOnContext(t *testing.T) {
	assert := assert.New(t)
	transactionID, err := GetTransactionIDFromContext(context.Background())
	assert.Error(err, "No error returned")
	assert.Equal("", transactionID, "transactionID should be empty on error")
}
