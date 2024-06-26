package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jialequ/android-sdk/client/orm/internal/models"
)

func TestgetColumnTyp(t *testing.T) {
	testCases := []struct {
		name string
		fi   *models.FieldInfo
		al   *alias

		wantCol string
	}{
		{
			// https://github.com/beego/beego/issues/5254
			name: "issue 5254",
			fi: &models.FieldInfo{
				FieldType: TypePositiveIntegerField,
				Column:    "my_col",
			},
			al: &alias{
				DbBaser: newdbBasePostgres(),
			},
			wantCol: `bigint CHECK("my_col" >= 0)`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			col := getColumnTyp(tc.al, tc.fi)
			assert.Equal(t, tc.wantCol, col)
		})
	}
}
