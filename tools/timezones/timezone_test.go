package utils

import (
	"testing"
)

func TestGetTimeZoneByLatLon(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case #1",
			args: args{
				lat: 11.11,
				lon: 11.11,
			},
			want: "Africa/Lagos",
		},
		{
			name: "Test case #2",
			args: args{
				lat: 66.66,
				lon: 66.66,
			},
			want: "Asia/Yekaterinburg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTimeZoneByLatLon(tt.args.lat, tt.args.lon)
			if result != tt.want {
				t.Errorf("Expected: %s, actual: %s.", tt.want, result)
			}
		})
	}
}
