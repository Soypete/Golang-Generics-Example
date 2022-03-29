package main

import (
	"fmt"
	"testing"
)

// func TestDataMap(t *testing.T) {

// 	tt := []struct {
// 		name string
// 		data map[string]any
// 		want want
// 	}{
// 		{
// 			name: "test1",
// 			data: map[string]any{
// 				"name": "John",
// 				"age":  30,
// 				"phone": map[string]any{
// 					"home":  "555-1234",
// 					"work":  "555-6789",
// 					"email": "example@test.com",
// 				},
// 				"lastAppointmet": "2020-01-01",
// 			},
// 			want: want{
// 				name: "John",
// 				age:  30,
// 				contact: struct {
// 					phone  string
// 					mobile string
// 					email  string
// 				}{
// 					phone:  "555-1234",
// 					mobile: "555-6789",
// 					email:  "example@test.com",
// 				},
// 			},
// 		},
// 		{
// 			name: "test2",
// 			data: map[string]any{
// 				"name": "John",
// 				"age":  30,
// 				"phone": map[string]any{
// 					"home":  4055551234,
// 					"work":  4055556789,
// 					"email": "example@test.com",
// 				},
// 				"lastAppointmet": 1577862000000,
// 			},
// 			want: want{
// 				name: "John",
// 				age:  30,
// 				contact: struct {
// 					phone  string
// 					mobile string
// 					email  string
// 				}{
// 					phone:  "555-1234",
// 					mobile: "555-6789",
// 					email:  "example@test.com",
// 				},
// 			},
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got := DataMap(tc.data)
// 			if got != tc.want {
// 				t.Errorf("got %v, want %v", got, tc.want)
// 			}
// 		})
// 	}
// }

func Test_makePhone(t *testing.T) {

	tt := []struct {
		name string
		data any
		// data interface{string|int64} // interface{}/any
		want    string
		wantErr error
	}{
		{
			name: "test1",
			data: "(405)-555-2345",
			want: "405-555-2345",
		},
		{
			name:    "test2",
			data:    "555-1234",
			want:    "",
			wantErr: fmt.Errorf("invalid phone number"),
		},
		{
			name:  "test3",
			data:  "405-555-1234",
			want1: "405-555-1234",
		},
		{
			name: "test4",
			data: 4055551234,
			want: "405-555-1234",
		},
		{
			name:    "test5",
			data:    555234,
			want:    "",
			wantErr: fmt.Errorf("invalid phone number"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := makePhone(tc.data)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
			if err != tc.wantErr {
				t.Errorf("got %v, want %v", err, tc.wantErr)
			}
		})
	}
}
