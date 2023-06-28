import subprocess
import os
import json


def execute():
    return 1


def addDevDependencies():
    dependencies = ["@jest/globals", "jest", "nodemon", "ts-jest", "typescript"]
    npm_install = ["npm", "i", "-D"]
    install_command = npm_install
    for dependency in dependencies:
        install_command.append(" ")
        install_command.append(dependency)

    subprocess.run(install_command, cwd=os.getcwd())


def addScripts():
    loaded_json = None
    with open("package.json", "r") as package_json:
        loaded_json = json.load(package_json)

    with open("package.json", "w") as package_json:
        scripts = {
            "test": "jest unit",
            "integration": "jest int",
            "dev": "nodemon",
            "prod": "node build/src/main.js",
            "build": "tsc",
        }
        loaded_json["scripts"] = scripts
        json.dump(loaded_json, package_json)


def setupTsConfig():
    with open("tsconfig.json", "w") as tsconfig_json:
        tsconfig = {
            "compilerOptions": {"outDir": "build/src"},
            "exclude": ["**/*.test.ts"],
        }
        json.dump(tsconfig, tsconfig_json)


def setupGitIgnore():
    ignored_list = ["node_modules", "build"]
    with open(".gitignore", "w") as gitignore:
        ignored_string = ""
        for ignored in ignored_list:
            ignored_string = ignored_string + ignored + "\n"
        gitignore.write(ignored_string)


def setupJestConfig():
    config = """/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'node',
};"""
    with open("jest.config.js", "w") as jestConfig:
        jestConfig.write(config)


if __name__ == "__main__":
    addDevDependencies()
    addScripts()
    setupTsConfig()
    setupGitIgnore()
    setupJestConfig()
