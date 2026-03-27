package structs_interfaces

import "testing"

func TestBirthdayUpdatesAge(t *testing.T) {
	user := User{Name: "Alice", Age: 30}
	user.Birthday()

	if user.Age != 31 {
		t.Fatalf("user.Age = %d, want 31", user.Age)
	}
}

func TestInterfaceImplementation(t *testing.T) {
	var speaker Speaker

	user := User{Name: "Alice", Age: 30}
	admin := Admin{
		User: User{Name: "Bob", Age: 40},
		Role: "SuperAdmin",
	}

	speaker = user
	if speaker == nil {
		t.Fatalf("speaker should not be nil after assigning user")
	}

	speaker = admin
	if speaker == nil {
		t.Fatalf("speaker should not be nil after assigning admin")
	}
}
