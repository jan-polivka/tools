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
        print(loaded_json)
        json.dump(loaded_json, package_json)


if __name__ == "__main__":
    execute()
