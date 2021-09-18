#HTTPS证书
#正式发布的时候，是需要购买正规的证书的。测试程序时，如果没有，我们可以使用openssl来生成私人的证书。

#（1）首先我们先生成证书私钥 （生成CA私钥）(无加密)
echo 1
openssl genrsa -out server.key 2048  #用于生成服务端私钥文件server.key 后面的参数2048单位是bit，是私钥的长度。


#（2）根据私钥生成公钥
echo 2
openssl rsa -in server.key -out server.key.public  #openssl生成的私钥中包含了公钥的信息，我们可以根据私钥生成公钥

#（3）根据私钥生成证书 （生成CA证书）
echo 3
openssl req -new -x509 -key server.key -out server.crt -days 365  #我们也可以根据私钥直接生成自签发的数字证书

#注意以上命令是生成在当前文件夹下的

#C  => Country
#ST => State
#L  => City
#O  => Organization
#OU => Organization Unit
#CN => Common Name (证书所请求的域名)
#emailAddress => main administrative point of contact for the certificate
