package ini

import (
	"reflect"
	"testing"
)

func Test_openFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			"invalid path", args{path: "./testdata"}, []byte{}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := openFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("openFile() gotData = \n%v\n, want = \n%v\n", gotData, tt.wantData)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *sections
		wantErr bool
	}{
		{"valid ini", args{path: "./testdata/valid.ini"}, nil, false},
		{"invalid ini", args{"./testdata/invalid.ini"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Parse() got = %v, want %v", got, tt.want)
			//}
			if err == nil {
				t.Log(got.Section("default"))
				if got.Section("default").Key("key1").String() == "" {
					t.Errorf("Parse() error:not get default.key1 value")
				}
				if _, e := got.Section("custom").Key("customKey2").Bool(); e != nil {
					t.Errorf("Parse() error:not get custom.customKey2 value")
				}
			}
		})
	}
}
