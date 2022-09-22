build-mock-device-image:
	sudo docker buildx build -t mock-device:v0.0.1 -f demo-device/device/dockerfile . --load
	(cd smartCity && sudo docker buildx build -t smartcity:v0.0.1 . --load)

kind-load-mock-device:
	sudo kind load docker-image mock-device:v0.0.1
	sudo kind load docker-image smartcity:v0.0.1

install-shifu:
	sudo kubectl apply -f shifu/pkg/k8s/crd/install/shifu_install.yml

start-mockdevice-and-devicehsifu:
	sudo kubectl apply -R -f demo-device/conf
	sudo kubectl apply -f smartCity/frontend_deploy.yaml

delete-mockdevice-and-deviceshifu:
	sudo kubectl delete -R -f demo-device/conf
	sudo kubectl delete -f smartCity/frontend_deploy.yaml
