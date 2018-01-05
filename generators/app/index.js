'use strict';
const Generator = require('yeoman-generator');
const fs = require('fs');
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
      console.log(parsedJson.endpoints);

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/.vscode'),
        this.destinationPath(parsedJson.app_name + '/.vscode'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/docker'),
        this.destinationPath(parsedJson.app_name + '/docker'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/src'),
        this.destinationPath(parsedJson.app_name + '/src'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_dockerignore'),
        this.destinationPath(parsedJson.app_name + '/.dockerignore'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_gitignore'),
        this.destinationPath(parsedJson.app_name + '/.gitignore'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/_env'),
        this.destinationPath(parsedJson.app_name + '/.env'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/makefile'),
        this.destinationPath(parsedJson.app_name + '/makefile'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/makefile'),
        this.destinationPath(parsedJson.app_name + '/makefile'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/Gopkg.lock'),
        this.destinationPath(parsedJson.app_name + '/Gopkg.lock'),
        parsedJson
      );

      this.fs.copyTpl(
        this.templatePath('go-kit-seed/Gopkg.toml'),
        this.destinationPath(parsedJson.app_name + '/Gopkg.toml'),
        parsedJson
      );
    });
  }
  install() {
    // This.installDependencies();
  }
};
