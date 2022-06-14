package entity

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
	"time"
)

func TestMessage_ToJSON(t *testing.T) {
	type fields struct {
		ID        string
		UserID    string
		Title     string
		Content   string
		Long      string
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				ID:        "f32164df-d048-4574-b4a4-49c652bf911f",
				UserID:    "test",
				Title:     "test",
				Content:   "test",
				Long:      "test",
				CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			want: "{\"content\":\"test\",\"created_at\":\"2009-11-10T23:00:00Z\",\"id\":\"f32164df-d048-4574-b4a4-49c652bf911f\",\"long\":\"test\",\"title\":\"test\",\"user_id\":\"test\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
				Long:      tt.fields.Long,
				CreatedAt: tt.fields.CreatedAt,
			}
			got, err := m.ToJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_ToGinH(t *testing.T) {
	type fields struct {
		ID        string
		UserID    string
		Title     string
		Content   string
		Long      string
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.H
	}{
		{
			name: "test",
			fields: fields{
				ID:        "f32164df-d048-4574-b4a4-49c652bf911f",
				UserID:    "test",
				Title:     "test",
				Content:   "test",
				Long:      "test",
				CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			want: gin.H{
				"id":         "f32164df-d048-4574-b4a4-49c652bf911f",
				"user_id":    "test",
				"title":      "test",
				"content":    "test",
				"long":       "test",
				"created_at": "2009-11-10T23:00:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
				Long:      tt.fields.Long,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := m.ToGinH(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToGinH() = %v, want %v", got, tt.want)
			}
		})
	}
}
