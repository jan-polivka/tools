import unittest
import node_project
import os
import shutil
import json

class NodeProjectTest(unittest.TestCase):
    
    current_dir = None
    test_directory = "test"
    
    def setUp(self):
        self.current_dir = os.getcwd()
        os.mkdir(self.test_directory)
        os.chdir(self.test_directory)
        
    def tearDown(self):
        os.chdir(self.current_dir)
        shutil.rmtree(self.test_directory)

    @unittest.skip("takes too long")
    def testAddDevDependencies(self):
        node_project.addDevDependencies()
        
        self.assertTrue(os.path.exists("node_modules"))
        self.assertTrue(os.path.exists("package.json"))
            
    def testAddScripts(self):
        test_object = {"test":"test"}
        with open('package.json', 'w') as package_json:
            json.dump(test_object, package_json)
        
        node_project.addScripts()

        with open('package.json', 'r') as package_json:
            result = json.load(package_json)
            keys = result.keys()
            self.assertTrue("test" in keys)
            self.assertTrue("scripts" in keys)
            self.assertTrue("integration" in result["scripts"])

    def testSetupTsConfig(self):
        node_project.setupTsConfig()

        with open('tsconfig.json', 'r') as tsconfig_json:
            result = json.load(tsconfig_json)
            keys = result.keys()
            self.assertTrue("compilerOptions" in keys)
            self.assertTrue("exclude" in keys)
            
    def testSetupGitIgnore(self):
        node_project.setupGitIgnore()

        with open('.gitignore', 'r') as gitignore:
            result = gitignore.read()

            self.assertNotEqual(-1, result.find("node_modules"))
            self.assertNotEqual(-1, result.find("build"))
            
        