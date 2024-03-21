package domainerr

// Case represents a specific error condition. For example: purchase_limit_exceeded, insufficient_inventory.
type Case interface {
	// Identifier returns a string that uniquely identifies this error case. It can be
	// a numerical value or a descriptive title/name. For example, two numerical values:
	// 1000, 1_1_1000; a descriptive title/name: purchase_limit_exceeded.
	Identifier() string

	// StatusCode returns the operation status Code to which this error case is mapped.
	StatusCode() Code
}
