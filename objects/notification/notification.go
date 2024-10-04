package notification

type SendNotification struct {
	Title        string
	Body         string
	UserIDs      []int
	TargetTokens []string
}
