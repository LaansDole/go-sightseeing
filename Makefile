SHELL := /bin/bash

.PHONY: add
add:
	bash -c "source ./scripts/fuzzy-add.sh"

.PHONY: commit
commit:
	@while true; do \
		echo "Enter commit type:"; \
		echo "1. feat"; \
		echo "2. fix"; \
		echo "3. docs"; \
		echo "4. chore"; \
		read -p "Choose a number from 1-4: " num; \
		case $$num in \
			1) type="feat"; break;; \
			2) type="fix"; break;; \
			3) type="docs"; break;; \
			4) type="chore"; break;; \
			*) echo "Please choose again";; \
		esac; \
	done; \
	read -p "Enter commit message: " message; \
	git commit -m "$$type: $$message"; \
	git pull; \
	git push -u

.PHONY: push
push: add commit

SHELL := /bin/bash
.PHONY: project
project:
	chmod +x ./go-projects/scripts/init_setup.sh
	read -p "Enter project name: " project; \
	bash -c "source ./go-projects/scripts/init_setup.sh -d go-projects/$$project"