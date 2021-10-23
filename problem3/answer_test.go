package main

import "testing"

func TestBoxCount(t *testing.T) {
	type args struct {
		cake  int
		apple int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				cake:  10,
				apple: 5,
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				cake:  20,
				apple: 25,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoxesCount(tt.args.cake, tt.args.apple); got != tt.want {
				t.Errorf("BoxesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInEachBox(t *testing.T) {
	type args struct {
		cake  int
		apple int
	}
	tests := []struct {
		name      string
		args      args
		wantCake  int
		wantApple int
	}{
		{
			name: "test1",
			args: args{
				cake:  10,
				apple: 5,
			},
			wantCake:  2,
			wantApple: 1,
		},
		{
			name: "test2",
			args: args{
				cake:  20,
				apple: 25,
			},
			wantCake:  4,
			wantApple: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCake, gotApple := InEachBox(tt.args.cake, tt.args.apple)
			if gotCake != tt.wantCake {
				t.Errorf("InEachBox() = %v, wantCake %v", gotCake, tt.wantCake)
			}
			if gotApple != tt.wantApple {
				t.Errorf("InEachBox() = %v, wantApple %v", gotApple, tt.wantApple)
			}
		})
	}
}
