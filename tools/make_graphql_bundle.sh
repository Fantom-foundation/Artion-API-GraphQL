#!/bin/sh
#
# Generate bundled GraphQL schema from the schema definition.
# The schema definition is split by modules and components,
# but we want to have a bundled single schema file for the GraphQL
# package to work with.
#
# Usage: make_graphql_bundle.sh <target file> <source path> <source path> ...

# the first argument is the output file to be generated
TARGET_FILE="$1"
shift 1
echo "Schema ${TARGET_FILE} is generated."

# make sure we have the target directory
TARGET_DIR=$(dirname "${TARGET_FILE}")
[ ! -d "${TARGET_DIR}" ] && mkdir -p "${TARGET_DIR}"

# remove previous build if any
[ -f "${TARGET_FILE}" ] && rm -f "${TARGET_FILE}"

# loop source files and process them
COUNT=1
MAX=$#
while [ ${COUNT} -le ${MAX} ]; do
  {
    # base schema file inside?
    [ -f "${1}/schema.graphql" ] && cat "${1}/schema.graphql" && echo ""

    # export all the other files
    find "${1}" -name '*.graphql' ! -name 'schema.graphql' -print0 | xargs -0 -I{} sh -c "cat {}; echo ''"
  } >>"${TARGET_FILE}"

  # advance counter and shift the arguments
  COUNT=$((COUNT + 1))
  shift 1
done

# we are done
exit 0
