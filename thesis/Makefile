MAIN = main
buildName = thesis-$(shell date +'%Y-%m-%dT%H:%M:%SZ')-$(shell git describe --always --tags --dirty)

all: pdf

pdf: ${MAIN}.pdf

${MAIN}.pdf: ${MAIN}.tex
	xelatex -no-pdf ${MAIN}
	# bibtex ${MAIN}
	xelatex -no-pdf ${MAIN}
	xelatex ${MAIN}

.PHONY: dev
dev: clean all

.PHONY: build
build: dev
	cp ${MAIN}.pdf $(buildName).pdf

.PHONY: clean
clean:
	git clean -Xf