package shim

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessStringsSlice(t *testing.T) {
	type args struct {
		strs   []string
		filter func(string) bool
		fn     func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{
			strs: []string{"1", "2"},
		}, []string{"1", "2"}},
		{"t2", args{
			strs: []string{"1", "2"},
			filter: func(s string) bool {
				return s == "1"
			},
			fn: nil,
		}, []string{"2"}},
		{"t3", args{
			strs: []string{"1", "2"},
			fn: func(s string) string {
				return fmt.Sprintf("0x%v", s)
			},
		}, []string{"0x1", "0x2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ProcessStringsSlice(tt.args.strs, tt.args.filter, tt.args.fn), "ProcessStringsSlice(%v, %v, %v)", tt.args.strs, tt.args.filter, tt.args.fn)
		})
	}
}

func TestParseStrIDToUint(t *testing.T) {
	t.Logf("%T, %[1]v", ParseStrIDToUint("1", uint64(0)))
	t.Logf("%T, %[1]v", ParseStrIDToUint("1", uint64(0)))
	t.Logf("%T, %[1]v", ParseStrIDToUint("1", uint(0)))
}

// func TestGetMapKeyValue(t *testing.T) {
// 	m1 := map[string]string{
// 		"a": "va",
// 		"b": "vb",
// 	}
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m1, "a", ""))
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m1, "x", "-"))
//
// 	m2 := map[string]int{
// 		"a": 100,
// 		"b": 200,
// 	}
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m2, "a", 0))
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m2, "x", 0))
//
// 	m3 := map[string]uint64{
// 		"a": 1000,
// 		"b": 2000,
// 	}
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m3, "a", uint64(0)))
// 	t.Logf("%T, %[1]v", GetMapKeyValue(m3, "x", uint64(0)))
// }

func TestGetMapKeyValue1(t *testing.T) {

	s := `{
    "productCode": "p_trade_t_s",
    "sv_trade_t_s_software_test1": 1,
    "subProductCode": "sp_trade_t_s_cloudapp",
    "timeSpan": 1,
    "resourceId": "shop-173103811172FY97GTS",
    "originate": "qcloud.directEnter.home",
    "curDeadline": "2024-12-08 11:55:21",
    "productInfo": [
        {
            "name": "version",
            "value": "basic"
        },
        {
            "name": "scale",
            "value": "single"
        }
    ],
    "extparam": {
        "trafficParams": "***%24%3Btimestamp%3D1732525847000%3Bfrom_type%3Dserver%3Btrack%3D97525dd5-43da-4815-96a3-b76d6f9fe4c0%3B%24***",
        "useragent": "Mozilla%2F5.0%20(Macintosh%3B%20Intel%20Mac%20OS%20X%2010_15_7)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F131.0.0.0%20Safari%2F537.36",
        "token": "MnxmsKz13u8cY3p4dA9DFcKy5c6Ganh51ddbc037a47f9c29642fc6d60d3f30eaBAKmLy8RdzhVd-pYvQo51B3-6-ek0hG5f74CA0l0JyWuG3NDpg2RwMdHxhgmyUNMbEGhSJcV-krpihvy2XsSVWJO3A-04xDk5gbjBvIgaRp20Ah6fKqi2M0Mhtf9N-YSQwFngiOlNxCOxncYPJILIiA7eTn-vKPm0230nRXylVJqIIZRd8f4aCvuFDCXf1Fg",
        "referer": "https://console.cloud.tencent.com/account/renewal"
    },
    "timeUnit": "m",
    "goodsNum": 1
}`

	var gm map[string]any
	err := json.Unmarshal([]byte(s), &gm)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", gm)

	t.Logf("%T, %[1]v", GetMapKeyValue(gm, "timeSpan", 0))
	t.Logf("%T, %[1]v", GetMapKeyValue(gm, "timeUnit", "-"))
	t.Logf("%T, %[1]v", GetMapKeyValue(gm, "goodsNum", uint32(0)))

}

func TestGenRandomLengthStr(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{5}},
	}
	for _, tt := range tests {
		t.Logf("%s", GenRandomLengthStr(tt.args.length))
	}
}

func TestHashStringToUint64(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"t1", args{"a", 32}, 2248273036},
		{"t1", args{"b", 32}, 2248274341},
		{"t1", args{"ab", 32}, 3041237098},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := HashStringToUint64(tt.args.s, tt.args.n)
			if v != tt.want {
				t.Error("HashStringToUint64 err")
			}
		})
	}
}
