local config = import 'default.jsonnet';

config {
  'cronos_777-1'+: {
    'account-prefix': 'crc',
    'coin-type': 60,
    'precompiled-contract-address': '0x0000000000000000000000000000000000000065',
    key_name: 'signer1',
    accounts: super.accounts[:std.length(super.accounts) - 1] + [super.accounts[std.length(super.accounts) - 1] {
      coins: super.coins + ',100000000000ibcfee',
    }],
    'app-config'+: {
      'index-events': super['index-events'] + ['message.action'],
    },
    genesis+: {
      app_state+: {
        feemarket+: {
          params+: {
            no_base_fee: true,
            base_fee: '0',
          },
        },
      },
    },
  },
  'chainmain-1': {
    cmd: 'chain-maind',
    'start-flags': '--trace',
    'account-prefix': 'cro',
    'coin-type': 394,
    'app-config': {
      'minimum-gas-prices': '500basecro',
    },
    validators: [
      {
        coins: '2234240000000000000cro',
        staked: '10000000000000cro',
        mnemonic: '${VALIDATOR1_MNEMONIC}',
        client_config: {
          'broadcast-mode': 'block',
        },
        base_port: 26800,
      },
      {
        coins: '987870000000000000cro',
        staked: '20000000000000cro',
        mnemonic: '${VALIDATOR2_MNEMONIC}',
        client_config: {
          'broadcast-mode': 'block',
        },
        base_port: 26810,
      },
    ],
    accounts: [
      {
        name: 'community',
        coins: '10000000000000cro',
        mnemonic: '${COMMUNITY_MNEMONIC}',
      },
      {
        name: 'relayer',
        coins: '10000000000000cro',
        mnemonic: '${SIGNER1_MNEMONIC}',
      },
      {
        name: 'signer2',
        coins: '10000000000000cro',
        mnemonic: '${SIGNER2_MNEMONIC}',
      },
    ],
    genesis: {
      app_state: {
        staking: {
          params: {
            unbonding_time: '1814400s',
          },
        },
        gov: {
          voting_params: {
            voting_period: '1814400s',
          },
          deposit_params: {
            max_deposit_period: '1814400s',
            min_deposit: [
              {
                denom: 'basecro',
                amount: '10000000',
              },
            ],
          },
        },
        transfer: {
          params: {
            receive_enabled: true,
            send_enabled: true,
          },
        },
        interchainaccounts: {
          host_genesis_state: {
            params: {
              allow_messages: [
                '/cosmos.bank.v1beta1.MsgSend',
              ],
            },
          },
        },
      },
    },
  },
  relayer: {
    mode: {
      clients: {
        enabled: true,
        refresh: true,
        misbehaviour: true,
      },
      connections: {
        enabled: true,
      },
      channels: {
        enabled: true,
      },
      packets: {
        enabled: true,
        tx_confirmation: true,
      },
    },
    rest: {
      enabled: true,
      host: '127.0.0.1',
      port: 3000,
    },
    chains: [
      {
        id: 'cronos_777-1',
        max_gas: 500000,
        gas_multiplier: 2,
        address_type: {
          derivation: 'ethermint',
          proto_type: {
            pk_type: '/ethermint.crypto.v1.ethsecp256k1.PubKey',
          },
        },
        gas_price: {
          price: 10000000000000000,
          denom: 'basetcro',
        },
        extension_options: [{
          type: 'ethermint_dynamic_fee',
          value: '1000000',
        }],
      },
      {
        id: 'chainmain-1',
        max_gas: 500000,
        gas_price: {
          price: 1000000,
          denom: 'basecro',
        },
      },
    ],
  },
}
