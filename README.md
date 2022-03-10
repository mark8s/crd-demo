
### k8s client 调用 crd 

参考： 
>  https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html

>  https://www.cnblogs.com/double12gzh/p/13734830.html

```yaml
go get k8s.io/client-go@v0.17.0

go get k8s.io/apimachinery@v0.17.0

go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5

controller-gen object paths=./api/types/v1alpha1/project.go

```

#### 效果
自定义一个crd和cd，并apply到k8s中。通过注册client，使用client进行cr的crud以及watch操作。
使用了`controller-gen`生成了部分代码。

