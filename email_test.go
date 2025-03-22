package main

import (
	"testing"
)

func TestEmail(t *testing.T) {
	es := NewEmailSender()
	err := es.Send("xxx@qq.com", "1", "test")
	if err != nil {
		t.Error(err)
	}
	t.Log("successfully sent email")
}
