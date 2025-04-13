# Eternal Store DFS

Eternal Store is a DFS (Distributed File System) that distributes a file you saved on one node to many nodes to ensure that they are "eternally" preserved so that even if one node goes down, your files will still be retrievable from another node.

# Instructions

1. Run `go mod tidy`
1. Run `make run` (only run in WSL or unix terminals since Windows file systems cant create folders with the symbol ':')

# Notes

- This was meant to be a project to use to learn more in depth skills on GO.
