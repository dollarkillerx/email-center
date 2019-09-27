/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:38 2019-09-25
 */
package test

import (
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"log"
	"strings"
	"sync"
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

// 压力测试
func TestAb(t *testing.T) {
	url := "https://translate.dollarkiller.com/translate?sl=&tl=fr&text=hello"
	gcf(url,5000)
}

func gcf(str string,num int)  {
	sy := sync.WaitGroup{}
	for i:=0;i<num;i++ {
		sy.Add(1)
		go func() {
			defer sy.Done()
			bytes, e := httplib.EuUserGet(str)
			if e != nil {
				clog.PrintWa(e)
			}else {
				clog.Println(string(bytes))
			}
		}()
	}

	sy.Wait()
}
