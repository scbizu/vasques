package build

import (
	"testing"

	"github.com/scbizu/vasques/internal/pb"
	"github.com/stretchr/testify/require"
)

func TestUppercaseName(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{n: "test"}, "Test"},
		{"t2", args{n: "Test"}, "Test"},
		{"t2", args{n: "tEST"}, "TEST"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UppercaseName(tt.args.n); got != tt.want {
				t.Errorf("UppercaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenMessage(t *testing.T) {
	raw := map[string]interface{}{
		"foo": "bar",
	}
	got := flattenMessage("AutoGenerate", raw)
	f := map[string]pb.Field{"foo": pb.NewField(pb.KeywordTypeString, false, "foo")}
	m := map[string]map[string]pb.Field{
		"AutoGenerate": f,
	}
	require.Equal(t, m, got)
}

func TestFlatten_ONLY_STRING(t *testing.T) {
	raw := map[string]interface{}{
		"AutoGenerate": map[string]interface{}{
			"foo": "bar",
		},
	}
	got := flatten(raw)
	f := map[string]pb.Field{"foo": pb.NewField(pb.KeywordTypeString, false, "foo")}
	m := map[string]map[string]pb.Field{
		"AutoGenerate": f,
	}
	require.Equal(t, m, got)
}

func TestFlatten_ONLY_INT(t *testing.T) {
	raw := map[string]interface{}{
		"AutoGenerate": map[string]interface{}{
			"foo": float64(1),
		},
	}
	got := flatten(raw)
	f := map[string]pb.Field{"foo": pb.NewField(pb.KeywordTypeInt32, false, "foo")}
	m := map[string]map[string]pb.Field{
		"AutoGenerate": f,
	}
	require.Equal(t, m, got)
}
