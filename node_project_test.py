import unittest
import node_project
import os
import shutil
import json

class NodeProjectTest(unittest.TestCase):

    @unittest.skip("takes too long")
    def testAddDevDependencies(self):
        current_dir = os.getcwd()
        test_directory = "test"
        os.mkdir(test_directory)
        os.chdir(test_directory)

        try:
            node_project.addDevDependencies()
        
            self.assertTrue(os.path.exists("node_modules"))
            self.assertTrue(os.path.exists("package.json"))
        finally:
            os.chdir(current_dir)
            shutil.rmtree(test_directory)
            
    def testAddScripts(self):
        current_dir = os.getcwd()
        test_directory = "test"
        test_object = {"test":"test"}
        os.mkdir(test_directory)
        os.chdir(test_directory)
        with open('package.json', 'w') as package_json:
            json.dump(test_object, package_json)
        
        try:
            node_project.addScripts()

            with open('package.json', 'r') as package_json:
                result = json.load(package_json)
                keys = result.keys()
                self.assertTrue("test" in keys)
                self.assertTrue("scripts" in keys)
                self.assertTrue("integration" in result["scripts"])
        finally:
            os.chdir(current_dir)
            shutil.rmtree(test_directory)

