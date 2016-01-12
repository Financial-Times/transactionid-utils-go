# Transaction ID library

If you want to log transactionID, you need to pass the context as the first field in your function (see TODO add link to website article).

For logrus, use like this:

	log := log.WithFields(log.Fields{
		transactionIdKey: ctx.Value(transactionIdKey),
	})

	
