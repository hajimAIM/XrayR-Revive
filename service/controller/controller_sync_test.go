package controller

import (
	"reflect"
	"testing"

	"github.com/XrayR-project/XrayR/api"
)

func TestCompareUserList(t *testing.T) {
	tests := []struct {
		name        string
		oldList     []api.UserInfo
		newList     []api.UserInfo
		wantDeleted []api.UserInfo
		wantAdded   []api.UserInfo
	}{
		{
			name:        "No changes",
			oldList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			newList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			wantDeleted: nil,
			wantAdded:   nil,
		},
		{
			name:        "User added",
			oldList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			newList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}, {UID: 2, Email: "user2@example.com"}},
			wantDeleted: nil,
			wantAdded:   []api.UserInfo{{UID: 2, Email: "user2@example.com"}},
		},
		{
			name:        "User deleted",
			oldList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}, {UID: 2, Email: "user2@example.com"}},
			newList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			wantDeleted: []api.UserInfo{{UID: 2, Email: "user2@example.com"}},
			wantAdded:   nil,
		},
		{
			name:        "Users swapped",
			oldList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			newList:     []api.UserInfo{{UID: 2, Email: "user2@example.com"}},
			wantDeleted: []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			wantAdded:   []api.UserInfo{{UID: 2, Email: "user2@example.com"}},
		},
		{
			name:        "Empty to Non-empty",
			oldList:     []api.UserInfo{},
			newList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			wantDeleted: nil,
			wantAdded:   []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
		},
		{
			name:        "Non-empty to Empty",
			oldList:     []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			newList:     []api.UserInfo{},
			wantDeleted: []api.UserInfo{{UID: 1, Email: "user1@example.com"}},
			wantAdded:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDeleted, gotAdded := compareUserList(&tt.oldList, &tt.newList)

			// Helper to check slice equality (nil vs empty slice handling)
			checkSlice := func(got, want []api.UserInfo, name string) {
				if len(got) == 0 && len(want) == 0 {
					return
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s: got %v, want %v", name, got, want)
				}
			}

			checkSlice(gotDeleted, tt.wantDeleted, "Deleted")
			checkSlice(gotAdded, tt.wantAdded, "Added")
		})
	}
}
