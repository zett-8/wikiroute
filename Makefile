DATE ?= 20250301

download-pages:
	curl -O https://dumps.wikimedia.org/jawiki/$(DATE)/jawiki-$(DATE)-page.sql.gz

download-pagelinks:
	curl -O https://dumps.wikimedia.org/jawiki/$(DATE)/jawiki-$(DATE)-pagelinks.sql.gz

download-linktargets:
	curl -O https://dumps.wikimedia.org/jawiki/$(DATE)/jawiki-$(DATE)-linktarget.sql.gz

download-all:
	make download-pages
	make download-pagelinks
	make download-linktargets

generate-linktargets:
	go run ./cmd/generate --type=linktargets

generate-pages:
	go run ./cmd/generate --type=pages

generate-pagelinks:
	go run ./cmd/generate --type=pagelinks

generate-all:
	make generate-linktargets
	make generate-pages
	make generate-pagelinks
	