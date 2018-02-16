'use strict';
const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);

    this.setup = {
      protoName: 'sample',
      appName: 'be-sample',
      serviceName: 'SamplePairService',
      methods: [
        {
          methodName: 'Get'
        }
      ]
    };
  }

  writing() {
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_vscode'),
      this.destinationPath(this.setup.appName + '/.vscode'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/build'),
      this.destinationPath(this.setup.appName + '/build'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/cmd'),
      this.destinationPath(this.setup.appName + '/cmd'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/internal'),
      this.destinationPath(this.setup.appName + '/internal'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pkg/endpoints'),
      this.destinationPath(this.setup.appName + '/pkg/endpoints'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pkg/pb/proto/service.proto'),
      this.destinationPath(
        this.setup.appName +
          `/pkg/pb/${this.setup.protoName}/${this.setup.protoName}.proto`
      ),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pkg/service'),
      this.destinationPath(this.setup.appName + '/pkg/service'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pkg/transport'),
      this.destinationPath(this.setup.appName + '/pkg/transport'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/pkg/utils'),
      this.destinationPath(this.setup.appName + '/pkg/utils'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_dockerignore'),
      this.destinationPath(this.setup.appName + '/.dockerignore'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_gitignore'),
      this.destinationPath(this.setup.appName + '/.gitignore'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/_env'),
      this.destinationPath(this.setup.appName + '/.env'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/glide.yaml'),
      this.destinationPath(this.setup.appName + '/glide.yaml'),
      this.setup
    );
    this.fs.copyTpl(
      this.templatePath('go-kit-seed/makefile'),
      this.destinationPath(this.setup.appName + '/makefile'),
      this.setup
    );
  }
  end() {}
};
