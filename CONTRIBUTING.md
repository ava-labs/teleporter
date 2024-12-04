# How to Contribute to Teleporter

## Setup

To start developing on Teleporter, you'll need Solidity >= v0.8.25. [Foundry](https://book.getfoundry.sh/getting-started/installation) is the recommended development toolkit for working with Teleporter, and it comes bundled with the required Solidity version. To run the tests and linter locally, you'll need the dependencies described in [Setup](./README.md#setup), as well as [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

## Issues

### Security

- Do not open up a GitHub issue if it relates to a security vulnerability in Teleporter, and instead refer to our [security policy](./SECURITY.md).

### Making an Issue

- Check that the issue you're filing doesn't already exist by searching under [issues](https://github.com/ava-labs/icm-contracts/issues).
- If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/ava-labs/icm-contracts/issues/new/choose). Be sure to include a _title and clear description_ with as much relevant information as possible.

## Features

- If you want to start a discussion about the development of a new feature or the modfiication of an existing one, start a thread under GitHub [discussions](https://github.com/ava-labs/icm-contracts/discussions/categories/ideas).
- Post a thread about your idea and why it should be added to Teleporter.
- Don't start working on a pull request until you've received positive feedback from the maintainers.

## Pull Request Guidelines

- Open a new GitHub pull request containing your changes.
- All commits must be [signed](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits).
- Ensure the PR description clearly describes the problem and solution, and how the change was tested. Include the relevant issue number if applicable.
- If your PR isn't ready to be reviewed just yet, you can open it as a draft to collect early feedback on your changes.
- Once the PR is ready for review, mark it as ready-for-review and request review from one of the maintainers.

### Testing

See [E2E Tests](./README.md#e2e-tests)

### Linting

- Run the Solidity and Golang linters

```sh
./scripts/lint.sh
```

### Continuous Integration (CI)

- Pull requests will generally not be approved or merged unless they pass CI.

## Other

### Do you have questions about the source code?

- Ask any question about ICM or ICM contracts under GitHub [discussions](https://github.com/ava-labs/icm-contracts/discussions/categories/q-a).
