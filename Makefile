TARGET ?= armhf
ARCHS ?= amd64 armhf
BASE_ARCH ?= armhf


build: tmp-$(TARGET)/Dockerfile
	docker build --no-cache -t innovate/dj:$(TARGET)-latest tmp-$(TARGET)
	docker run --rm innovate/dj:$(TARGET)-latest uname -a


tmp-$(TARGET)/Dockerfile: Dockerfile $(shell find overlay-common overlay-$(TARGET))
	rm -rf tmp-$(TARGET)
	mkdir tmp-$(TARGET)
	cp Dockerfile $@
	cp -rf overlay-common tmp-$(TARGET)/
	cp -rf overlay-$(TARGET) tmp-$(TARGET)/
	for arch in $(ARCHS); do                     \
	  if [ "$$arch" != "$(TARGET)" ]; then       \
	    sed -i "/arch=$$arch/d" $@;              \
	  fi;                                        \
	done
	sed -i '/#[[:space:]]*arch=$(TARGET)/s/^#//' $@
	sed -i 's/#[[:space:]]*arch=$(TARGET)//g' $@
	cat $@
