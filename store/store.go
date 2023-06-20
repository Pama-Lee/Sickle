package store

import "Sickle/entity"

var GlobalConfig []entity.Webhook

var UUID2Config map[string]entity.Webhook = make(map[string]entity.Webhook)
