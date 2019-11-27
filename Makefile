NAME=opencc_jieba
DISTDIR=dist
VERSION=$(shell git describe --tags || echo "unknown version")
GOBUILD=CGO_ENABLED=0 go build -o $(DISTDIR)/$(NAME)

PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64 \

WINDOWS_ARCH_LIST = \
	windows-amd64

all: linux-amd64 darwin-amd64 windows-amd64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD)-$@ .

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD)-$@ .

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD)-$@.exe .

gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))
zip_releases=$(addsuffix .zip, $(WINDOWS_ARCH_LIST))

$(gz_releases): %.gz : %
	chmod +x $(DISTDIR)/$(NAME)-$(basename $@)
	gzip -f -S -$(VERSION).gz $(DISTDIR)/$(NAME)-$(basename $@)

$(zip_releases): %.zip : %
	zip -m -j $(DISTDIR)/$(NAME)-$(basename $@)-$(VERSION).zip $(DISTDIR)/$(NAME)-$(basename $@).exe

releases: $(gz_releases) $(zip_releases)

clean:
	rm -r $(DISTDIR)/*