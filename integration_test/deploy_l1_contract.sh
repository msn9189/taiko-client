#!/bin/bash

source scripts/common.sh

# load l1 chain deploy contracts environment variables
source integration_test/l1_env.sh

# check taiko-mono dir path environment.
check_env "TAIKO_MONO_DIR"

cd "$TAIKO_MONO_DIR"/packages/protocol &&
 forge script script/DeployOnL1.s.sol:DeployOnL1 \
  --fork-url "$L1_NODE_HTTP_ENDPOINT" \
  --broadcast \
  --ffi \
  -vvvvv \
  --private-key "$PRIVATE_KEY" \
  --block-gas-limit 100000000
  