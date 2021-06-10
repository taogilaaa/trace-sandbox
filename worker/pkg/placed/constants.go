package placed

const (
	NatsChannel     = "sandbox.saleorder.placed"
	NatsQueueGroup  = "sandbox.saleorder.placed.worker"
	NatsDurableName = "create-saleorder"
	NatsMaxInflight = 1
)
