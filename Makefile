all: install
install: git-redmine_install

git-redmine_install:
	go install ./git-redmine
