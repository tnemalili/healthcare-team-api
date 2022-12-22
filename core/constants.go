package core

import "os"

var (
	PhoneNumberValidationRegexp = `^\+?\d{1,4}?[-.\s]?[-.\s]?\d{1,4}[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`
	DARE                        = "dare"
	CREDENTIALS                 = "credentials"
	StorageUrl                  = os.Getenv("STORAGE_URL")
	StorageApiKey               = os.Getenv("STORAGE_API_KEY")
)
