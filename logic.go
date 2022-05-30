package minimal

func NOT (b1 bit)(bit){
	if b1 == 0 {
		return 1
	} else {
		return 0
	}
}

func AND (b1, b2 bit)(bit){
	if b1 == 1 && b2 == 1 {
		return 1
	} else {
		return 0
	}
}

func OR (b1, b2 bit)(bit) {
	if b1 == 0 && b2 == 0 {
		return 0
	} else {
		return 1
	}
}