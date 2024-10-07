package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

// ?

func processNotification(n notification) (string, int) {
	switch noti := n.(type) {
	case directMessage:
		return noti.senderUsername, noti.importance()
	case groupMessage:
		return noti.groupName, noti.importance()
	case systemAlert:
		return noti.alertCode, noti.importance()
	default:
		return "", 0
	}
}