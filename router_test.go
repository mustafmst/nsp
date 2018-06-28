package nsp

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicRouter_ServeHTTP(t *testing.T) {
	type fields struct {
		paths PathNode
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &basicRouter{
				paths: tt.fields.paths,
			}
			rt.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func Test_basicRouter_AddPath(t *testing.T) {
	type fields struct {
		paths PathNode
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PathNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &basicRouter{
				paths: tt.fields.paths,
			}
			if got := rt.AddPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("basicRouter.AddPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
