import os
import re
import sys


def get_docker_tag(branch_name):
    if branch_name == "main":
        return "latest"
    match = re.match(r"release/(.+)", branch_name)
    if match:
        return match.group(1)
    raise ValueError(
        f"Branch name '{branch_name}' does not match expected patterns.")


def write_to_github(docker_tag):
    print(f"Writing docker tag:: {docker_tag} to github environment and outputs")
    # Write the docker tag to github environment and github outputs
    env_file = os.getenv('GITHUB_ENV')
    with open(env_file, "a") as file:
        file.write(f"DOCKER_TAG={docker_tag}")

    github_output_file = os.getenv('GITHUB_OUTPUT')
    with open(github_output_file, "a") as file:
        file.write(f"DOCKER_TAG={docker_tag}")


if __name__ == "__main__":
    try:
        branch_name = os.getenv("BRANCH_NAME")
        if not branch_name:
            write_to_github("latest")
        else:
            write_to_github(get_docker_tag(branch_name))

    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
