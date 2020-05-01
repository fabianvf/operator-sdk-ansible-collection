module github.com/operator-framework/operator-sdk-ansible-collection

go 1.13

require (
	github.com/operator-framework/api v0.3.0
	github.com/operator-framework/operator-registry v1.6.2-0.20200330184612-11867930adb5
	github.com/operator-framework/operator-sdk v0.17.0
	github.com/sirupsen/logrus v1.5.0
)

replace k8s.io/client-go => k8s.io/client-go v0.18.2
