# Ec2ls

List AWS EC2 instances.

## Usage

Configure your aws credentials in the manner of [aws-cli](https://github.com/aws/aws-cli), and type:

    $ ec2ls

To specify profile in ~/.aws/credentials:

    $ ec2ls -p [profile name]

To add more information:

    $ ec2ls -k launch_time,image_id

To show available keys:

    $ ec2ls -l

## Credentials

This tools uses AWS SDK for Ruby, so credentials are loaded [as described in the library document](https://github.com/aws/aws-sdk-ruby#credentials) unless you specify any profile.

## Installation

This file is not distributed in rubygems.
You should install by following steps:

```bash
$ git clone https://github.com/nsaeki/ec2ls.git
$ rake install
```

## Contributing

1. Fork it ( https://github.com/nsaeki/ec2ls/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
