package main

import (
	"slices"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		desc     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			desc:     "default case",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			desc:     "empty nums2",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			desc:     "empty nums1",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !slices.Equal(tt.nums1, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, tt.nums1)
			}
		})
	}
}
