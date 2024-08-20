SHELL := /bin/bash

.PHONY: push
push:
	@while true; do \
		read -p "Do you want to add all files? (y/n): " yn; \
		case $$yn in \
			[Yy]*) git add .; break;; \
			[Nn]*) \
				echo "Enter the list of files to add (space-separated):"; \
				read -p "Files: " files; \
				git add $$files; \
				break;; \
			*) echo "Please answer y or n.";; \
		esac; \
	done; \
	while true; do \
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

SHELL := /bin/bash
.PHONY: project
project:
	chmod +x ./go-projects/scripts/init_setup.sh
	read -p "Enter project name: " project; \
	bash -c "source ./go-projects/scripts/init_setup.sh -d go-projects/$$project"