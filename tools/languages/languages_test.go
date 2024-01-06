package languages

import "testing"

func TestGetLang(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want Language
	}{
		{
			name: "Test case #1",
			args: args{lang: "ru"},
			want: RUSSIAN,
		},
		{
			name: "Test case #2",
			args: args{lang: "en"},
			want: ENGLISH,
		},
		{
			name: "Test case #3",
			args: args{lang: "zz"},
			want: ENGLISH,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLang(tt.args.lang); got != tt.want {
				t.Errorf("GetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
