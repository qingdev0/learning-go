TOPIC := "Available Recipes for Go Learning:\n"

_default:
    @just --list --unsorted --list-prefix "✨ " --list-heading '{{TOPIC}}'

commit:
    #!/usr/bin/env bash
    git status
    read -p "Enter commit message: " message
    git commit -m "$message"
    git push

# Run slice examples with interactive selection
slices:
    @go run main.go types/slices
