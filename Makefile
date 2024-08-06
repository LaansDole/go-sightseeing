SHELL := /bin/bash
.PHONY: push
push:
	git add .
	@while true; do \
		echo "Enter commit type:"; \
		echo "1. feat"; \
		echo "2. fix"; \
		echo "3. update"; \
		echo "4. docs"; \
		read -p "Choose a number from 1-4: " num; \
		case $$num in \
			1) type="feat"; break;; \
			2) type="fix"; break;; \
			3) type="update"; break;; \
			4) type="docs"; break;; \
			*) echo "Please choose again";; \
		esac; \
	done; \
	read -p "Enter commit message: " message; \
	git commit -m "$$type: $$message"
	git pull
	git push -u
.PHONY: project
project:
	chmod +x ./go-projects/init_setup.sh
	source ./go-projects/init_setup.sh