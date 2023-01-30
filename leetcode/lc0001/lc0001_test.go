package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		a, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
	fmt.Printf("%v", `[{"id":"dce83165-a672-413a-a444-a8920a85a38c","type":"title","items":[{"id":"59644888-4c06-4d20-9e37-93dafbd7a902","type":"sequence_title","start_time":880,"duration":3880,"width":300,"height":300,"position":{"x":171.53351771392028,"y":215.84176445382474},"scale":1,"asset_id":"622e853d27c711000155f53c@Public@CME","contents":{"text":"你好","text_style":{"font":"SimHei","font_size":60,"font_color":"#000000","font_alpha":100,"font_bold":0,"font_italic":0,"shadow_color":"","font_uline":0,"font_align":"center","bottom_color":"#ffffff","bottom_alpha":0,"background_color":"#ffffff","background_alpha":0,"border_width":0,"border_color":"#ffffff","font_asset_id":""}},"operations":[{"type":"image_rotate","params":{"angle":0}}]}]},{"id":"593af565-2535-4916-82c0-804ea3d818f4","type":"image","items":[{"id":"51b47a7c-202f-4397-89ef-b8b563a81686","type":"sequence","sequence_id":"70aa22e2-de92-4bfa-90e5-5b2f25539eef","operations":[{"type":"image_rotate","params":{"angle":0}}],"start_time":1800,"duration":3880,"width":300,"height":227,"position":{"x":748.4364807680388,"y":206.4229408805768},"scale":1,"asset_id":"6229f87c62a9ac0001089f69@Public@CME"}]},{"id":"05067aa8-3f4c-44bc-bcd7-735c1ba78901","type":"image","items":[{"id":"c024c37b-0513-4841-b76b-0e8680081d41","type":"sequence","sequence_id":"3eb160f4-bf6a-4cab-be43-0c1009fd8990","operations":[{"type":"image_rotate","params":{"angle":0}}],"start_time":2120,"duration":960,"width":476,"height":184,"position":{"x":480,"y":270},"scale":1,"asset_id":"622b262327c7110001552359@Public@CME"}]},{"id":"0df9040b-f5c9-4d0f-86b1-7450f57db364","type":"image","items":[{"id":"69dc199c-ef2c-41d7-9da6-63adfd359eaa","type":"sequence","sequence_id":"2a4a1690-9a71-440e-a72e-23d899698024","operations":[{"type":"image_rotate","params":{"angle":0}}],"start_time":2360,"duration":960,"width":316,"height":460,"position":{"x":480,"y":270},"scale":1,"asset_id":"622b26694c975a000166aa24@Public@CME"}]},{"id":"a890f463-b085-4b30-b9df-f9ab3daefefb","type":"audio","items":[{"id":"2fb72d99-e939-4525-8ef7-e4c46f86d1b7","start_time":1400,"duration":12160,"type":"audio","asset_id":"381921560306406981@Public@CME","section":{"from":0,"to":12160},"operations":[],"seq_key":""}]},{"id":"3082c41b-f97a-402d-b04a-9b8296972fa3","type":"video","items":[{"id":"b3319530-0e74-4f3d-949b-d5bf0cc9ec88","start_time":0,"duration":4520,"type":"video","section":{"from":4200,"to":8720},"asset_id":"63d49a150b229d00016813b4","filter_asset_id":"","width":960,"height":540,"position":{"x":480,"y":270},"operations":[{"type":"image_rotate","params":{"angle":0}},{"type":"audio_volumes","params":{"all":100}}]},{"id":"acdc2428-b771-4658-8210-23ec3ba37b4e","start_time":4520,"duration":4200,"type":"video","section":{"from":0,"to":4200},"asset_id":"63d49a150b229d00016813b4","filter_asset_id":"","width":960,"height":540,"position":{"x":480,"y":270},"operations":[{"type":"image_rotate","params":{"angle":0}},{"type":"audio_volumes","params":{"all":100}}]}]}]`)
}
