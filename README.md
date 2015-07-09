# Ec2ls

List AWS EC2 instances.

## Usage

Configure your AWS credentials in the manner of [aws-cli](https://github.com/aws/aws-cli), and type:

    $ ec2ls

To specify profile in ~/.aws/credentials:

    $ ec2ls -p [profile name]

To add more information:

    $ ec2ls -k launch_time,image_id

To show available keys:

    $ ec2ls -l

## Credentials

This tool uses [AWS SDK for Go](https://github.com/aws/aws-sdk-go). Your credentials are loaded by the library default priority unless you specify a profile name in AWS credentials.

## Installation

By following steps:

```bash
$ go get github.com/nsaeki/ec2ls
```

## Contributing

1. Fork it ( https://github.com/nsaeki/ec2ls/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
