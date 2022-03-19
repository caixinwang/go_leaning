package first_test

import "testing"

/**
单元测试功能：
1. 源码⽂件以 _test 结尾：xxx_test.go
2. 测试⽅法名以 Test 开头：func TestXXX(t *testing.T) {…}
注意这里的Test后面跟的第一个一定是大写
*/
func TestFirstTry(t *testing.T) {
	t.Logf("my first try")
}
