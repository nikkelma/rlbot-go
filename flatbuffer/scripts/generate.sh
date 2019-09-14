#!/bin/bash

# Remove old generated files
rm -rf ./flat

# Generate go files from flatbuffer schema
./flatbuffer/bin/flatc --go ./flatbuffer/schema/rlbot.fbs

# Move generated files to intended folder
mv ./rlbot/flat ./flat

# Clean up intermediate folder
rm -r ./rlbot

# Apply go fmt
go fmt ./flat/...
