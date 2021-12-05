package handler

import "testing"

func TestHandler_TranslateFromEnToZh(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{text: "water"},
			want: "水",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{}
			if got := h.TranslateFromEnToZh(tt.args.text); got != tt.want {
				t.Errorf("TranslateFromEnToZh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_TranslateFromZhToEn(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{text: "水"},
			want: "Water",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{}
			if got := h.TranslateFromZhToEn(tt.args.text); got != tt.want {
				t.Errorf("TranslateFromZhToEn() = %v, want %v", got, tt.want)
			}
		})
	}
}
