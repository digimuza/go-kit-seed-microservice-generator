'use strict';
const Generator = require('yeoman-generator');
const fs = require('fs');
// Const { execSync } = require('child_process');
module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);

    // This makes `config` a required argument.
    this.argument('config', { type: String, required: true });

    // And you can then access it later; e.g.
    this.log(this.options.config);
  }

  prompting() {
    const prompts = [];

    return this.prompt(prompts).then(props => {
      // To access props later use this.props.someAnswer;
      this.props = props;

      // /home/andrius/Desktop/proj/go-kit-generator.json

      this.dasd = this.props.someAnswer;
    });
  }

  writing() {
    fs.readFile(this.options.config, 'utf8', (err, data) => {
      if (err) {
        return console.log(err);
      }

      const parsedJson = JSON.parse(data);
      this.parsedJson = parsedJson;

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/.vscode'),
        this.destinationPath(parsedJson.appName + '/.vscode'),
        parsedJson
      );

      if (parsedJson.grpc) {
        this.fs.copyTpl(
          this.templatePath('go-kit-seed/pb/service.proto'),
          this.destinationPath(
            parsedJson.appName + '/pb/' + parsedJson.appName + '.proto'
          ),
          parsedJson
        );

        this.fs.copy(
          this.templatePath('go-kit-seed/protoc'),
          this.destinationPath(parsedJson.appName + '/protoc')
        );
      }
      this.fs.copy(
        this.templatePath('go-kit-seed/pkg'),
        this.destinationPath(parsedJson.appName + '/pkg')
      );
      this.fs.copyTpl(
        this.templatePath('go-kit-seed/docker'),
        this.destinationPath(parsedJson.appName + '/docker'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/src'),
        this.destinationPath(parsedJson.appName + '/src'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_dockerignore'),
        this.destinationPath(parsedJson.appName + '/.dockerignore'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_gitignore'),
        this.destinationPath(parsedJson.appName + '/.gitignore'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_env'),
        this.destinationPath(parsedJson.appName + '/.env'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/README.md'),
        this.destinationPath(parsedJson.appName + '/README.md'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/makefile'),
        this.destinationPath(parsedJson.appName + '/makefile'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/makefile'),
        this.destinationPath(parsedJson.appName + '/makefile'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/Gopkg.lock'),
        this.destinationPath(parsedJson.appName + '/Gopkg.lock'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/Gopkg.toml'),
        this.destinationPath(parsedJson.appName + '/Gopkg.toml'),
        parsedJson
      );
    });
  }
  end() {}
};
