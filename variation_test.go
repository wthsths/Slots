package slots

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewVariationFromConfig(t *testing.T) {
	tests := []struct {
		name              string
		variationVal      string
		variationFilePerm fs.FileMode
		want              *Variation
		wantErr           error
	}{
		{
			name:              "TestNewVariationFromConfig should return no error",
			variationVal:      `{ "slug": "test" }`,
			variationFilePerm: 0644,
			want: &Variation{
				Slug: "test",
			},
		},
		{
			name:              "TestNewVariationFromConfig should return invalid json error",
			variationVal:      `{ "slug": test" }`,
			variationFilePerm: 0644,
			wantErr:           errInvalidJson,
		},
		{
			name:              "TestNewVariationFromConfig should return no file or perm error",
			variationVal:      ``,
			variationFilePerm: 0000,
			wantErr:           errReadingConfigFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempConfigName := fmt.Sprintf("./TestNewVariationFromConfig_%d.json", time.Now().UnixMicro())
			ioutil.WriteFile(tempConfigName, []byte(tt.variationVal), tt.variationFilePerm)
			defer os.RemoveAll(tempConfigName)

			got, gotErr := NewVariationFromConfig(tempConfigName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestNewVariationFromString(t *testing.T) {
	tests := []struct {
		name         string
		variationVal string
		want         *Variation
		wantErr      error
	}{
		{
			name:         "TestNewVariationFromConfig should return no error",
			variationVal: `{ "slug": "test" }`,
			want: &Variation{
				Slug: "test",
			},
		},
		{
			name:         "TestNewVariationFromConfig should return invalid json error",
			variationVal: `{ "slug": test" }`,
			wantErr:      errInvalidJson,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := NewVariationFromString(tt.variationVal)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
