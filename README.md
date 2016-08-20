tfvars-out

This is a simple go application that you pipe the "outputs" of a `<name>.tfstate` file to, and it will turn them into a valid `<name>.tfvars` file. I don't know why this isn't built-in to Terraform.

Example Usage:

    terraform output -state=example.tfstate -no-color | tfvars-out > example.tfvars

Build a binary for your system, and send stdin input to it and it will output on stdout.

simple.
