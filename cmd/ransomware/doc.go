package main

/*
	Ransomware firedrill

	Simulation of a ransomware
		- Changes wallpapers through windows registry.
		- Drops ransom note on desktop.
		- Encrypts files present on filesystem (drops sample files in a temp folder and does not destroy actual files).
		- Read credentials from Chrome (does not print or decrypt any credentials).
*/
