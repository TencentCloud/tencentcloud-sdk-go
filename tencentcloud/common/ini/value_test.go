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
		{"valid bool", fields{raw: "true"}, true, false},
		{"valid bool", fields{raw: "y"}, true, false},
		{"invalid bool", fields{raw: "TrUe"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bool() got = %v, want %v", got, tt.want)
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
		{"valid float32", fields{raw: "1.23"}, 1.23, false},
		{"valid float32", fields{raw: "0.33333"}, 0.33333, false},
		{"invalid float32", fields{raw: "1.23.23"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.float32()
			if (err != nil) != tt.wantErr {
				t.Errorf("float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("float32() got = %v, want %v", got, tt.want)
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
		{"valid float64", fields{raw: "1.23"}, 1.23, false},
		{"valid float64", fields{raw: "0.33333"}, 0.33333, false},
		{"invalid float64", fields{raw: "1.23.23"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.float64()
			if (err != nil) != tt.wantErr {
				t.Errorf("float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("float64() got = %v, want %v", got, tt.want)
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
		{"valid int", fields{raw: "1"}, 1, false},
		{"valid int", fields{raw: "99887766"}, 99887766, false},
		{"invalid int", fields{raw: "1987a"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.int()
			if (err != nil) != tt.wantErr {
				t.Errorf("int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("int() got = %v, want %v", got, tt.want)
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
		{"valid int", fields{raw: "1"}, 1, false},
		{"valid int", fields{raw: "99887766"}, 99887766, false},
		{"invalid int", fields{raw: "1987a"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				raw: tt.fields.raw,
			}
			got, err := v.int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
