DEPS_FILES := \
	X509-ML-KEM-2025.asn \
	./example/ML-KEM-1024.crt \
	./example/ML-KEM-1024.crt.txt \
	./example/ML-KEM-1024-seed.priv \
	./example/ML-KEM-1024-expanded.priv \
	./example/ML-KEM-1024-both.priv \
	./example/ML-KEM-1024-seed.priv.txt \
	./example/ML-KEM-1024-expanded.priv.txt \
	./example/ML-KEM-1024-both.priv.txt \
	./example/ML-KEM-1024.pub \
	./example/ML-KEM-1024.pub.txt \
	./example/ML-KEM-512.crt \
	./example/ML-KEM-512.crt.txt \
	./example/ML-KEM-512-seed.priv \
	./example/ML-KEM-512-expanded.priv \
	./example/ML-KEM-512-both.priv \
	./example/ML-KEM-512-seed.priv.txt \
	./example/ML-KEM-512-expanded.priv.txt \
	./example/ML-KEM-512-both.priv.txt \
	./example/ML-KEM-512.pub \
	./example/ML-KEM-512.pub.txt \
	./example/ML-KEM-768.crt \
	./example/ML-KEM-768.crt.txt \
	./example/ML-KEM-768-seed.priv \
	./example/ML-KEM-768-expanded.priv \
	./example/ML-KEM-768-both.priv \
	./example/ML-KEM-768-seed.priv.txt \
	./example/ML-KEM-768-expanded.priv.txt \
	./example/ML-KEM-768-both.priv.txt \
	./example/ML-KEM-768.pub \
	./example/ML-KEM-768.pub.txt \

LIBDIR := lib
include $(LIBDIR)/main.mk

$(LIBDIR)/main.mk:
ifneq (,$(shell grep "path *= *$(LIBDIR)" .gitmodules 2>/dev/null))
	git submodule sync
	git submodule update $(CLONE_ARGS) --init
else
	git clone -q --depth 10 $(CLONE_ARGS) \
	    -b main https://github.com/martinthomson/i-d-template $(LIBDIR)
endif
