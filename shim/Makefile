VERSION_FILE = version

.PHONY: tag push

tag:
	@current_version=$$(cat $(VERSION_FILE)); \
	version_number=$${current_version#v}; \
	IFS=. read major minor patch <<< "$$version_number"; \
	new_patch=$$((patch + 1)); \
	new_version="v$${major}.$${minor}.$${new_patch}"; \
	echo "$$new_version" > $(VERSION_FILE); \
	git add .; \
	git commit -m "Bump tag version to $$new_version" > /dev/null; \
	git tag -a "$${new_version}" -m "--feat $${new_version}"; \
	echo "Created tag: $${new_version}"

push:
	@current_version=$$(cat $(VERSION_FILE)); \
	git push origin $$current_version