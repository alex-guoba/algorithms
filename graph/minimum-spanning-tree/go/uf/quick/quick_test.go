package quick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind_Count(t *testing.T) {
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
			qf := New(tt.fields.count)

			if got := qf.Count(); got != tt.want {
				t.Errorf("QuickFind.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickFind_Connected(t *testing.T) {
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
		want   bool
	}{
		// TODO: Add test cases.
		{
			"connected",
			fields{
				4,
			},
			args{
				0,
				1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qf := New(tt.fields.count)

			if got := qf.Connected(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("QuickFind.Connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickFind_Find(t *testing.T) {
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
			qf := New(tt.fields.count)

			if got := qf.Find(tt.args.v); got != tt.want {
				t.Errorf("QuickFind.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickFind_Union(t *testing.T) {
	qf := New(4)

	c := qf.Count()

	assert.False(t, qf.Connected(0, 1), "unconnect")
	qf.Union(0, 1)
	assert.True(t, qf.Connected(0, 1), "connected")
	assert.Equal(t, c-1, qf.Count(), "count incorretct")

	assert.False(t, qf.Connected(2, 3), "unconnect")
	qf.Union(2, 3)
	assert.True(t, qf.Connected(2, 3), "connected")
	assert.Equal(t, c-2, qf.Count(), "count incorretct")

	qf.Union(1, 2)
	assert.True(t, qf.Connected(0, 3), "connected")
	assert.True(t, qf.Connected(0, 2), "connected")
	assert.True(t, qf.Connected(1, 3), "connected")
	assert.True(t, qf.Connected(1, 2), "connected")
	assert.Equal(t, 1, qf.Count(), "count incorretct")
}
