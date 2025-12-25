package tasks

import "fmt"

func LabSeventeen() {
	var all_users = [...]string{"id3", "id5", "id9", "id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	// пользователи не в сети
	var offline_users = [...]string{"id3", "id9", "id7", "id2", "id4", "id6"}

	for _, user := range all_users {
		for _, offlineUser := range offline_users {
			if user == offlineUser {
				fmt.Printf("User %s is offline\n", user)
			}
		}
	}

}
