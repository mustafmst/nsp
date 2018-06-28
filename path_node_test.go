package nsp

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_pathNode_AddMethod(t *testing.T) {
	type fields struct {
		childNodes map[string]PathNode
		methods    map[string]http.Handler
	}
	type args struct {
		method  string
		handler func(http.ResponseWriter, *http.Request)
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
			p := &pathNode{
				childNodes: tt.fields.childNodes,
				methods:    tt.fields.methods,
			}
			if got := p.AddMethod(tt.args.method, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathNode.AddMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathNode_ServeHTTP(t *testing.T) {
	type fields struct {
		childNodes map[string]PathNode
		methods    map[string]http.Handler
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
			p := &pathNode{
				childNodes: tt.fields.childNodes,
				methods:    tt.fields.methods,
			}
			p.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func Test_pathMethodHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		f func(http.ResponseWriter, *http.Request)
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
			pmh := &pathMethodHandler{
				f: tt.fields.f,
			}
			pmh.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func Test_pathNode_AddPath(t *testing.T) {
	type fields struct {
		childNodes map[string]PathNode
		methods    map[string]http.Handler
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
			p := &pathNode{
				childNodes: tt.fields.childNodes,
				methods:    tt.fields.methods,
			}
			if got := p.AddPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathNode.AddPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPathNode(t *testing.T) {
	tests := []struct {
		name string
		want PathNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPathNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPathNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
