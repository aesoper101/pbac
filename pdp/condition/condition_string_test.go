package condition

import (
	"github.com/aesoper101/pbac/pdp/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newStringEqualsCondition(t *testing.T) {
	type args struct {
		key    string
		values []interface{}
	}
	tests := []struct {
		name  string
		args  args
		value interface{}
		want  bool
	}{
		{
			name: "test for string equals",
			args: args{
				key:    "test",
				values: []interface{}{"test", "test2"},
			},
			value: "test",
			want:  true,
		},
		{
			name: "test for string not equals",
			args: args{
				key:    "test",
				values: []interface{}{"test", "test2"},
			},
			value: "test3",
			want:  false,
		},
		{
			name: "test for string equals for empty values",
			args: args{
				key:    "test",
				values: []interface{}{},
			},
			value: "test3",
			want:  false,
		},
		{
			name: "test for string equals for empty value",
			args: args{
				key:    "test",
				values: []interface{}{"test", "test2"},
			},
			value: "",
			want:  false,
		},
		{
			name: "test for string equals for nil value",
			args: args{
				key:    "test",
				values: []interface{}{"test", "test2"},
			},

			want: false,
		},
		{
			name: "test for string equals for slice value",
			args: args{
				key:    "test",
				values: []interface{}{"test", "test2"},
			},
			value: []string{"test"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newStringEqualsCondition(tt.args.key, tt.args.values)

			ctrl := gomock.NewController(t)
			ctx := mock.NewMockEvalContextor(ctrl)

			ctx.EXPECT().GetAttribute(tt.args.key).Return(tt.value, true).AnyTimes()

			assert.Equal(t, tt.want, got.Evaluate(tt.value, ctx))
		})
	}
}
