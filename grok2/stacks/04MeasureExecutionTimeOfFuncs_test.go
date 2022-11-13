package stacks

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	tests := []struct {
		testID int
		input  []string
		want   []int
	}{
		{1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}, []int{8}},
		{2, []string{"0:start:0", "1:start:3", "1:end:6", "0:end:10"}, []int{7, 4}},
		{3, []string{"0:start:0", "0:end:0", "1:start:1", "1:end:1", "2:start:2", "2:end:2", "2:start:3", "2:end:3"}, []int{1, 1, 2}},
		{5, []string{"0:start:0", "1:start:5", "1:end:9", "4:start:10", "2:start:13", "2:end:15", "3:start:16", "3:end:18", "4:end:21", "0:end:22"}, []int{6, 5, 3, 3, 6}},
		{6, []string{"0:start:0", "1:start:5", "2:start:8", "3:start:12", "4:start:15", "5:start:19", "5:end:22", "4:end:24", "3:end:27", "2:end:32", "1:end:35", "0:end:36"}, []int{6, 6, 9, 6, 6, 4}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.testID), func(t *testing.T) {
			if got := exclusiveTime(tt.testID, tt.input); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
