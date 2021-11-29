import json
  

with open("config.txt", "r") as f:
    content = f.read()

config_j = json.loads(content)
for item in config_j["schedule"]:
    print("---")
    print("apiVersion: v1")
    print("kind: query")
    print("spec:")
    print("  description: {}".format(config_j["schedule"][item]["description"]))
    print("  name: {}".format(item))
    print("  query: {}".format(config_j["schedule"][item]["query"]))