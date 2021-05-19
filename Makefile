VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: dev

versioning:
	cd imls-playbook ; \
		sed 's/<<VERSION>>/$(VERSION)/g' Makefile.in > Makefile ; \
		cd ..


ifeq ($(shell git describe --tags --abbrev=0),$(VERSION))
release:
	@echo "Version needs to be updated from " $(VERSION)
dev:
	@echo "Version needs to be updated from " $(VERSION)
else	
# make VERSION=v1.2.3 release
release: versioning
	@echo $(VERSION) > prod-version.txt

# make dev
dev: versioning
	@echo $(VERSION) > dev-version.txt
	cd imls-raspberry-pi ; make dev ; cd ..
endif

