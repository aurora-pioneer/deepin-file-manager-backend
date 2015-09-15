.PHONY: mo install


POT?=file-manager-backend.pot
PKG_NAME?=file-manager-backend
PKG_VERSION?=unknown

# 1. find will sort the files.
# 2. xgettext:
#      -j will merge the non-exist messages.
#      -k disable default keyword.
#      --force-po will always generate a file even if it's not exist.
# TODO: just desktop for now.
all:
	find ../desktop -name '*.go' | xargs xgettext --force-po -k -kTr:1 -kNTr:1,2 -C --from-code=UTF-8 --package-name=$(PKG_NAME) --package-version=$(PKG_VERSION) -o $(POT)

clean:
	echo > $(POT)

generate_mo = msgfmt -o $(1).mo $(1)

mo:
	$(foreach file, $(wildcard *.po), $(call generate_mo, $(file)))

install_mo = cp $(1) /usr/share/locale/$(subst .po.mo,,$(1))/LC_MESSAGES/DFMB.mo

install:
	$(foreach file, $(wildcard *.mo), $(call install_mo,$(file)))