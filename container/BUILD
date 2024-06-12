load("@rules_oci//oci:defs.bzl", "oci_image", "oci_tarball")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

pkg_tar(
    name = "server_layer",
    srcs = ["//server"],
)

oci_image(
    name = "container_image",
    base = "@distroless_base",
    entrypoint = ["/server"],
    tars = [":server_layer"],
)

oci_tarball(
    name = "container",
    image = ":container_image",
    repo_tags = [
        "popovicu.com/foo:latest",
    ],
)
