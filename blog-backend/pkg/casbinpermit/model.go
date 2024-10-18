package casbinpermit

// var CasbinConf = `
// [request_definition]
// r = sub, obj, act

// [policy_definition]
// p = sub, obj, act

// [role_definition]
// g = _, _

// [policy_effect]
// e = some(where (p.eft == allow))

// [matchers]
// m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
// `
var CasbinConf = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

// 需要我们注意的是上面的匹配规则,如果我们使用第一种规则,那么我们在定义策略的时候,需要将对象的值写死
// Policy
// 代表策略，它表示具体的权限定义的规则是什么

// Effect
// ⽤来判断如果⼀个请求满⾜了规则，是否需要同意请求
// e = some(where (p.eft == allow))
// 这里的some表示括号中的表达式个数大于等于1就行。
// 将request和所有policy比对完之后，所有policy的策略结果(p.eft)为allow的个数>=1，整个请求的策略就是true

// 有请求，有规则，那么请求是否匹配某个规则，则是 matchers 进行判断的
// g(r.sub, p.sub) 表示 [用户-角色] 关联
