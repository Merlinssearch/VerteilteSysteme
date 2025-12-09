#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2025 Maxim Ott <maxim.ott@informatik.hs-fulda.de>
# SPDX-FileCopyrightText: 2025 Maximilian Krönung <maximilian.kroenung@informatik.hs-fulda.de>
#
# SPDX-License-Identifier: GPL-3.0-or-later

reuse_lint() {
	if ! command -v reuse >/dev/null 2>&1; then
		printf "Reuse tool is not installed.\nSkip checking copyright information"
		return 0
	fi
	reuse lint
}

format() {
	if ! command -v go >/dev/null 2>&1; then
		echo "Go is not installed."
		return 1
	fi

	go fmt $(go list ./... | grep -v /vendor/) || return "$?"
	go vet $(go list ./... | grep -v /vendor/) || return "$?"
	go test -race $(go list ./... | grep -v /vendor/) || return "$?"
}

go_build() {
	mkdir -p target >/dev/null 2>&1
	go build -o target/controller ./cmd/controller || return "$?"
	go build -o target/robot ./cmd/robot || return "$?"
}

container_build() {
	local REGISTRY="$1"
	local PROJECT_NAMESPACE="$2"
	local PROJECT_NAME="$3"
	local COMMIT_SHORT_SHA="$(git rev-parse --short HEAD)"

	docker build \
		-t "$REGISTRY/$PROJECT_NAMESPACE/$PROJECT_NAME:controller" \
		-t "$REGISTRY/$PROJECT_NAMESPACE/$PROJECT_NAME:controller-$COMMIT_SHORT_SHA" \
		-f "docker/controller.Dockerfile" . || return "$?"


        docker build \
                -t "$REGISTRY/$PROJECT_NAMESPACE/$PROJECT_NAME:robot" \
                -t "$REGISTRY/$PROJECT_NAMESPACE/$PROJECT_NAME:robot-$COMMIT_SHORT_SHA" \
                -f "docker/robot.Dockerfile" . || return "$?"
}

main() {
	local REGISTRY="git-ce.rwth-aachen.de"
	local PROJECT_NAMESPACE="hfd-distributed-systems-ws2526"
	local PROJECT_NAME="group11"

	reuse_lint || return "$?"
	format || return "$?"
	go_build || return "$?"
	container_build "$REGISTRY" "$PROJECT_NAMESPACE" "$PROJECT_NAME" || return "$?"
}

main "$@"
