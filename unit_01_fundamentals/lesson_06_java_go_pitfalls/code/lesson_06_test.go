package lesson_06_java_go_pitfalls

import (
	"context"
	"testing"
	"time"
)

func TestSetAge(t *testing.T) {
	u := UserValue{Name: "Bob", Age: 20}
	u.SetAgeWrong(30)
	if u.Age == 30 {
		t.Error("SetAgeWrong should not modify age")
	}

	u.SetAgeCorrect(30)
	if u.Age != 30 {
		t.Error("SetAgeCorrect should modify age")
	}
}

func TestGetUser(t *testing.T) {
	errWrong := GetUserWrong(1)
	if errWrong == nil {
		t.Error("GetUserWrong should return non-nil error interface")
	}

	errCorrect := GetUserCorrect(1)
	if errCorrect != nil {
		t.Error("GetUserCorrect should return nil error")
	}
}

func TestSafeIncrement(t *testing.T) {
	count = 0
	go safeIncrement()
	go safeIncrement()
	time.Sleep(200 * time.Millisecond)
	if count != 20000 {
		t.Errorf("Expected count to be 20000, got %d", count)
	}
}

func TestNoLeakGoroutine(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	noLeakGoroutine(ctx)
}
