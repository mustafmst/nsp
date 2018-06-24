package nsp

import (
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
		// TODO: Add test cases.
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
