Terraform `statevar` Provider
===================================

The `statevar` provider allows you to store arbitrary strings in the terraform state. It is a fork of [terraform-provider-secret](https://github.com/tweag/terraform-provider-secret) but adds the ability to assign a default value and is not designed for storing secrets. If you wish to store secrets then keep using [terraform-provider-secret](https://github.com/tweag/terraform-provider-secret).

Terraform workspaces are great for sharing similar environments, for example you could have a `prod` workspace, a `staging` workspace and even a `david` workspace, each defining a separate but similar infrastructure. However these infrastructures likely have variables that define the differences. Traditionally you store these variables in a `*.tfvars` file however you now have the problem of sharing this file. What would be great is if you could store these values directly in the terraform state, that way you would only have to switch workspaces to work on a different infrastructure.

With the `statevar` provider you can create a string resource that is stored in the terraform state and can have an optional default value. If you want to override this value you `terraform import` a new value.

Lets say I want to store an EC2 instance size for my machines:
```terraform
resource "statevar_string" "ec2_instance_type" {
  default = "t2.small"
}
```

I can now reference this as `"${statevar_string.ec2_instance_type.value}"` in my `aws_instance` definition.

A `t2.small` is fine for development environments but I want something a bit bigger in staging:
```
terraform state rm statevar_string.ec2_instance_type
terraform import statevar_string.ec2_instance_type "t2.large"
```
An unfortunate limitation of terraform is that when I import, this will set the `default` value in state to `""`. This is no big deal but this means next time I `terraform apply` I will get:
```
  ~ statevar_string.ec2_instance_type
      default: "" => "t2.small"
```
This does not change the `value` you imported to state, it just corrects the default value in state to the one you have configured.

Now lets say I want to store the environment name:
```
resource "statevar_string" "environment" {}
```
I haven't provided a `default` value because we don't want clashing names, so we will need to import our value:
```
terraform import statevar_string.environment "staging"
```

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

## How to install

### Using pre-built binary

1. Download the binary from the project [releases page](https://github.com/shmish111/terraform-provider-statevar/releases/latest)
2. Extract provider binary from tar file.
3. Copy to `$PATH` or the `~/.terraform.d/plugins` directory so Terraform can find it.

### Building from source

1. Follow these [instructions](https://golang.org/doc/install) to setup a Golang development environment.
2. Use `go get` to pull down this repository and compile the binary:

```
go get -u -v github.com/shmish111/terraform-provider-statevar
```

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/shmish111/terraform-provider-statevar`

```sh
$ git clone git@github.com:shmish111/terraform-provider-statevar $GOPATH/src/github.com/shmish111/terraform-provider-statevar
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/shmish111/terraform-provider-statevar
$ make build
```

Using the provider
----------------------

### `secret_resource`

**Schema**:

* `value`, string: Returns the value of the string
* `default`, string: The default value if no value is imported

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-statevar
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## License

This work is licensed under the Mozilla Public License 2.0. See
[LICENSE](LICENSE) for more details.