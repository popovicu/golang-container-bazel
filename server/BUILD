load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server_lib",
    srcs = ["server.go"],
    importpath = "github.com/developer-guy/golang-container-bazel/server",
    visibility = ["//visibility:private"],
    deps = ["@com_github_fatih_color//:color"],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    goarch = "arm64",
    goos = "linux",
    visibility = ["//visibility:public"],
)
