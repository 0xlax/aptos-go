package aptostypes

var codeAddress = [16]uint8{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
}

// Currency converts given currency code string into Move TypeTag that is required by
// move script argument.
func Currency(code string) TypeTag {
	return &TypeTag__Struct{
		Value: StructTag{
			Address:    codeAddress,
			Module:     Identifier(code),
			Name:       Identifier(code),
			TypeParams: []TypeTag{},
		},
	}
}
