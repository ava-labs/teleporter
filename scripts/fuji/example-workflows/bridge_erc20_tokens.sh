#!/usr/bin/env bash

set -e # Stop on first error

# Mint 0.5 ERC20 tokens on Amplify.
cast send $subnet_a_native_erc20_contract "mint(uint256)" 500000000000000000 --private-key $user_private_key --rpc-url $subnet_a_rpc_url

# Check the user's current ERC20 token balances on each chain before bridging.
native_erc20_balance=$(cast call $subnet_a_native_erc20_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_a_rpc_url)
subnet_b_bridged_token_balance=$(cast call $subnet_b_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_rpc_url)
subnet_c_bridged_token_balance=$(cast call $subnet_c_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_c_rpc_url)

echo "Native ERC20 balance on Amplify Subnet: $native_erc20_balance"
echo "Bridged ERC20 balance on Bulletin Subnet: $subnet_b_bridged_token_balance"
echo "Bridged ERC20 balance on Conduit Subnet: $subnet_c_bridged_token_balance"

# Bridge 0.4 ERC20 tokens from Amplify to Bulletin.
# Step 1: Approve the bridge contract on Amplify Subnet to spend ERC20 tokens from the user account we're using to send transactions
cast send $subnet_a_native_erc20_contract "approve(address,uint256)(bool)" $subnet_a_erc20_bridge_contract 400000000000000000 \
    --private-key $user_private_key --rpc-url $subnet_a_rpc_url
result=$(cast call $subnet_a_native_erc20_contract "allowance(address,address)(uint256)" $user_address $subnet_a_erc20_bridge_contract --rpc-url $subnet_a_rpc_url)
if [[ $result -ne 400000000000000000 ]]; then
    echo "Allowance $result"
    echo "Error approving bridge contract on Amplify Subnet to spend ERC20 from user account."
    exit 1
fi
echo "Approved the Amplify Subnet bridge contract to spend the test ERC20 token from the user account."

# Step 2: Send a transaction to the bridge contract to bridge the ERC20 tokens from Amplify Subnet to Bulletin Subnet.
total_amount=400000000000000000
primary_fee_amount=10000000000000000
secondary_fee_amount=0
cast send $subnet_a_erc20_bridge_contract "bridgeTokens(bytes32,address,address,address,uint256,uint256,uint256)" \
    $subnet_b_chain_id_hex \
    $subnet_b_erc20_bridge_contract \
    $subnet_a_native_erc20_contract \
    $user_address \
    $total_amount \
    $primary_fee_amount \
    $secondary_fee_amount \
    --private-key $user_private_key --rpc-url $subnet_a_rpc_url
echo "Sent a transaction to bridgeTokens from Amplify Subnet to Bulletin Subnet."

# Check bridged token balances
native_erc20_expected_balance=$(echo "$native_erc20_balance" '-' "$total_amount" | bc -l)
subnet_b_bridged_token_expected_balance=$(echo "$subnet_b_bridged_token_balance" '+' "$total_amount" '-' "$primary_fee_amount" '-' "$secondary_fee_amount" | bc -l)

native_erc20_balance=$(cast call $subnet_a_native_erc20_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_a_rpc_url)
subnet_b_bridged_token_balance=$(cast call $subnet_b_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_rpc_url)

retry_count=0
until [[ $native_erc20_balance = $native_erc20_expected_balance ]] && [[ $subnet_b_bridged_token_balance = $subnet_b_bridged_token_expected_balance ]]; do
    if [[ retry_count -ge 10 ]]; then
        if [[ $native_erc20_balance -ne $native_erc20_expected_balance ]]; then
            echo "Native ERC20 balance for $user_address did not match expected on Amplify Subnet."
            echo "Actual balance: $native_erc20_balance, Expected: $native_erc20_expected_balance"
        fi
        if [[ $subnet_b_bridged_token_balance -ne $subnet_b_bridged_token_expected_balance ]]; then
            echo "Bridge token balance for $user_address did not match expected on Bulletin Subnet."
            echo "Actual balance: $subnet_b_bridged_token_balance, Expected: $subnet_b_bridged_token_expected_balance"
        fi
        exit 1
    fi
    echo "Waiting for transaction to be processed. Retry count: $retry_count"
    retry_count=$((retry_count+1))
    sleep 5
    native_erc20_balance=$(cast call $subnet_a_native_erc20_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_a_rpc_url)
    subnet_b_bridged_token_balance=$(cast call $subnet_b_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_rpc_url)
done

echo "Token balances after bridging from Amplify Subnet to Bulletin Subnet:"
echo "Native ERC20 balance on Amplify Subnet: $native_erc20_balance"
echo "Bridged ERC20 balance on Bulletin Subnet: $subnet_b_bridged_token_balance"
echo "Bridge token balance matches expected."

# Bridge 0.2 ERC20 tokens from Bulletin to Conduit.
# Step 1: Approve the bridge contract on Bulletin Subnet to spent the wrapped tokens in the user account.
cast send $subnet_b_bridge_token_contract "approve(address,uint256)(bool)" $subnet_b_erc20_bridge_contract 200000000000000000 \
    --private-key $user_private_key --rpc-url $subnet_b_rpc_url
result=$(cast call $subnet_b_bridge_token_contract "allowance(address,address)(uint256)" $user_address $subnet_b_erc20_bridge_contract --rpc-url $subnet_b_rpc_url)
if [[ $result -ne 200000000000000000 ]]; then
    echo $result
    echo "Error approving bridge contract on Bulletin Subnet to spend bridged ERC20 from user account."
    exit 1
fi
echo "Approved the Bulletin Subnet bridge contract to spend the bridged ERC20 token from the user account."

# Step 2: Send a transaction to the bridge contract to bridge the ERC20 tokens from Bulletin Subnet to Conduit Subnet.
total_amount=200000000000000000
primary_fee_amount=10000000000000000
secondary_fee_amount=10000000000000000
cast send $subnet_b_erc20_bridge_contract "bridgeTokens(bytes32,address,address,address,uint256,uint256,uint256)" \
    $subnet_c_chain_id_hex \
    $subnet_c_erc20_bridge_contract \
    $subnet_b_bridge_token_contract \
    $user_address \
    $total_amount \
    $primary_fee_amount \
    $secondary_fee_amount \
    --private-key $user_private_key --rpc-url $subnet_b_rpc_url
echo "Sent a transaction to bridge tokens from Bulletin Subnet to Conduit Subnet."

# Check bridged token balances
subnet_b_bridged_token_expected_balance=$(echo "$subnet_b_bridged_token_balance" '-' "$total_amount" | bc -l)
subnet_c_bridged_token_expected_balance=$(echo "$subnet_c_bridged_token_balance" '+' "$total_amount" '-' "$primary_fee_amount" '-' "$secondary_fee_amount" | bc -l)

subnet_b_bridged_token_balance=$(cast call $subnet_b_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_rpc_url)
subnet_c_bridged_token_balance=$(cast call $subnet_c_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_c_rpc_url)

retry_count=0
until [[ $subnet_b_bridged_token_balance = $subnet_b_bridged_token_expected_balance ]] && [[ $subnet_c_bridged_token_balance = $subnet_c_bridged_token_expected_balance ]]; do
    if [[ retry_count -ge 10 ]]; then
        if [[ $subnet_b_bridged_token_balance -ne $subnet_b_bridged_token_expected_balance ]]; then
            echo "Bridge token balance for $user_address did not match expected on Bulletin Subnet."
            echo "Actual balance: $subnet_b_bridged_token_balance, Expected: $subnet_b_bridged_token_expected_balance"
        fi
        if [[ $subnet_c_bridged_token_balance -ne $subnet_c_bridged_token_expected_balance ]]; then
            echo "Bridge token balance for $user_address did not match expected on Conduit Subnet."
            echo "Actual balance: $subnet_c_bridged_token_balance, Expected: $subnet_c_bridged_token_expected_balance"
        fi
        exit 1
    fi
    echo "Waiting for transaction to be processed. Retry count: $retry_count"
    retry_count=$((retry_count+1))

    sleep 5
    subnet_b_bridged_token_balance=$(cast call $subnet_b_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_rpc_url)
    subnet_c_bridged_token_balance=$(cast call $subnet_c_bridge_token_contract "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_c_rpc_url)
done

echo "Token balances after bridging:"
echo "Native ERC20 balance on Bulletin Subnet: $subnet_b_bridged_token_balance"
echo "Bridged ERC20 balance on Conduit Subnet: $subnet_c_bridged_token_balance"
echo "Bridge token balance matches expected."

exit 0