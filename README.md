# Transaction ID library

This library supports transaction id handling. It provides methods for checking an http request for an
'X-Request-Id' header, and if present using the value of this as a transaction id.

If the header isn't present, the library will generate a transaction with a 'tid_' prefix
 and a random 10 character string.

Go provides support for passing around variables associated with a request lifecycle via the [context package](http://www.gorillatoolkit.org/pkg/context)

Best practice for using this is to specify a context as the first argument to your function
 (see [here](https://blog.golang.org/context) for examples that this library draws on).

For outputting a transactionID to logrus, use like this:

	log := log.WithFields(log.Fields{
		transactionIdKey: ctx.Value(transactionIdKey),
	})
