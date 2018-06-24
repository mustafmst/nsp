package nsp

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_controllersMap_AddController(t *testing.T) {
	type fields struct {
		controllers map[string]Controller
	}
	type args struct {
		controllerName string
		controller     Controller
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"AddController",
			fields{make(map[string]Controller)},
			args{"testController", NewController("testController")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &controllersMap{
				controllers: tt.fields.controllers,
			}
			cm.AddController(tt.args.controller)
			if c, ok := cm.controllers[tt.args.controllerName]; c == nil || !ok {
				t.Error("Controller wasn't added to map.")
			}
		})
	}
}

func Test_controllersMap_getController(t *testing.T) {
	type fields struct {
		controllers map[string]Controller
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Controller
		want1  bool
	}{
		{"Get existing controller",
			fields{map[string]Controller{"testController": NewController("testController")}},
			args{"testController"},
			NewController("testController"),
			true},
		{"Get controller that doesn't exists",
			fields{map[string]Controller{"otherTestController": NewController("testController")}},
			args{"testController"},
			nil,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &controllersMap{
				controllers: tt.fields.controllers,
			}
			got, got1 := cm.getController(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controllersMap.getController() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("controllersMap.getController() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_controllersMap_GetControllerMethod(t *testing.T) {
	createMap := func(c, m string, f func(http.ResponseWriter, *http.Request)) map[string]Controller {
		cm := make(map[string]Controller)
		ctrl := NewController(c)
		ctrl.AddMethod(m, f)
		cm[c] = ctrl
		return cm
	}
	var r = new(http.Request)
	// -----------------
	type fields struct {
		controllers map[string]Controller
	}
	type args struct {
		controller string
		method     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Get existing controller and method",
			fields{createMap("testController", "testMethod", getTestMethod("test"))},
			args{"testController", "testMethod"},
			"test"},
		{"Get non existing controller",
			fields{createMap("testController", "testMethod", getTestMethod("test"))},
			args{"otherTestController", "testMethod"},
			"There is no otherTestController controller!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &controllersMap{
				controllers: tt.fields.controllers,
			}
			m := cm.GetControllerMethod(tt.args.controller, tt.args.method)
			w := writer{""}
			m(&w, r)
			if w.msg != tt.want {
				t.Error("Wrong method returned")
			}
		})
	}
}

func TestNewControllersMap(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Create new ControllersMap", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewControllersMap(); got == nil && len(got.(*controllersMap).controllers) != tt.want {
				t.Errorf("NewControllersMap() controllers map is not empty")
			}
		})
	}
}
