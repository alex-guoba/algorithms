package union

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestUnionFind_Count(t *testing.T) {
	type fields struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			"count",
			fields{
				4,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := New(tt.fields.count)

			if got := uf.Count(); got != tt.want {
				t.Errorf("UnionFind.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Connected(t *testing.T) {
	type fields struct {
		count int
	}
	type arg struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		ags    []arg
	}{
		// TODO: Add test cases.
		{
			"connected",
			fields{
				4,
			},
			[]arg{arg{0, 1}, arg{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := New(tt.fields.count)

			for _, a := range tt.ags {
				if got := uf.Connected(a.p, a.q); got != false {
					t.Errorf("UnionFind.Connected() error")
				}
				uf.Union(a.p, a.q)
				if got := uf.Connected(a.p, a.q); got != true {
					t.Errorf("UnionFind.Connected() failed")
				}
			}
		})
	}
}

func TestUnionFind_Find(t *testing.T) {
	type fields struct {
		count int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			"find",
			fields{
				4,
			},
			args{
				2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := New(tt.fields.count)
			if got := uf.Find(tt.args.v); got != tt.want {
				t.Errorf("UnionFind.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Depth(t *testing.T) {
	uf := New(6)

	// 0 -> 1[root]
	uf.Union(0, 1)
	assert.Equal(t, uf.Depth(0, 1), 2, "depth error")
	assert.Equal(t, uf.Depth(1, 1), 1, "depth error")

	// 3 -> 2[root]
	uf.Union(3, 2)
	assert.Equal(t, uf.Depth(3, 1), 2, "depth error")
	assert.Equal(t, uf.Depth(2, 1), 1, "depth error")

	// 3 -> 2[root]
	// 4 -> 2[root]
	uf.Union(3, 4)
	assert.Equal(t, uf.Depth(3, 1), 2, "depth error")
	assert.Equal(t, uf.Depth(4, 1), 2, "depth error")
	assert.Equal(t, uf.Depth(2, 1), 1, "depth error")

	// 3 -> 2[root]
	// 4 -> 2[root]
	// 5 -> 2[root]
	uf.Union(5, 4)
	assert.Equal(t, uf.Depth(5, 1), 2, "depth error")

	// 0 -> 1 -> 2[root]
	// 		3 -> 2[root]
	// 		4 -> 2[root]
	// 		5 -> 2[root]
	uf.Union(0, 4)
	assert.Equal(t, uf.Depth(1, 1), 2, "depth error")
	assert.Equal(t, uf.Depth(0, 1), 3, "depth error")
	assert.Equal(t, uf.Depth(4, 1), 2, "depth error")

	assert.Equal(t, uf.Connected(0, 4), true, "connected error")
}

func TestUnionFind_Union(t *testing.T) {
	type fields struct {
		count int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"union",
			fields{
				4,
			},
			args{
				0,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := New(tt.fields.count)

			if uf.Connected(tt.args.p, tt.args.q) {
				t.Errorf("should not conneted before union")
			}
			c := uf.Count()

			uf.Union(tt.args.p, tt.args.q)

			if !uf.Connected(tt.args.p, tt.args.q) {
				t.Errorf("should conneted fater union")
			}

			assert.Equal(t, c-1, uf.Count(), "count incorretct")
		})
	}
}
