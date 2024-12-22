import os
import re
import sys

def get_docker_tag(branch_name):
    if branch_name == "main":
        return "latest"
    match = re.match(r"release/(.+)", branch_name)
    if match:
        return match.group(1)
    raise ValueError(f"Branch name '{branch_name}' does not match expected patterns.")

if __name__ == "__main__":
    try:
        branch_name = os.getenv("BRANCH_NAME")
        if not branch_name:
            raise EnvironmentError("BRANCH_NAME environment variable is not set.")
        docker_tag = get_docker_tag(branch_name)
        print(f"::set-output name=docker_tag::{docker_tag}")
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)