format:
	python3 -m black .

check:
	python3 -m ruff check .

lock:
	python3 -m pigar generate

bundle:
	rm -rf bin
	python3 -m nuitka \
		--onefile \
		--standalone \
		--remove-output \
		--lto=yes \
		--clang \
		--assume-yes-for-downloads \
		--follow-imports \
		--output-filename=ytts \
		--enable-plugin=pylint-warnings \
		--show-progress \
		main.py
	mkdir -p bin
	mv ytts bin/
