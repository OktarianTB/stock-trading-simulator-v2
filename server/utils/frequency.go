package util

const (
	DAILY = "daily"
	WEEKLY = "weekly"
	MONTHLY = "monthly"
	ANNUALLY = "annually"
)

func IsValidFrequency(frequency string) bool {
	switch frequency {
		case DAILY, WEEKLY, MONTHLY, ANNUALLY:
			return true
	}
	return false
}
