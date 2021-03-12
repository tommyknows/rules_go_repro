def _template_impl(ctx):
    out = ctx.actions.declare_file(ctx.attr.destination_name)
    args = ctx.actions.args()
    args.add("-template", ctx.file.template)
    args.add("-out", out)

    ctx.actions.run(
        executable = ctx.executable._templater,
        arguments = [args],
        outputs = [out],
        inputs = [ctx.file.template],
    )

    return [
        DefaultInfo(
            files = depset([out]),
        ),
    ]

template = rule(
    implementation = _template_impl,
    attrs = {
        "template": attr.label(
            allow_single_file = True,
        ),
        "destination_name": attr.string(),
        "_templater": attr.label(
            executable = True,
            cfg = "host",
            allow_single_file = True,
            default = Label("//build/templater:templater"),
        ),
    },
)
