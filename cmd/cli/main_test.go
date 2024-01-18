package main

import (
	"reflect"
	"testing"
)

func Test_mp4Info(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
		wantErr bool
	}{
		{
			name: "clouds.mp4",
			args: args{
				fn: "./clouds.mp4",
			},
			wantRes: []string{
				"video: 720 486",
			},
			wantErr: false,
		},
		{
			name: "lemon.mp4",
			args: args{
				fn: "./lemon.mp4",
			},
			wantRes: []string{
				"video: 1280 720",
				"audio: 48000 hz",
			},
			wantErr: false,
		},
		{
			name: "dji.mp4",
			args: args{
				fn: "./dji.mp4",
			},
			wantRes: []string{
				"video: 1920 1080",
				"audio: 44100 hz",
			},
			wantErr: false,
		},
		{
			name: "space.mp4",
			args: args{
				fn: "./space.mp4",
			},
			wantRes: []string{
				"video: 1280 720",
				"audio: 48000 hz",
			},
			wantErr: false,
		},
		{
			name: "no file",
			args: args{
				fn: "./xxx.mp4",
			},
			wantRes: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mp4Info(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("mp4Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("mp4Info() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
