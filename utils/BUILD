load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/bcpitutor/gosharedlibs
gazelle(name = "gazelle")

go_library(
    name = "gosharedlibs",
    srcs = ["browser.go"],
    importpath = "github.com/bcpitutor/gosharedlibs",
    visibility = ["//visibility:public"],
)
