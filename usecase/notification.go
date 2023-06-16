package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetNotifications(id uint) (resp []payload.GetNotification, err error) {
	user, err := database.GetUser(id)
	if err != nil {
		return
	}
	for _, value := range user.Notifications {
		resp = append(resp, payload.GetNotification{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			Text:      value.Text,
			IsRead:    value.IsRead,
		})
	}
	return
}

func GetNotificationById(id uint) (resp payload.GetNotification, err error) {
	notif, err := database.GetNotificationById(id)
	if err != nil {
		return
	}

	resp = payload.GetNotification{
		ID:        notif.ID,
		CreatedAt: notif.CreatedAt,
		Text:      notif.Text,
		IsRead:    notif.IsRead,
	}
	return
}

func UpdateNotificationStatus(id uint) error {
	notif, err := database.GetNotificationById(id)
	if err != nil {
		return err
	}
	if notif.IsRead {
		notif.IsRead = false
	} else {
		notif.IsRead = true
	}
	if err := database.SaveNotification(notif); err != nil {
		return err
	}
	return nil
}

func DeleteNotification(id uint) error {
	if err := database.DeleteNotificationById(id); err != nil {
		return err
	}
	return nil
}
