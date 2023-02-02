package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  [][]string
	}{
		"1": {
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		"2": {
			input: []string{""},
			want: [][]string{
				{""},
			},
		},
		"3": {
			input: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := groupAnagrams(tc.input)

			if diff := cmp.Diff(got, tc.want); diff != "" {
				fmt.Printf("got: %v\n want: %v\n", got, tc.want)
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		// 文字列ごとにアルファベット順（正確にはASCII順）に並べる
		// ex. bat -> abt, ate -> aet
		// ソートしたものをキーとしてそれぞれのグループを作成
		sorted := sortString(s)
		m[string(sorted)] = append(m[string(sorted)], s)
	}

	var res [][]string
	for _, v := range m {
		// {"tea", "ate", "eat"} -> {"ate", "eat", "tea"}のようにソート
		sort.Strings(v)
		res = append(res, v)
	}

	// グループ数の少ない順にソート
	sort.Slice(res, func(i, j int) bool { return len(res[i]) < len(res[j]) })
	return res
}

func sortString(str string) []byte {
	var tmp []byte
	for i := 0; i < len(str); i++ {
		tmp = append(tmp, str[i])
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
	return tmp
}
