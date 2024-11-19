# Configuration

The configuration file is a YAML file that contains the AWS account information. The default location for the configuration file is `~/.whoiam/whoiam.yaml`.

The configuration file should have the following structure:

```yaml
accounts:
   foo-account: 123456789012 # where foo-account is the account name and 123456789012 is the account ID
   bar-account: 210987654321 # where bar-account is the account name and 210987654321 is the account ID

## Managing Config
Configruation entries can either be made by editing the configuration file directly or by using the `whoiam config` command.
Run `whoiam config --help` for more information.

