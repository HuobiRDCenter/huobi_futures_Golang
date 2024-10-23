package config

// 这里的sign代表着是使用hmac256签名方法还是Ed25519签名方法。这里的accesskey和secretkey代表了公钥和私钥。如果你选择ed25519方式，那么听他们分别代表publickey和privatekey
// The sign indicates whether to use the hmac256 signature method or the Ed25519 signature method. Here accesskey and secretkey represent the public and private keys. If you choose ed25519, then listen to them stand for publickey and privatekey respectively
var SecretKey = ``

var AccessKey = ""
var Sign = "25519"
