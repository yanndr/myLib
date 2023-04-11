package api

import "testing"

func Test_model_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		m       Serializable
		want    string
		wantErr bool
	}{
		{name: "empty", m: Author{}, want: "f21575bb3d7b27b22f6c9b345f4250dc", wantErr: false},
		{name: "not empty", m: Author{
			ID: 1,
			AuthorBase: AuthorBase{
				LastName:   "Salut",
				FirstName:  "Les",
				MiddleName: "Gars",
			},
		}, want: "3115a613fe9c96622488db1700c33b26", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.m.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
