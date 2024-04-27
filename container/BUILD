load("@rules_pkg//pkg:tar.bzl", "pkg_tar")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_tarball")

pkg_tar(
    name = "server_layer",
    srcs = ["//server"],
)

oci_image(
    name = "container_image",
    base = "@distroless_base",
    tars = [":server_layer"],
    entrypoint = ["/server"],
)

oci_tarball(
    name = "container",
    image = ":container_image",
    repo_tags = [
        "popovicu.com/foo:latest",
    ],
)
