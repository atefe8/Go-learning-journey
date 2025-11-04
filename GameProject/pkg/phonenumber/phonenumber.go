package phonenumber

func IsVaild(phoneNumber string) bool {
	if len(phoneNumber) != 11 {
		return false
	}

	if phoneNumber[0:3] != "09" {
		return false
	}

	return true
}
