'use strict';
const Generator = require('yeoman-generator');

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);

    // This makes `config` a required argument.
    this.argument('appName', { type: String, required: true });

    this.serviceName = this.options.appName
      .split('-')
      .map(element => {
        return capitalizeFirstLetter(element);
      })
      .join('');

    this.userInput = {
      appName: this.options.appName,
      org: 'awpc',
      serviceName: this.serviceName
    };
  }

  writing() {
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/.vscode'),
      this.destinationPath(this.userInput.appName + '/.vscode'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pb/service.proto'),
      this.destinationPath(
        this.userInput.appName + '/pb/' + this.userInput.appName + '.proto'
      ),
      this.userInput
    );

    this.fs.copy(
      this.templatePath('go-kit-seed/pkg'),
      this.destinationPath(this.userInput.appName + '/pkg')
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/docker'),
      this.destinationPath(this.userInput.appName + '/docker'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/src'),
      this.destinationPath(this.userInput.appName + '/src'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/internal'),
      this.destinationPath(this.userInput.appName + '/internal'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_dockerignore'),
      this.destinationPath(this.userInput.appName + '/.dockerignore'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_gitignore'),
      this.destinationPath(this.userInput.appName + '/.gitignore'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_env'),
      this.destinationPath(this.userInput.appName + '/.env'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/README.md'),
      this.destinationPath(this.userInput.appName + '/README.md'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/makefile'),
      this.destinationPath(this.userInput.appName + '/makefile'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/Gopkg.lock'),
      this.destinationPath(this.userInput.appName + '/Gopkg.lock'),
      this.userInput
    );

    this.fs.copyTpl(
      this.templatePath('go-kit-seed/Gopkg.toml'),
      this.destinationPath(this.userInput.appName + '/Gopkg.toml'),
      this.userInput
    );
  }
  end() {}
};
