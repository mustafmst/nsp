package nsp

import (
	"net/http"
	"reflect"
	"testing"
)

type writer struct {
	msg []byte
}

func (w *writer) Header() http.Header {
	return *new(http.Header)
}
func (w *writer) Write(b []byte) (int, error) {
	w.msg = b
	return 1, nil
}
func (w *writer) WriteHeader(statusCode int) {}

func testMethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

var r *http.Request = new(http.Request)

func Test_controller_GetMethod(t *testing.T) {
	type fields struct {
		name    string
		methods map[string]func(w http.ResponseWriter, r *http.Request)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{"GetMethod return proper method",
			fields{"testController", map[string]func(w http.ResponseWriter, r *http.Request){"testMethod": testMethod}},
			args{"testMethod"},
			[]byte("test")},
		{"GetMethod return proper method",
			fields{"testController", map[string]func(w http.ResponseWriter, r *http.Request){"otherTestMethod": testMethod}},
			args{"testMethod"},
			[]byte("There is no testMethod method in testController controller!")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				name:    tt.fields.name,
				methods: tt.fields.methods,
			}
			m := c.GetMethod(tt.args.name)
			w := writer{[]byte("")}
			m(&w, r)
			if string(w.msg) != string(tt.want) {
				t.Errorf("controller.GetMethod() returned wrong method.")
			}
		})
	}
}

func Test_controller_AddMethod(t *testing.T) {
	type fields struct {
		name    string
		methods map[string]func(w http.ResponseWriter, r *http.Request)
	}
	type args struct {
		name   string
		method func(w http.ResponseWriter, r *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Method is added to controller",
			fields{"testController", make(map[string]func(w http.ResponseWriter, r *http.Request))},
			args{"testMethod", testMethod}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				name:    tt.fields.name,
				methods: tt.fields.methods,
			}
			c.AddMethod(tt.args.name, tt.args.method)
			if c.methods[tt.args.name] == nil {
				t.Errorf("controller.AddMethod() didn't added method to controller.")
			}
		})
	}
}

func Test_controller_GetName(t *testing.T) {
	type fields struct {
		name    string
		methods map[string]func(w http.ResponseWriter, r *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"GetName returns controller real name",
			fields{"testName", make(map[string]func(w http.ResponseWriter, r *http.Request))}, "testName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				name:    tt.fields.name,
				methods: tt.fields.methods,
			}
			if got := c.GetName(); got != tt.want {
				t.Errorf("controller.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewController(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Controller
	}{
		{"home", args{"home"}, &controller{"home", make(map[string]func(w http.ResponseWriter, r *http.Request))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
