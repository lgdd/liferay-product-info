#!/bin/bash

container_name=get-dxp-dev-trial
modules_path=/opt/liferay/osgi/modules
trial_file_pattern=*.xml
trial_file=""

function destroyContainer {
  docker stop "$container_name" >/dev/null 2>&1
  docker rm "$container_name" >/dev/null 2>&1
}

if command -v docker >/dev/null 2>&1; then
  destroyContainer
  echo "Starting liferay/dxp:latest..."
  docker run --name $container_name -d liferay/dxp:latest >/dev/null 2>&1
  echo "Waiting for the trial DXP license to be deployed..."
  while [[ -z "$trial_file" ]]; do
    trial_file=$(docker exec "$container_name" find "$modules_path" -maxdepth 1 -name "$trial_file_pattern" -print -quit 2>/dev/null)

    if [[ -z "$trial_file" ]]; then
      sleep 5
    fi
  done
  echo "Found: '$trial_file'"
  docker cp "$container_name":"$trial_file" .
  trial_file_name=$(echo $trial_file | cut -d'/' -f6)
  mv $trial_file_name trial.xml
  echo "Stopping liferay/dxp:latest..."
  destroyContainer
  echo "Done!"
else
  echo "Error: docker command not found. Please install it and try again."
  exit 1
fi