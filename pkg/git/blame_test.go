package git

import (
	"testing"
)

func TestParseOwnerRepo(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOwner  string
		wantRepo   string
		wantErrNil bool
	}{
		{
			name:       "HTTPS URL",
			input:      "owner/repo.git",
			wantOwner:  "owner",
			wantRepo:   "repo",
			wantErrNil: true,
		},
		{
			name:       "SSH URL",
			input:      "owner/repo",
			wantOwner:  "owner",
			wantRepo:   "repo",
			wantErrNil: true,
		},
		{
			name:       "Invalid URL",
			input:      "invalid",
			wantOwner:  "",
			wantRepo:   "",
			wantErrNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOwner, gotRepo, err := parseOwnerRepo(tt.input)
			if (err == nil) != tt.wantErrNil {
				t.Errorf("parseOwnerRepo() error = %v, wantErrNil %v", err, tt.wantErrNil)
				return
			}
			if gotOwner != tt.wantOwner {
				t.Errorf("parseOwnerRepo() gotOwner = %v, want %v", gotOwner, tt.wantOwner)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("parseOwnerRepo() gotRepo = %v, want %v", gotRepo, tt.wantRepo)
			}
		})
	}
}