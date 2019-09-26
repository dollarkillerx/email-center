/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:55 2019-09-25
 */
package defs

type SeedEmail struct {
	Email string `json:"email" form:"email"` // 输入参考 a@a.com,b@b.com,v@c.com
	Head  string `json:"head" form:"head"`
	Body  string `json:"body" form:"body"`
	Key   string `json:"key" form:"key"`
}
