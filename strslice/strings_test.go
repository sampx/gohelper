package strslice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	type args struct {
		vs []string
		f  func(string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "异常用例-非全数字",
			args: args{
				vs: []string{"1", "2", "3", "apple", "asdf"},
				f: func(str string) bool {
					if _, err := strconv.Atoi(str); err == nil {
						return true
					}
					return false
				},
			},
			want: false,
		},
		{
			name: "正常用例-全数字",
			args: args{
				vs: []string{"1", "2", "3"},
				f: func(str string) bool {
					if _, err := strconv.Atoi(str); err == nil {
						return true
					}
					return false
				},
			},
			want: true,
		},
		{
			name: "正常用例-全空字符串",
			args: args{
				vs: []string{},
				f: func(str string) bool {
					if len(str) == 0 {
						return true
					}
					return false
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.vs, tt.args.f); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		vs []string
		f  func(string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常用例-全空字符串",
			args: args{
				vs: []string{},
				f: func(str string) bool {
					if len(str) == 0 {
						return true
					}
					return false
				},
			},
			want: true,
		},
		{
			name: "正常用例-部分空字符串",
			args: args{
				vs: []string{"1", ""},
				f: func(str string) bool {
					if len(str) == 0 {
						return true
					}
					return false
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.vs, tt.args.f); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		vs []string
		f  func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "正常用例-筛选奇数",
			args: args{
				vs: []string{"1", "3", "5", "2", "4", "6"},
				f: func(str string) bool {
					if s, err := strconv.Atoi(str); err == nil {
						if s%2 != 0 {
							return true
						}
					}
					return false
				},
			},
			want: []string{"1", "3", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInclude(t *testing.T) {
	type args struct {
		vs []string
		t  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常用例-包含所选字符",
			args: args{
				vs: []string{"1", "3", "5", "2", "4", "6"},
				t:  "4",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Include(tt.args.vs, tt.args.t); got != tt.want {
				t.Errorf("Include() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		vs []string
		t  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常用例",
			args: args{
				vs: []string{"1", "3", "5", "2", "4", "6"},
				t:  "4",
			},
			want: 4,
		},
		{
			name: "异常用例",
			args: args{
				vs: []string{"1", "3", "5", "2", "4", "6"},
				t:  "8",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.vs, tt.args.t); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
