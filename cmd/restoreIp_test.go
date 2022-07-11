package cmd 

import "testing"

func TestRestoreIp(t *testing.T) {
    tables := []struct {
        s string
        result []string
    }{
        {"25525511135",[]string{"255.255.11.135","255.255.111.35"}},
        {"0000", []string{"0.0.0.0"}},
        {"101023", []string{"1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3", "101.0.2.3"}},
        {"abcasdasd", []string{}},
    }
    for _, table := range tables {
        total := RestoreIp(table.s)
        if testEq(total, table.result) == false {
            t.Errorf("Incorrect for %v, expected %v, got %v\n", table.s, table.result, total)
        }
    }
}

func testEq(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
