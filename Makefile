.PHONY : build-k8s
build-k8s:
	./script/ci/build-k8s.sh

.PHONY : build-oc
build-oc:
	./script/ci/build-oc.sh

.PHONY : update-dep
update-dep:
	./script/ci/update-dep.sh