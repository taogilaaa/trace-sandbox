# Path to this plugin, Note this must be an abolsute path on Windows
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"

# Path to the grpc_node_plugin
PROTOC_GEN_GRPC_PATH="./node_modules/.bin/grpc_tools_node_protoc_plugin"

# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="./src/generated"

protoc \
  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
  --plugin="protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH}" \
  --js_out="import_style=commonjs,binary:${OUT_DIR}" \
  --ts_out="service=grpc-node:${OUT_DIR}" \
  --grpc_out="${OUT_DIR}" \
  -I ../api/protobuf-spec sandbox_sales_v1.proto
