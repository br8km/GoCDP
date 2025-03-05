# Requirements

---

- Generate Basic Files -- https://ffprofile.com/
	- pref.js -> firefox.cfg
	- policies.json

- Start New Profile -> Customize

- Config With Firefox autoconfig
	- FIREFOX-INSTALL-FOLDER
		- defaults/pref/autoconfig.js
		- firefox.cfg
		- distribution/policies.json
	- PROFILE-FOLDER
		- pref.js
		- extensions/***.xpi
			- chameleon -> settings.json
	- REMOTE-URL

- Python script testing
	- profiles|extensions create|modify|delete
	- test chameleon rules for proxy|headers
	- test fingerprintjs

## Debugging

- Check Firefox Installed Extension IDs
	- about:debugging#/runtime/this-firefox
	- about:config -> extensions.webextensions.uuids