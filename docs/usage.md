# Usage

## Basic Commands

### Display Help

To display the help message for `whoiam`, run:

```sh
whoiam --help
```

## Run Command for an AWS Account
most organizations often have multiple AWS accounts for different environments (e.g., development, staging, production). `whoiam` allows you to run a command first checking that you are
authenticated with the expected account, before running the command. This prevents accidental deployment or destruction of resources unexpectedly via "Fat Fingers".

To run a command for a specific AWS account, use the following command:

```sh
whoiam exec --account my-account -- your command here
```
where the account is the name of the account associated with an account number in the `~/.whoiam/whoiam.yaml` configuration file.

## Retrieve AWS IAM Role Information
To simply check who you are authenticated as, run the following command:
```sh
whoiam
```
This will display the account name and account ID of the AWS account you are authenticated with.

## List AWS Accounts
To list all the AWS accounts in the configuration file, run the following command:
```sh
whoiam config view
```

## Add AWS Account
To add an AWS account to the configuration file, run the following command:
```sh
whoiam config add --name my-account --account 123456789012
```
where the name is the name of the account and account is the account ID.

## Remove AWS Account
To remove an AWS account from the configuration file, run the following command:
```sh
whoiam config delete --name my-account
```
where the name is the name of the account.