/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:38 2019-09-25
 */
package test

import (
	"log"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	str := "he,sc,ps,go,"
	if str[len(str)-1:] == "," {
		str = str[:len(str)-1]
	}

	split := strings.Split(str, ",")
	log.Println(split)
	log.Println(len(split))
}
