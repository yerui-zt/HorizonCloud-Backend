
Main Service:
 - User


Sub Service:
 - Identity

错误处理

rpc服务：
// 如果是 xerr 的自定义错误，则返回自定义的错误内容
// 如果是 一般错误，则只记录错误日志
// 一般的，推荐使用一般错误，不要使用自定义错误
推荐：
return nil, errors.Wrapf(err, "redis setnxex failed [key: %s]", key)
在数据库找不到记录等常用情况下，使用自定义错误，如：
DB_NOT_FOUND_ERR等

api服务：
1）api 服务想把 rpc 返回给前端友好的错误提示信息，我们想直接返回给前端不做任何处理（比如 rpc 已经返回了 “用户已存在”，api 不想做什么处理，就想把这个错误信息直接返回给前端）
return nil, errors.Wrapf(err, "login error [email: %s]", req.Email)
2）api 服务不管 rpc 返回的是什么错误信息，我就想自己再重新定义给前端返回错误信息（比如 rpc 已经返回了 “用户已存在”，api 想调用 rpc 时只要有错误我就返回给前端 “用户注册失败”）
return nil, errors.Wrapf(xerr.NewErrCodeMsg(400,"test"), "login error [email: %s]", req.Email)
