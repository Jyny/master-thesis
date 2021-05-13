MAIN = main

all: pdf

pdf: ${MAIN}.pdf

${MAIN}.pdf: ${MAIN}.tex
	xelatex -no-pdf ${MAIN}
	# bibtex ${MAIN}
	xelatex -no-pdf ${MAIN}
	xelatex ${MAIN}

.PHONY: dev
dev: clean all

.PHONY: clean
clean:
	git clean -Xf