package ini

import "testing"

func Test_value_String(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"valid", fields{"valid string"}, "valid string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_bool(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{"valid Bool", fields{raw: "true"}, true, false},
		{"valid Bool", fields{raw: "y"}, true, false},
		{"invalid Bool", fields{raw: "TrUe"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_float32(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float32
		wantErr bool
	}{
		{"valid Float32", fields{raw: "1.23"}, 1.23, false},
		{"valid Float32", fields{raw: "0.33333"}, 0.33333, false},
		{"invalid Float32", fields{raw: "1.23.23"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.Float32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_float64(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{"valid Float64", fields{raw: "1.23"}, 1.23, false},
		{"valid Float64", fields{raw: "0.33333"}, 0.33333, false},
		{"invalid Float64", fields{raw: "1.23.23"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.Float64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_int(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"valid Int", fields{raw: "1"}, 1, false},
		{"valid Int", fields{raw: "99887766"}, 99887766, false},
		{"invalid Int", fields{raw: "1987a"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_int64(t *testing.T) {
	type fields struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"valid Int", fields{raw: "1"}, 1, false},
		{"valid Int", fields{raw: "99887766"}, 99887766, false},
		{"invalid Int", fields{raw: "1987a"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.Int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
